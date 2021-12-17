package slice_test

import (
	"calculator-vecty/model"
	"calculator-vecty/slice"
	"strconv"
	"testing"
)

func TestPushDigits(t *testing.T) {
	r := slice.NewCalculaterReducer()
	s := r.DigitReducer(r.State, slice.DigitAction{Digit: "1"})
	if s.Next != "1" || s.Display != "1" {
		t.Errorf("faild: %s CalculaterState=%+v", "1つ目の数字追加失敗", s)
	}
	s = r.DigitReducer(s, slice.DigitAction{Digit: "9"})
	if s.Next != "19" || s.Display != "19" {
		t.Errorf("faild: %s CalculaterState=%+v", "2つ目の数字追加失敗", s)
	}
}

func TestPushZero(t *testing.T) {
	r := slice.NewCalculaterReducer()
	s := r.DigitReducer(r.State, slice.DigitAction{Digit: "0"})
	if s.Next != "" || s.Display != "0" {
		t.Errorf("faild: %s CalculaterState=%+v", "1つ目の数字追加失敗", s)
	}
}

func pushButton(r *slice.CalculaterReducer, s slice.CalculaterState, p interface{}) slice.CalculaterState {
	switch v := p.(type) {
	case string:
		return r.DigitReducer(s, slice.DigitAction{Digit: v})
	case model.Operator:
		return r.OperatorReducer(s, slice.OperatorAction{Operator: v})
	case slice.InversionAction:
		return r.InversionReducer(s, p.(slice.InversionAction))
	case slice.PercentageAction:
		return r.PercentageReducer(s, p.(slice.PercentageAction))
	}
	panic("test broken")
}

func TestPushOperator(t *testing.T) {
	tests := []struct {
		name     string
		payloads []interface{}
		result   slice.CalculaterState
	}{
		{
			name:     "=",
			payloads: []interface{}{model.Equals},
			result:   slice.CalculaterState{Total: 0, Display: "0", Next: "", Operation: model.None},
		},
		{
			name:     "12+29=41",
			payloads: []interface{}{"1", "2", model.Plus, "2", "9", model.Equals},
			result:   slice.CalculaterState{Total: 12 + 29, Display: "41", Next: "", Operation: model.None},
		},
		{
			name:     "29-12=17",
			payloads: []interface{}{"2", "9", model.Minus, "1", "2", model.Equals},
			result:   slice.CalculaterState{Total: 29 - 12, Display: "17", Next: "", Operation: model.None},
		},
		{
			name:     "12-29=-17",
			payloads: []interface{}{"1", "2", model.Minus, "2", "9", model.Equals},
			result:   slice.CalculaterState{Total: 12 - 29, Display: "-17", Next: "", Operation: model.None},
		},
		{
			name:     "12-29=-17 演算子二回押し(+ -? -)",
			payloads: []interface{}{"1", "2", model.Plus, model.Minus, "2", "9", model.Equals},
			result:   slice.CalculaterState{Total: 12 - 29, Display: "-17", Next: "", Operation: model.None},
		},
		{
			name:     "-29=-29",
			payloads: []interface{}{model.Minus, "2", "9", model.Equals},
			result:   slice.CalculaterState{Total: -29, Display: "-29", Next: "", Operation: model.None},
		},
		{
			name:     "*29=0",
			payloads: []interface{}{model.Multiply, "2", "9", model.Equals},
			result:   slice.CalculaterState{Total: 0, Display: "0", Next: "", Operation: model.None},
		},
		{
			name:     "12*29=348",
			payloads: []interface{}{"1", "2", model.Multiply, "2", "9", model.Equals},
			result:   slice.CalculaterState{Total: 12 * 29, Display: "348", Next: "", Operation: model.None},
		},
		{
			name:     "12/29=",
			payloads: []interface{}{"1", "2", model.Divide, "2", "9", model.Equals},
			result:   slice.CalculaterState{Total: 12.0 / 29.0, Display: strconv.FormatFloat(12.0/29.0, 'f', -1, 64), Next: "", Operation: model.None},
		},
		{
			name:     "12/29+",
			payloads: []interface{}{"1", "2", model.Divide, "2", "9", model.Plus},
			result:   slice.CalculaterState{Total: 12.0 / 29.0, Display: strconv.FormatFloat(12.0/29.0, 'f', -1, 64), Next: "", Operation: model.Plus},
		},
		{
			name:     "12/29+0",
			payloads: []interface{}{"1", "2", model.Divide, "2", "9", model.Plus, "0"},
			result:   slice.CalculaterState{Total: 12.0 / 29.0, Display: "0", Next: "", Operation: model.Plus},
		},
		{
			name:     "12/29+01",
			payloads: []interface{}{"1", "2", model.Divide, "2", "9", model.Plus, "0", "1"},
			result:   slice.CalculaterState{Total: 12.0 / 29.0, Display: "1", Next: "1", Operation: model.Plus},
		},
		{
			name:     "12/29+12",
			payloads: []interface{}{"1", "2", model.Divide, "2", "9", model.Plus, "1", "2"},
			result:   slice.CalculaterState{Total: 12.0 / 29.0, Display: "12", Next: "12", Operation: model.Plus},
		},
		{
			name:     "12/29+12=",
			payloads: []interface{}{"1", "2", model.Divide, "2", "9", model.Plus, "1", "2", model.Equals},
			result:   slice.CalculaterState{Total: 12.0/29.0 + 12.0, Display: strconv.FormatFloat(12.0/29.0+12.0, 'f', -1, 64), Next: "", Operation: model.None},
		},
		{
			name:     "12/29+12*3=",
			payloads: []interface{}{"1", "2", model.Divide, "2", "9", model.Plus, "1", "2", model.Multiply, "3", model.Equals},
			result:   slice.CalculaterState{Total: (12.0/29.0 + 12.0) * 3.0, Display: strconv.FormatFloat((12.0/29.0+12.0)*3.0, 'f', -1, 64), Next: "", Operation: model.None},
		},
		{
			name:     "12/29+12*3-2=",
			payloads: []interface{}{"1", "2", model.Divide, "2", "9", model.Plus, "1", "2", model.Multiply, "3", model.Minus, "2", model.Equals},
			result:   slice.CalculaterState{Total: (12.0/29.0+12.0)*3.0 - 2.0, Display: strconv.FormatFloat((12.0/29.0+12.0)*3.0-2.0, 'f', -1, 64), Next: "", Operation: model.None},
		},
		{
			name:     "1/0=",
			payloads: []interface{}{"1", model.Divide, "0", model.Equals},
			result:   slice.CalculaterState{Total: 1, Display: "0", Next: "", Operation: model.None},
		},
	}

	for _, test := range tests {
		r := slice.NewCalculaterReducer()
		s := r.State
		for _, p := range test.payloads {
			s = pushButton(r, s, p)
		}
		if s != test.result {
			t.Errorf("faild: %s CalculaterState=%+v≠%+v", test.name, s, test.result)
		}
	}

}

