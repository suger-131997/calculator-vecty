package components

import (
	"todos-vecty/model"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/hexops/vecty/prop"
)

type link struct {
	vecty.Core

	Type     model.FilterType         `vecty:"prop"`
	IsActive bool                     `vecty:"prop"`
	Label    string                   `vecty:"prop"`
	OnClick  func(event *vecty.Event) `vecty:"prop"`
}

func (l *link) Render() vecty.ComponentOrHTML {
	return elem.Button(
		vecty.Markup(
			vecty.Class("button"),
			prop.Disabled(l.IsActive),
			event.Click(l.OnClick).PreventDefault(),
		),
		vecty.Text(l.Label),
	)
}
