package router

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

// Request stores all information from AppSync request
type Request struct {
	Field     string          `json:"field"`
	Arguments json.RawMessage `json:"arguments"`
}

// Router stores all routes and handlers
type Router map[string]Handler

// Add stores a new route with handler
func (r Router) Add(route string, function interface{}) error {
	if kind := reflect.TypeOf(function).Kind(); kind != reflect.Func {
		return fmt.Errorf("Handler is not a function, but %s", kind)
	}

	if reflect.TypeOf(function).NumIn() > 1 {
		return errors.New("Router only supports handler with one or none parameter")
	}

	if reflect.TypeOf(function).NumOut() != 2 {
		return errors.New("Router only supports handler with two return values")
	}

	r[route] = Handler{function}

	return nil
}

// Get return handler for route or error
func (r Router) Get(route string) (*Handler, error) {
	handler, found := r[route]
	if !found {
		return nil, fmt.Errorf("No handler for request found: %s", route)
	}

	return &handler, nil
}

// Serve parses the AppSync request
func (r Router) Serve(req Request) (interface{}, error) {
	handler, err := r.Get(req.Field)
	if err != nil {
		return nil, err
	}

	arguments, err := handler.Prepare(req.Arguments)
	if err != nil {
		return nil, err
	}

	return handler.Call(arguments)
}

// New returns a new Router
func New() Router {
	return Router{}
}