func TestPushInversion(t *testing.T) {
	tests := []struct {
		name     string
		payloads []interface{}
		result   slice.CalculaterState
	}{
		{
			name:     "12->Inv",
			payloads: []interface{}{"1", "2", slice.InversionAction{}},
			result:   slice.CalculaterState{Total: 0, Display: "-12", Next: "-12", Operation: model.None},
		},
		{
			name:     "12+1+2->Inv",
			payloads: []interface{}{"1", "2", model.Plus, "1", model.Plus, "2", slice.InversionAction{}},
			result:   slice.CalculaterState{Total: 12 + 1, Display: "-2", Next: "-2", Operation: model.None},
		},
		{
			name:     "12+1+2->InvInv",
			payloads: []interface{}{"1", "2", model.Plus, "1", model.Plus, "2", slice.InversionAction{}, slice.InversionAction{}},
			result:   slice.CalculaterState{Total: 12 + 1, Display: "2", Next: "2", Operation: model.None},
		},
		{
			name:     "12+1=->Inv",
			payloads: []interface{}{"1", "2", model.Plus, "1", model.Equals, slice.InversionAction{}},
			result:   slice.CalculaterState{Total: (12 + 1) * -1, Display: "-13", Next: "", Operation: model.None},
		},
		{
			name:     "12+1=->InvInv",
			payloads: []interface{}{"1", "2", model.Plus, "1", model.Equals, slice.InversionAction{}, slice.InversionAction{}},
			result:   slice.CalculaterState{Total: (12 + 1) * -1 * -1, Display: "13", Next: "", Operation: model.None},
		},
		{
			name:     "12/7=->Inv",
			payloads: []interface{}{"1", "2", model.Divide, "7", model.Equals, slice.InversionAction{}},
			result:   slice.CalculaterState{Total: (12.0 / 7) * -1, Display: strconv.FormatFloat((12.0/7)*-1, 'f', -1, 64), Next: "", Operation: model.None},
		},
		{
			name:     "12/7=->InvInv",
			payloads: []interface{}{"1", "2", model.Divide, "7", model.Equals, slice.InversionAction{}, slice.InversionAction{}},
			result:   slice.CalculaterState{Total: (12.0 / 7) * -1 * -1, Display: strconv.FormatFloat((12.0/7)*-1*-1, 'f', -1, 64), Next: "", Operation: model.None},
		},
	}

	for _, test := range tests {
		r := slice.NewCalculaterReducer()
		s := r.State
		for _, p := range test.payloads {
			s = pushButton(r, s, p)
		}
		if s != test.result {
			t.Errorf("faild: %s CalculaterState=%+v≠%+v", test.name, s, test.result)
		}
	}
}

func TestPushPercentage(t *testing.T) {
	tests := []struct {
		name     string
		payloads []interface{}
		result   slice.CalculaterState
	}{
		{
			name:     "12%",
			payloads: []interface{}{"1", "2", slice.PercentageAction{}},
			result:   slice.CalculaterState{Total: 0, Display: "0.12", Next: "0.12", Operation: model.None},
		},
		{
			name:     "12%%",
			payloads: []interface{}{"1", "2", slice.PercentageAction{}, slice.PercentageAction{}},
			result:   slice.CalculaterState{Total: 0, Display: "0.0012", Next: "0.0012", Operation: model.None},
		},
		{
			name:     "12+1+2%",
			payloads: []interface{}{"1", "2", model.Plus, "1", model.Plus, "2", slice.PercentageAction{}},
			result:   slice.CalculaterState{Total: 12 + 1, Display: "0.02", Next: "0.02", Operation: model.None},
		},
		{
			name:     "12+1=%",
			payloads: []interface{}{"1", "2", model.Plus, "1", model.Equals, slice.PercentageAction{}},
			result:   slice.CalculaterState{Total: (12.0 + 1) / 100, Display: "0.13", Next: "", Operation: model.None},
		},
		{
			name:     "12/7=%",
			payloads: []interface{}{"1", "2", model.Divide, "7", model.Equals, slice.PercentageAction{}},
			result:   slice.CalculaterState{Total: 0.01714285714285714, Display: "0.01714285714285714", Next: "", Operation: model.None},
		},
	}

	for _, test := range tests {
		r := slice.NewCalculaterReducer()
		s := r.State
		for _, p := range test.payloads {
			s = pushButton(r, s, p)
		}
		if s != test.result {
			t.Errorf("faild: %s CalculaterState=%+v≠%+v", test.name, s, test.result)
		}
	}
}
