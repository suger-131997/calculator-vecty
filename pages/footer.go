package pages

import (
	"todos-vecty/components"
	"todos-vecty/model"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

type footer struct {
	vecty.Core
}

func (f *footer) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("level", "mt-3"),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("level-left"),
			),
			elem.Paragraph(
				vecty.Markup(
					vecty.Class("level-item"),
				),
				elem.Strong(
					vecty.Markup(
						vecty.Class("subtitle", "mr-3"),
					),
					vecty.Text("Show:"),
				),
			),
			elem.Paragraph(
				vecty.Markup(
					vecty.Class("level-item", "mx-0"),
				),
				&components.FilterLink{Type: model.All, Label: "All"},
			),
			elem.Paragraph(
				vecty.Markup(
					vecty.Class("level-item", "mx-0"),
				),
				&components.FilterLink{Type: model.Active, Label: "Active"},
			),
			elem.Paragraph(
				vecty.Markup(
					vecty.Class("level-item", "mx-0"),
				),
				&components.FilterLink{Type: model.Completed, Label: "Completed"},
			),
		),
	)
}
