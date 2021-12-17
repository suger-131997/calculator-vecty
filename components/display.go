package components

import (
	"calculator-vecty/slice"
	"calculator-vecty/storeutil"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/style"
)

type display struct {
	vecty.Core
}

func (d *display) Render() vecty.ComponentOrHTML {
	display := storeutil.UseState(slice.CalculaterStateType, d).(slice.CalculaterState).Display

	return elem.Div(
		vecty.Markup(
			vecty.Class("grey-light"),
		),
		elem.Paragraph(
			vecty.Markup(
				vecty.Class("title", "has-text-right"),
				style.Width("100%"),
				style.Height("23.5vh"),
				vecty.Style("font-size", "18vh"),
			),
			vecty.Text(display),
		),
	)
}
