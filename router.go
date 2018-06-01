package router

import (
	"encoding/json"
	"fmt"
)

// Request stores all information from AppSync request
type Request struct {
	Field     string          `json:"field"`
	Arguments json.RawMessage `json:"arguments"`
}

// RouteHandler defines the function to handle the request
type RouteHandler func(req json.RawMessage) (interface{}, error)

// Router stores all routes and handlers
type Router struct {
	Routes map[string]RouteHandler
}

// Add stores a new route with handler
func (r *Router) Add(route string, handler RouteHandler) {
	r.Routes[route] = handler
}

// Serve parses the AppSync request
func (r *Router) Serve(req Request) (interface{}, error) {
	if handler, found := r.Routes[req.Field]; found {
		return handler(req.Arguments)
	}

	return nil, fmt.Errorf("Unable to handle request: %s", req.Field)
}

// New returns a new Router
func New() Router {
	return Router{
		map[string]RouteHandler{},
	}
}
