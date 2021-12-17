package helper

import (
	"calculator-vecty/model"
	"calculator-vecty/slice"
	"calculator-vecty/storeutil"
	"unicode"

	"github.com/hexops/vecty"
)

func PushDigit(digit string) func(event *vecty.Event) {
	if len(digit) > 1 || !unicode.IsDigit(rune(digit[0])) {
		panic("Invalid digit")
	}

	return func(event *vecty.Event) {
		storeutil.Dispatch(slice.DigitAction{Digit: digit})
	}
}

func PushDot() func(event *vecty.Event) {
	return func(event *vecty.Event) {
		storeutil.Dispatch(slice.DotAction{})
	}
}

func PushOperator(operator model.Operator) func(event *vecty.Event) {
	return func(event *vecty.Event) {
		storeutil.Dispatch(slice.OperatorAction{Operator: operator})
	}
}

func PushAllClear() func(event *vecty.Event) {
	return func(event *vecty.Event) {
		storeutil.Dispatch(slice.AllClearAction{})
	}
}

func PushInversion() func(event *vecty.Event) {
	return func(event *vecty.Event) {
		storeutil.Dispatch(slice.InversionAction{})
	}
}

func PushPercentage() func(event *vecty.Event) {
	return func(event *vecty.Event) {
		storeutil.Dispatch(slice.PercentageAction{})
	}
}
