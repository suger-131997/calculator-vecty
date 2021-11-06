package slice

import (
	"reflect"
	"todos-vecty/model"

	"github.com/dannypsnl/redux/v2/rematch"
)

var TodosStateType = reflect.TypeOf([]model.Todo{})

type TodosReducer struct {
	rematch.Reducer
	State []model.Todo

	Add      *rematch.Action `action:"AddTodo"`
	Complete *rematch.Action `action:"CompleteTodo"`
}

type addTodoAction struct {
	payload model.Todo
}

var nextTodoId int

func NewAddTodoAction(t string) addTodoAction {
	nextTodoId++
	return addTodoAction{payload: model.Todo{Id: nextTodoId, Title: t, Completed: false}}
}

func (t *TodosReducer) AddTodo(s []model.Todo, a addTodoAction) []model.Todo {
	return append(s, a.payload)
}

type completeTodoAction struct {
	payload int
}

func NewCompleteTodoAction(id int) completeTodoAction {
	return completeTodoAction{payload: id}
}

func (t *TodosReducer) CompleteTodo(s []model.Todo, a completeTodoAction) []model.Todo {
	newState := make([]model.Todo, 0)

	for _, todo := range s {
		if todo.Id == a.payload {
			todo.Completed = !todo.Completed
		}
		newState = append(newState, todo)
	}

	return newState
}

func init() {
	nextTodoId = 0
}
