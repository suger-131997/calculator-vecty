package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/hexops/vecty/style"
)

type tileButton struct {
	vecty.Core

	Wide  bool `vecty:"prop"`
	White bool `vecty:"prop"`

	Label   string                   `vecty:"prop"`
	OnClick func(event *vecty.Event) `vecty:"prop"`
}

func (b *tileButton) Render() vecty.ComponentOrHTML {
	var width vecty.Applyer
	if b.Wide {
		width = style.Width("50vw")
	} else {
		width = style.Width("25vw")
	}

	var bc vecty.Applyer
	if b.White {
		bc = vecty.Class("button", "is-light")
	} else {
		bc = vecty.Class("button", "is-primary")
	}

	return elem.Div(
		vecty.Markup(width),
		elem.Button(
			vecty.Markup(
				bc,
				style.Width("100%"),
				style.Height("100%"),
				vecty.Style("font-size", "6vmin"),
				vecty.Style("border-radius", "0"),
				event.Click(b.OnClick).PreventDefault(),
			),
			vecty.Text(b.Label),
		),
	)
}
