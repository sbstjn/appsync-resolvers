package resolvers

import (
	"encoding/json"
	"reflect"
)

type resolver struct {
	function interface{}
}

func (r *resolver) hasArguments() bool {
	return reflect.TypeOf(r.function).NumIn() == 1
}

func (r *resolver) hasArgumentsAndIdentity() bool {
	return reflect.TypeOf(r.function).NumIn() == 2
}

func (r *resolver) call(p json.RawMessage, i string) (interface{}, error) {
	var args []reflect.Value
	var err error

	if r.hasArguments() {
		pld := payload{p}
		args, err = pld.parse(reflect.TypeOf(r.function).In(0))

		if err != nil {
			return nil, err
		}
	} else if r.hasArgumentsAndIdentity() {
		pld := payload{p}
		args, err = pld.parse(reflect.TypeOf(r.function).In(0))
		args = append(args, reflect.ValueOf(i))

		if err != nil {
			return nil, err
		}
	}

	returnValues := reflect.ValueOf(r.function).Call(args)
	var returnData interface{}
	var returnError error

	if len(returnValues) == 2 {
		returnData = returnValues[0].Interface()
	}

	if err := returnValues[len(returnValues)-1].Interface(); err != nil {
		returnError = returnValues[len(returnValues)-1].Interface().(error)
	}

	return returnData, returnError
}
