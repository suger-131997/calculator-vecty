package slice

import (
	"reflect"
	"todos-vecty/model"

	"github.com/dannypsnl/redux/v2/rematch"
)

var FilterStateType = reflect.TypeOf((*model.FilterType)(nil)).Elem() //reflect.TypeOf(model.All)

type FilterReducer struct {
	rematch.Reducer
	State model.FilterType

	Chenge *rematch.Action `action:"ChengeFilter"`
}

func (t *FilterReducer) ChengeFilter(s, a model.FilterType) model.FilterType {
	return a
}
