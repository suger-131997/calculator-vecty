package components

import (
	"todos-vecty/model"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

type todoList struct {
	vecty.Core

	Todos   []model.Todo `vecty:"prop"`
	OnClick func(int)    `vecty:"prop"`
}

func (t *todoList) Render() vecty.ComponentOrHTML {
	var items vecty.List

	for _, todo := range t.Todos {
		_todo := todo
		items = append(items, &todoListItem{
			Text:      todo.Title,
			Completed: todo.Completed,
			OnClick: func(e *vecty.Event) {
				t.OnClick(_todo.Id)
			},
		})
	}

	return elem.UnorderedList(
		vecty.Markup(),
		items,
	)
}
