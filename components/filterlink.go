package components

import (
	"todos-vecty/model"
	"todos-vecty/storeutil"
	"todos-vecty/storeutil/slice"

	"github.com/hexops/vecty"
)

type FilterLink struct {
	vecty.Core

	Type  model.FilterType `vecty:"prop"`
	Label string           `vecty:"prop"`
}

func (f *FilterLink) Render() vecty.ComponentOrHTML {
	isActive := storeutil.UseState(slice.FilterStateType, f).(slice.FilterState).Type == f.Type

	return &link{
		Type:     f.Type,
		IsActive: isActive,
		Label:    f.Label,
		OnClick: func(event *vecty.Event) {
			storeutil.Dispatch(slice.ChengeFilterAction{Payload: f.Type})
		},
	}
}
