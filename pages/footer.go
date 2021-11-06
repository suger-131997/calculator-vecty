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
			vecty.Style("display", "flex"),
		),
		elem.Span(
			vecty.Markup(),
			vecty.Text("Show:"),
		),
		&components.FilterLink{Type: model.All, Label: "All"},
		&components.FilterLink{Type: model.Active, Label: "Active"},
		&components.FilterLink{Type: model.Completed, Label: "Completed"},
	)
}
