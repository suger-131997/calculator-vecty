package main

import (
	"calculator-vecty/components"
	"calculator-vecty/slice"
	"calculator-vecty/storeutil"
	"syscall/js"

	_ "github.com/dannypsnl/redux/v2"

	"github.com/hexops/vecty"
)

func main() {
	vecty.SetTitle("Vecty Calculater")

	meta := js.Global().Get("document").Call("createElement", "meta")
	meta.Set("name", "viewport")
	meta.Set("content", "width=device-width, initial-scale=1, user-scalable=no")
	js.Global().Get("document").Get("head").Call("appendChild", meta)

	vecty.AddStylesheet("https://cdn.jsdelivr.net/npm/bulma@0.9.3/css/bulma.min.css")
	vecty.AddStylesheet("./style/scrollbar.css")

	storeutil.Init(slice.NewCalculaterReducer())

	p := &components.PageView{}
	vecty.RenderBody(p)
}
