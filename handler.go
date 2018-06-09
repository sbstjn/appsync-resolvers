package router

import (
	"encoding/json"
	"reflect"
)

type handler struct {
	function interface{}
}

func (h *handler) prepare(payload json.RawMessage) ([]reflect.Value, error) {
	if reflect.TypeOf(h.function).NumIn() == 0 {
		return nil, nil
	}

	argsType := reflect.TypeOf(h.function).In(0)
	args := reflect.New(argsType)

	if err := json.Unmarshal(payload, args.Interface()); err != nil {
		return nil, err
	}

	return append([]reflect.Value{}, args.Elem()), nil
}

func (h *handler) call(payload json.RawMessage) (interface{}, error) {
	args, err := h.prepare(payload)

	if err != nil {
		return nil, err
	}

	returnValues := reflect.ValueOf(h.function).Call(args)
	var returnData interface{}
	var returnError error

	if len(returnValues) == 2 {
		returnData = returnValues[0].Interface()
	}

	if err := returnValues[len(returnValues)-1].Interface(); err != nil {
		returnError = err.(error)
	}

	return returnData, returnError
}

func (h handler) validate() error {
	handler := reflect.TypeOf(h.function)

	return validators.run(handler)
}
