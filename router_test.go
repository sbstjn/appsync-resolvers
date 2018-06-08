package router_test

import (
	"encoding/json"
	"os"
	"testing"

	router "github.com/sbstjn/appsync-router"
	"github.com/stretchr/testify/assert"
)

var (
	r = router.New()
)

func TestMain(m *testing.M) {
	r.Add("validWithoutParameter", func() error { return nil })
	r.Add("validWithOneParameter", func(args struct {
		Name string `json:"name"`
	}) error {
		return nil
	})

	os.Exit(m.Run())
}

func TestRouteMiss(t *testing.T) {
	undefiendRoute, err := r.Handle(router.Request{
		Field:     "invalid",
		Arguments: json.RawMessage("{}"),
	})

	assert.Nil(t, undefiendRoute)
	assert.Equal(t, "No handler for request found: invalid", err.Error())
}

func TestRouteMatchWithInvalidPayload(t *testing.T) {
	validRoute, err := r.Handle(router.Request{
		Field:     "validWithOneParameter",
		Arguments: json.RawMessage("{}}"),
	})

	assert.Nil(t, validRoute)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid character '}' after top-level value", err.Error())
}
