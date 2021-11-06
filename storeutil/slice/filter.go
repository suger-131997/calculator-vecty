package slice

import (
	"reflect"
	"todos-vecty/model"

	"github.com/dannypsnl/redux/v2/rematch"
)

var FilterStateType = reflect.TypeOf((*FilterState)(nil)).Elem()

type FilterState struct {
	Type model.FilterType
}

type FilterReducer struct {
	rematch.Reducer
	State FilterState

	Chenge *rematch.Action `action:"ChengeFilter"`
}

type ChengeFilterAction struct {
	Payload model.FilterType
}

func (t *FilterReducer) ChengeFilter(s FilterState, a ChengeFilterAction) FilterState {
	return FilterState{Type: a.Payload}
}
