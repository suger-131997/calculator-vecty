package main

import (
	"syscall/js"
	"todos-vecty/pages"
	"todos-vecty/storeutil"
	"todos-vecty/storeutil/slice"

	_ "github.com/dannypsnl/redux/v2"

	"github.com/hexops/vecty"
)

func main() {
	vecty.SetTitle("Todos")

	meta := js.Global().Get("document").Call("createElement", "meta")
	meta.Set("name", "viewport")
	meta.Set("content", "width=device-width, initial-scale=1")
	js.Global().Get("document").Get("head").Call("appendChild", meta)

	vecty.AddStylesheet("https://cdn.jsdelivr.net/npm/bulma@0.9.3/css/bulma.min.css")

	storeutil.Init(slice.NewTodosReducer(), slice.NewFilterReducer())

	p := &pages.PageView{}
	vecty.RenderBody(p)
}
