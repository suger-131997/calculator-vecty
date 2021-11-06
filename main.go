package main

import (
	"todos-vecty/pages"
	"todos-vecty/storeutil"
	"todos-vecty/storeutil/slice"

	_ "github.com/dannypsnl/redux/v2"

	"github.com/hexops/vecty"
)

func main() {
	vecty.SetTitle("Todos")

	storeutil.Init(slice.NewTodosReducer(), slice.NewFilterReducer())

	p := &pages.PageView{}
	vecty.RenderBody(p)
}
