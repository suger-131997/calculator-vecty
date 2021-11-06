package pages

import (
	"todos-vecty/components"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

type PageView struct {
	vecty.Core
}

func (p *PageView) Render() vecty.ComponentOrHTML {
	return elem.Body(
		&components.AddTodo{},
		&components.VisibleTodoList{},
		&footer{},
	)
}
