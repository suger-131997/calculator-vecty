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
		elem.Form(
			elem.Input(
				vecty.Markup(
					prop.Value(a.newItemTitle),
					event.Input(a.onNewItemTitleInput),
				),
			),
			elem.Button(
				vecty.Markup(
					event.Click(a.onAdd).PreventDefault(),
				),
				vecty.Text("Add Todo"),
			),
		),
	)
}
