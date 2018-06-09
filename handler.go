package router

import (
	"encoding/json"
	"reflect"
)

// Handler is an abstract function
type Handler struct {
	function interface{}
}

// Prepare parses event payload
func (h *Handler) Prepare(payload json.RawMessage) ([]reflect.Value, error) {
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

// Call the handler and pass event
func (h *Handler) Call(payload json.RawMessage) (interface{}, error) {
	args, err := h.Prepare(payload)

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

// Validate checks if passed handler is valid
func (h Handler) Validate() error {
	handler := reflect.TypeOf(h.function)

	return validators.run(handler)
}
