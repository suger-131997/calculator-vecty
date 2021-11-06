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
	todosState, _ := storeutil.UseState(slice.TodosStateType, a).([]model.Todo)
	filterState, _ := storeutil.UseState(slice.FilterStateType, a).(model.FilterType)

	todosArray := make([]model.Todo, 0)

	if filterState == model.All {
		todosArray = append(todosArray, todosState...)
	} else if filterState == model.Active {

		for _, todo := range todosState {
			if !todo.Completed {
				todosArray = append(todosArray, todo)
			}
		}
	} else if filterState == model.Completed {
		for _, todo := range todosState {
			if todo.Completed {
				todosArray = append(todosArray, todo)
			}
		}
	}

	return &todoList{
		Todos: todosArray,
		OnClick: func(id int) {
			storeutil.Dispatch(slice.NewCompleteTodoAction(id))
		},
	}
}
