package resolvers

import (
	"encoding/json"
	"fmt"
)

type context struct {
	Arguments json.RawMessage `json:"arguments"`
	Source    json.RawMessage `json:"source"`
}

type invocation struct {
	Resolve string  `json:"resolve"`
	Context context `json:"context"`
}

// Resolvers stores all routes and handlers
type Resolvers map[string]resolver

// Add stores a new resolver
func (r Resolvers) Add(resolve string, f interface{}) error {
	handler := resolver{f}

	if err := handler.validate(); err != nil {
		return err
	}

	r[resolve] = handler

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

	if handler, err = r.get(req.Resolve); err != nil {
		return nil, err
	}

	if req.Context.Source != nil && string(req.Context.Source) != "null" {
		return handler.call(req.Context.Source)
	}

	return handler.call(req.Context.Arguments)
}

// New returns a new and empty list of Resolvers
func New() Resolvers {
	return Resolvers{}
}
