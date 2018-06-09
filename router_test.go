package router

import (
	"encoding/json"
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	r = New()
)

func createHandleResponse(route string) (interface{}, error) {
	return r.Handle(request{
		Field:     route,
		Arguments: json.RawMessage("{}"),
	})
}

func TestMain(m *testing.M) {
	r.Add("routeWithArgumentsAndSingleReturnThatReturnsError", func(args struct{ Foo bool }) error {
		return errors.New("Example Error")
	})
	r.Add("routeWithArgumentsAndSingleReturnThatReturnsNil", func(args struct{ Foo bool }) error {
		return nil
	})
	r.Add("routeWithArgumentsAndTwoReturnsThatReturnsError", func(args struct{ Foo bool }) (interface{}, error) {
		return nil, errors.New("Example Error")
	})
	r.Add("routeWithArgumentsAndTwoReturnsThatReturnsData", func(args struct{ Foo bool }) (interface{}, error) {
		type data struct {
			Foo bool
		}
		return data{true}, nil
	})
	r.Add("routeWithArgumentsAndTwoReturnsThatReturnsNil", func(args struct{ Foo bool }) (interface{}, error) {
		return nil, nil
	})

	r.Add("validWithoutParameter", func() error { return nil })
	r.Add("validWithOneParameter", func(args struct {
		Name string `json:"name"`
	}) error {
		return nil
	})

	os.Exit(m.Run())
}

func TestRouteValidationOnAdd(t *testing.T) {
	assert.Nil(t, r.Add("validWithoutParameter-again", func() error { return nil }))
	assert.Error(t, r.Add("invalidHandler", func(args bool) error { return nil }))
}

func TestRegisteredRoutes(t *testing.T) {
	var data interface{}
	var err error

	data, err = createHandleResponse("routeWithArgumentsAndSingleReturnThatReturnsError")
	assert.Nil(t, data)
	assert.Error(t, err)

	data, err = createHandleResponse("routeWithArgumentsAndSingleReturnThatReturnsNil")
	assert.Nil(t, data)
	assert.Nil(t, err)

	data, err = createHandleResponse("routeWithArgumentsAndTwoReturnsThatReturnsError")
	assert.Nil(t, data)
	assert.Error(t, err)

	data, err = createHandleResponse("routeWithArgumentsAndTwoReturnsThatReturnsData")
	assert.NotNil(t, data)
	assert.Nil(t, err)

	data, err = createHandleResponse("routeWithArgumentsAndTwoReturnsThatReturnsNil")
	assert.Nil(t, data)
	assert.Nil(t, err)
}

func TestRouteMiss(t *testing.T) {
	undefiendRoute, err := r.Handle(request{
		Field:     "invalid",
		Arguments: json.RawMessage("{}"),
	})

	assert.Nil(t, undefiendRoute)
	assert.Equal(t, "No handler for request found: invalid", err.Error())
}

func TestRouteMatchWithInvalidPayload(t *testing.T) {
	validRoute, err := r.Handle(request{
		Field:     "validWithOneParameter",
		Arguments: json.RawMessage("{}}"),
	})

	assert.Nil(t, validRoute)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid character '}' after top-level value", err.Error())
}
