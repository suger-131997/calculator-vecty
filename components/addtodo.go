package components

import (
	"todos-vecty/storeutil"
	"todos-vecty/storeutil/slice"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/hexops/vecty/prop"
)

type AddTodo struct {
	vecty.Core

	newItemTitle string
}

func (a *AddTodo) onNewItemTitleInput(e *vecty.Event) {
	a.newItemTitle = e.Target.Get("value").String()
	vecty.Rerender(a)
}

func (a *AddTodo) onAdd(e *vecty.Event) {
	storeutil.Dispatch(slice.NewAddTodoAction(a.newItemTitle))
	a.newItemTitle = ""
	vecty.Rerender(a)
}

func (a *AddTodo) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("block", "mb-0"),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("columns", "is-grouped"),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("column", "is-half"),
				),
				elem.Input(
					vecty.Markup(
						vecty.Class("input", "is-rounded"),
						prop.Value(a.newItemTitle),
						event.Input(a.onNewItemTitleInput),
					),
				),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("column", "is-one-quarter"),
				),
				elem.Button(
					vecty.Markup(
						vecty.Class("button", "is-primary"),
						event.Click(a.onAdd).PreventDefault(),
					),
					vecty.Text("Add Todo"),
				),
			),
		),
	)
}
