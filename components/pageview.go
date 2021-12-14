package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/style"
)

type PageView struct {
	vecty.Core
}

func (p *PageView) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Div(
			vecty.Markup(
				vecty.Class("container", "is-fluid", "has-background-grey-light"),
				vecty.Style("padding", "0 0 0 0"),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("tile", "is-vertical", "is-ancestor"),
					style.Margin(style.Size("0 0 0 0")),
				),
				elem.Div(
					vecty.Markup(
						vecty.Class("tile", "is-vertical", "is-parent"),
						style.Margin(style.Size("0 0 0 0")),
						vecty.Style("padding", "0 0 0 0"),
					),
					&display{},
					&tilePanel{},
				),
			),
		),
	)
}
