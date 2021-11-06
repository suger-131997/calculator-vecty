package storeutil

import (
	"reflect"

	"github.com/dannypsnl/redux/v2/rematch"
	"github.com/dannypsnl/redux/v2/store"
	"github.com/hexops/vecty"
)

var stores = make(map[reflect.Type]interface{})
var stateTypeToStoreType = make(map[reflect.Type]reflect.Type)
var stateTypeToComponents = make(map[reflect.Type][]vecty.Component)
var actTypeToStateType = make(map[reflect.Type]reflect.Type)
var actTypeToRematcher = make(map[reflect.Type]string)
var rootStore *store.Store

func UseState(stateType reflect.Type, c vecty.Component) interface{} {
	stateTypeToComponents[stateType] = append(stateTypeToComponents[stateType], c)
	return rootStore.StateOf(stores[stateType])
}

func Dispatch(a interface{}) {
	at := reflect.TypeOf(a)
	st := actTypeToStateType[at]

	r := reflect.Indirect(
		reflect.ValueOf(stores[st]).Convert(stateTypeToStoreType[st]),
	).FieldByName(actTypeToRematcher[at]).MethodByName("With")

	rootStore.Dispatch(
		r.Call(
			[]reflect.Value{
				reflect.ValueOf(a),
			},
		)[0].Interface().(*rematch.Action),
	)

	copyed := make([]vecty.Component, 0)
	copyed = append(copyed, stateTypeToComponents[st]...)
	stateTypeToComponents[st] = make([]vecty.Component, 0)
	for _, c := range copyed {
		vecty.Rerender(c)
	}
}

func initReducer(reducer interface{}) {
	rpv := reflect.ValueOf(reducer)
	rpt := reflect.TypeOf(reducer)
	rt := rpt.Elem()
	sv, _ := rt.FieldByName("State")
	st := sv.Type

	stores[st] = reducer
	stateTypeToStoreType[st] = rpt
	stateTypeToComponents[st] = make([]vecty.Component, 0)
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		if funcName := f.Tag.Get("action"); funcName != "" {
			at := rpv.MethodByName(funcName).Type().In(1)
			actTypeToStateType[at] = st
			actTypeToRematcher[at] = f.Name
		}
	}
}

func Init(reducers ...interface{}) {
	for _, r := range reducers {
		initReducer(r)
	}

	rootStore = store.New(reducers...)
}
