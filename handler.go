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
	argsType := reflect.TypeOf(h.function).In(0)
	args := reflect.New(argsType)

	if err := json.Unmarshal(payload, args.Interface()); err != nil {
		return nil, err
	}

	return append([]reflect.Value{}, args.Elem()), nil
}

// Call the handler and pass event
func (h *Handler) Call(args []reflect.Value) (interface{}, error) {
	var err error
	response := reflect.ValueOf(h.function).Call(args)

	if response[1].Interface() != nil {
		err = response[1].Interface().(error)
	}

	return response[0].Interface(), err
}
