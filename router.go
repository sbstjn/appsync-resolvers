package router

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Request stores all information from AppSync request
type Request struct {
	Field     string          `json:"field"`
	Arguments json.RawMessage `json:"arguments"`
}

// Router stores all routes and handlers
type Router struct {
	Routes map[string]Handler
}

// Add stores a new route with handler
func (r *Router) Add(route string, function interface{}) error {
	if kind := reflect.TypeOf(function).Kind(); kind != reflect.Func {
		return fmt.Errorf("Handler is not a function, but %s", kind)
	}

	r.Routes[route] = Handler{function}
	return nil
}

// Serve parses the AppSync request
func (r *Router) Serve(req Request) (interface{}, error) {
	if handler, found := r.Routes[req.Field]; found {
		args, err := handler.Prepare(req.Arguments)
		if err != nil {
			return nil, err
		}

		return handler.Call(args)
	}

	return nil, fmt.Errorf("No handler for request found: %s", req.Field)
}

// New returns a new Router
func New() Router {
	return Router{map[string]Handler{}}
}
