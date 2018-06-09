package router

import (
	"encoding/json"
	"fmt"
)

type request struct {
	Field     string          `json:"field"`
	Arguments json.RawMessage `json:"arguments"`
}

// Router stores all routes and handlers
type Router map[string]Handler

// Add stores a new route with handler
func (r Router) Add(route string, function interface{}) error {
	handler := Handler{function}

	if err := handler.validate(); err != nil {
		return err
	}

	r[route] = handler

	return nil
}

func (r Router) get(route string) (*Handler, error) {
	handler, found := r[route]
	if !found {
		return nil, fmt.Errorf("No handler for request found: %s", route)
	}

	return &handler, nil
}

// Handle responds to the AppSync request
func (r Router) Handle(req request) (interface{}, error) {
	var handler *Handler
	var err error

	if handler, err = r.get(req.Field); err != nil {
		return nil, err
	}

	return handler.call(req.Arguments)
}

// New returns a new Router
func New() Router {
	return Router{}
}
