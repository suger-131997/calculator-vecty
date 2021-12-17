package slice

import (
	"calculator-vecty/model"
	"reflect"
	"strconv"

	"github.com/dannypsnl/redux/v2/rematch"
)

var CalculaterStateType = reflect.TypeOf((*CalculaterState)(nil)).Elem()

type CalculaterState struct {
	Total     float64
	Display   string
	Next      string
	Operation model.Operator
}

type CalculaterReducer struct {
	rematch.Reducer
	State CalculaterState

	Digit      *rematch.Action `action:"DigitReducer"`
	Dot        *rematch.Action `action:"DotReducer"`
	Operator   *rematch.Action `action:"OperatorReducer"`
	AllClear   *rematch.Action `action:"AllClearReducer"`
	Inversion  *rematch.Action `action:"InversionReducer"`
	Percentage *rematch.Action `action:"PercentageReducer"`
}

func NewCalculaterReducer() *CalculaterReducer {
	return &CalculaterReducer{State: CalculaterState{Total: 0, Display: "0", Next: "", Operation: model.None}}
}

type DigitAction struct {
	Digit string
}

func (r *CalculaterReducer) DigitReducer(s CalculaterState, a DigitAction) CalculaterState {
	if s.Next == "" && a.Digit == "0" {
		return CalculaterState{Total: s.Total, Display: "0", Next: "", Operation: s.Operation}
	}

	next := s.Next + a.Digit
	return CalculaterState{Total: s.Total, Display: next, Next: next, Operation: s.Operation}
}

type DotAction struct {
}

func (r *CalculaterReducer) DotReducer(s CalculaterState, a DotAction) CalculaterState {
	if s.Next == "" {
		return CalculaterState{Total: s.Total, Display: "0.", Next: "0.", Operation: s.Operation}
	}

	next := s.Next + "."
	return CalculaterState{Total: s.Total, Display: next, Next: next, Operation: s.Operation}
}

type OperatorAction struct {
	Operator model.Operator
}

func operation(total float64, next string, o model.Operator) (float64, string) {
	i, _ := strconv.ParseFloat(next, 64)
	switch o {
	case model.Plus:
		t := total + i
		d := strconv.FormatFloat(t, 'f', -1, 64)
		return t, d
	case model.Minus:
		t := total - i
		d := strconv.FormatFloat(t, 'f', -1, 64)
		return t, d
	case model.Multiply:
		t := total * i
		d := strconv.FormatFloat(t, 'f', -1, 64)
		return t, d
	case model.Divide:
		if i != 0.0 {
			t := total / i
			d := strconv.FormatFloat(t, 'f', -1, 64)
			return t, d
		}
	}

	return total, next
}

func (r *CalculaterReducer) OperatorReducer(s CalculaterState, a OperatorAction) CalculaterState {
	if a.Operator != model.Equals {
		if s.Display == s.Next {
			if s.Operation == model.None {
				t, _ := strconv.ParseFloat(s.Next, 64)
				return CalculaterState{Total: t, Display: s.Display, Next: "", Operation: a.Operator}
			}
			t, d := operation(s.Total, s.Next, s.Operation)
			return CalculaterState{Total: t, Display: d, Next: "", Operation: a.Operator}
		}
		return CalculaterState{Total: s.Total, Display: s.Display, Next: s.Next, Operation: a.Operator}
	}

	if s.Operation != model.None {
		t, d := operation(s.Total, s.Next, s.Operation)
		if d == "" {
			return CalculaterState{Total: t, Display: "0", Next: "", Operation: model.None}
		}
		return CalculaterState{Total: t, Display: d, Next: "", Operation: model.None}
	}

	ns := s
	return ns
}

type AllClearAction struct {
}

func (r *CalculaterReducer) AllClearReducer(s CalculaterState, a AllClearAction) CalculaterState {
	return CalculaterState{Total: 0, Display: "0", Next: "", Operation: model.None}
}

type InversionAction struct {
}

func (r *CalculaterReducer) InversionReducer(s CalculaterState, a InversionAction) CalculaterState {
	if s.Next == "" {
		if string(s.Display[0]) == "-" {
			return CalculaterState{Total: s.Total * -1, Display: s.Display[1:], Next: "", Operation: model.None}
		}
		return CalculaterState{Total: s.Total * -1, Display: "-" + s.Display, Next: "", Operation: model.None}
	}
	n, _ := strconv.ParseFloat(s.Next, 64)
	n = -n
	d := strconv.FormatFloat(n, 'f', -1, 64)
	return CalculaterState{Total: s.Total, Display: d, Next: d, Operation: model.None}
}

type PercentageAction struct {
}

func (r *CalculaterReducer) PercentageReducer(s CalculaterState, a PercentageAction) CalculaterState {
	if s.Next == "" {
		t := s.Total / 100
		return CalculaterState{Total: t, Display: strconv.FormatFloat(t, 'f', -1, 64), Next: "", Operation: model.None}
	}
	n, _ := strconv.ParseFloat(s.Next, 64)
	n = n / 100
	d := strconv.FormatFloat(n, 'f', -1, 64)
	return CalculaterState{Total: s.Total, Display: d, Next: d, Operation: model.None}
}
