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
	paramType := reflect.TypeOf(h.function).In(0)
	param := reflect.New(paramType)

	if err := json.Unmarshal(payload, param.Interface()); err != nil {
		return nil, err
	}

	return append([]reflect.Value{}, param.Elem()), nil
}

// Call the handler and pass event
func (h *Handler) Call(in []reflect.Value) (interface{}, error) {
	response := reflect.ValueOf(h.function).Call(in)

	return response[0].Interface(), response[1].Interface().(error)
}
