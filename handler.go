package router

import (
	"encoding/json"
	"errors"
	"fmt"
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

	if kind := handler.Kind(); kind != reflect.Func {
		return fmt.Errorf("Handler is not a function, but %s", kind)
	}

	if handler.NumIn() > 1 {
		return errors.New("Handler must not have more than one argument")
	}

	if handler.NumIn() == 1 && handler.In(0).Kind() != reflect.Struct {
		return errors.New("Handler argument must be struct")
	}

	if handler.NumOut() > 2 {
		return errors.New("Handler must not have more than two return values")
	}

	if handler.NumOut() < 1 {
		return errors.New("Handler must have at least one return value")
	}

	if last := handler.Out(handler.NumOut() - 1); last.String() != "error" {
		return errors.New("Last or only return value must be an error")
	}

	return nil
}
