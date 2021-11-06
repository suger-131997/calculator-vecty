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

	todosArray := make([]model.Todo, 0)

	if filterState.Type == model.All {
		todosArray = append(todosArray, todosState.Todos...)
	} else if filterState.Type == model.Active {

		for _, todo := range todosState.Todos {
			if !todo.Completed {
				todosArray = append(todosArray, todo)
			}
		}
	} else if filterState.Type == model.Completed {
		for _, todo := range todosState.Todos {
			if todo.Completed {
				todosArray = append(todosArray, todo)
			}
		}
	}

	return &todoList{
		Todos: todosArray,
		OnClick: func(id int) {
			storeutil.Dispatch(slice.CompleteTodoAction{Payload: id})
		},
	}
}
