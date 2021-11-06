package components

import (
	"todos-vecty/model"
	"todos-vecty/storeutil"
	"todos-vecty/storeutil/slice"

	"github.com/hexops/vecty"
)

type VisibleTodoList struct {
	vecty.Core
}

func (a *VisibleTodoList) Render() vecty.ComponentOrHTML {
	todosState, _ := storeutil.UseState(slice.TodosStateType, a).(slice.TodosState)
	filterState, _ := storeutil.UseState(slice.FilterStateType, a).(slice.FilterState)

	todos := make([]model.Todo, 0)

	switch filterState.Type {
	case model.All:
		todos = append(todos, todosState.Todos...)
	case model.Active:
		for _, todo := range todosState.Todos {
			if !todo.Completed {
				todos = append(todos, todo)
			}
		}
	case model.Completed:
		for _, todo := range todosState.Todos {
			if todo.Completed {
				todos = append(todos, todo)
			}
		}
	}

	return &todoList{
		Todos: todos,
		OnClick: func(id int) {
			storeutil.Dispatch(slice.CompleteTodoAction{Payload: id})
		},
	}
}
