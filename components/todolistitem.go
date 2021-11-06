package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
)

type todoListItem struct {
	vecty.Core

	Text      string             `vecty:"prop"`
	Completed bool               `vecty:"prop"`
	OnClick   func(*vecty.Event) `vecty:"prop"`
}

func (t *todoListItem) Render() vecty.ComponentOrHTML {
	var textdeco = "none"
	if t.Completed {
		textdeco = "line-through"
	}

	return elem.ListItem(
		vecty.Markup(
			vecty.Class("is-size-3"),
			vecty.Style("text-decoration", textdeco),
			event.Click(t.OnClick),
		),
		vecty.Text(
			"・ "+t.Text,
		),
	)
}
