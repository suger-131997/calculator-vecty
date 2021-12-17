package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

type PageView struct {
	vecty.Core
}

func (p *PageView) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Div(
			vecty.Markup(
				vecty.Class("container", "is-fluid", "has-background-grey-light"),
				vecty.Style("padding", "0"),
			),
			&display{},
			&tilePanel{},
		),
	)
}
