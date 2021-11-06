package main

import (
	"todos-vecty/pages"

	_ "github.com/dannypsnl/redux/v2"

	"github.com/hexops/vecty"
)

func main() {
	vecty.SetTitle("Todos")

	// storeutil.TodosStoreInit()
	// storeutil.FilterStoreInit()

	p := &pages.PageView{}
	vecty.RenderBody(p)
}

// func main() {
// 	storeutil.Dispatch(todos.NewAddTodoAction("Hogehoge"))
// 	storeutil.Dispatch(todos.NewAddTodoAction("FooFoo"))
// 	a := storeutil.UseState(todos.StateType).([]model.Todo)
// 	fmt.Println(a)

// 	b := storeutil.UseState(filter.StateType).(model.FilterType)
// 	fmt.Println(b)
// 	storeutil.Dispatch(model.Active)
// 	b = storeutil.UseState(filter.StateType).(model.FilterType)
// 	fmt.Println(b)

// 	storeutil.Dispatch(todos.NewCompleteTodoAction(1))
// 	a = storeutil.UseState(todos.StateType).([]model.Todo)
// 	fmt.Println(a)
// }
