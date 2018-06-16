package router

import (
	"encoding/json"
	"fmt"
)

type invocation struct {
	Field     string          `json:"field"`
	Arguments json.RawMessage `json:"arguments"`
}

// Resolvers stores all routes and handlers
type Resolvers map[string]resolver

// Add stores a new resolver
func (r Resolvers) Add(field string, f interface{}) error {
	handler := resolver{f}

	if err := handler.validate(); err != nil {
		return err
	}

	r[field] = handler

	return nil
}

func (r Resolvers) get(route string) (*resolver, error) {
	handler, found := r[route]
	if !found {
		return nil, fmt.Errorf("No handler for request found: %s", route)
	}

	return &handler, nil
}

// Handle responds to the AppSync request
func (r Resolvers) Handle(req invocation) (interface{}, error) {
	var handler *resolver
	var err error

	if handler, err = r.get(req.Field); err != nil {
		return nil, err
	}

	return handler.call(req.Arguments)
}

// New returns a new lsit of Resolvers
func New() Resolvers {
	return Resolvers{}
}
