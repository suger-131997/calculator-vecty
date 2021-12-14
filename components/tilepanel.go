package components

import (
	"calculator-vecty/helper"
	"calculator-vecty/model"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/style"
)

type tilePanel struct {
	vecty.Core
}

func (b *tilePanel) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("tile", "is-parent", "is-vertical"),
			vecty.Style("padding", "0"),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("tile", "is-parent"),
				style.MinHeight("15.3vh"),
				vecty.Style("padding", "0"),
			),
			&tileButton{Label: "AC", OnClick: helper.PushAllClear()},
			&tileButton{Label: "+/-"},
			&tileButton{Label: "%"},
			&tileButton{Label: "÷", White: true, OnClick: helper.PushOperator(model.Divide)},
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("tile", "is-parent"),
				style.MinHeight("15.3vh"),
				vecty.Style("padding", "0"),
			),
			&tileButton{Label: "7", OnClick: helper.PushDigit("7")},
			&tileButton{Label: "8", OnClick: helper.PushDigit("8")},
			&tileButton{Label: "9", OnClick: helper.PushDigit("9")},
			&tileButton{Label: "x", White: true, OnClick: helper.PushOperator(model.Multiply)},
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("tile", "is-parent"),
				style.MinHeight("15.3vh"),
				vecty.Style("padding", "0"),
			),
			&tileButton{Label: "4", OnClick: helper.PushDigit("4")},
			&tileButton{Label: "5", OnClick: helper.PushDigit("5")},
			&tileButton{Label: "6", OnClick: helper.PushDigit("6")},
			&tileButton{Label: "-", White: true, OnClick: helper.PushOperator(model.Minus)},
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("tile", "is-parent"),
				style.MinHeight("15.3vh"),
				vecty.Style("padding", "0"),
			),
			&tileButton{Label: "1", OnClick: helper.PushDigit("1")},
			&tileButton{Label: "2", OnClick: helper.PushDigit("2")},
			&tileButton{Label: "3", OnClick: helper.PushDigit("3")},
			&tileButton{Label: "+", White: true, OnClick: helper.PushOperator(model.Plus)},
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("tile", "is-parent"),
				style.MinHeight("15.3vh"),
				vecty.Style("padding", "0"),
			),
			&tileButton{Label: "0", Wide: true, OnClick: helper.PushDigit("0")},
			&tileButton{Label: "."},
			&tileButton{Label: "=", White: true, OnClick: helper.PushOperator(model.Equals)},
		),
	)
}
