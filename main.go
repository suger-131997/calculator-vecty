package main

import (
	"todos-vecty/pages"

	_ "github.com/dannypsnl/redux/v2"

	"github.com/hexops/vecty"
)

func main() {
	vecty.SetTitle("Todos")

	p := &pages.PageView{}
	vecty.RenderBody(p)
}
