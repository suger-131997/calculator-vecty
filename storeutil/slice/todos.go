package slice

import (
	"reflect"
	"todos-vecty/model"

	"github.com/dannypsnl/redux/v2/rematch"
)

var TodosStateType = reflect.TypeOf((*TodosState)(nil)).Elem()

type TodosState struct {
	Todos []model.Todo
}

type TodosReducer struct {
	rematch.Reducer
	State TodosState

	Add      *rematch.Action `action:"AddTodo"`
	Complete *rematch.Action `action:"CompleteTodo"`
}

func NewTodosReducer() *TodosReducer {
	return &TodosReducer{State: TodosState{Todos: make([]model.Todo, 0)}}
}

type addTodoAction struct {
	payload model.Todo
}

var nextTodoId int

func NewAddTodoAction(t string) addTodoAction {
	nextTodoId++
	return addTodoAction{payload: model.Todo{Id: nextTodoId, Title: t, Completed: false}}
}

func (t *TodosReducer) AddTodo(s TodosState, a addTodoAction) TodosState {
	return TodosState{Todos: append(s.Todos, a.payload)}
}

type CompleteTodoAction struct {
	Payload int
}

func (t *TodosReducer) CompleteTodo(s TodosState, a CompleteTodoAction) TodosState {
	newTodos := make([]model.Todo, 0)

	for _, todo := range s.Todos {
		if todo.Id == a.Payload {
			todo.Completed = !todo.Completed
		}
		newTodos = append(newTodos, todo)
	}

	return TodosState{Todos: newTodos}
}

func init() {
	nextTodoId = 0
}
