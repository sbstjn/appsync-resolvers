package resolvers

import (
	"fmt"
	"log"
	"reflect"
)

// Repository stores all resolvers
type Repository map[string]resolver

// Add stores a new resolver
func (r Repository) Add(resolve string, handler interface{}) error {
	err := validators.run(reflect.TypeOf(handler))

	if err == nil {
		r[resolve] = resolver{handler}
	}

	return err
}

// Handle responds to the AppSync request
func (r Repository) Handle(in invocation) (interface{}, error) {
	handler, found := r[in.Resolve]

	log.Printf("calling handler with arguments: %v\n", in)
	if found {
		return handler.call(in.payload(), in.identity())
	}

	return nil, fmt.Errorf("No resolver found: %s", in.Resolve)
}
