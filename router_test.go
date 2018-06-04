package router_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	router "github.com/sbstjn/appsync-router"
	"github.com/stretchr/testify/assert"
)

type ParamsRouteA struct {
	Foo string `json:"foo"`
}

type ParamsRouteB struct {
	Bar string `json:"bar"`
}

func handleRouteA(args ParamsRouteA) (interface{}, error) {
	return nil, fmt.Errorf("Nothing here in route A: %s", args.Foo)
}

func handleRouteB(args ParamsRouteB) (interface{}, error) {
	return nil, fmt.Errorf("Nothing here in route B: %s", args.Bar)
}

var (
	r = router.New()
)

func TestMain(m *testing.M) {
	r.Add("fieldA", handleRouteA)
	r.Add("fieldB", handleRouteB)

	os.Exit(m.Run())
}

func TestRouteMustBeFunction(t *testing.T) {
	err := r.Add("fieldD", true)

	assert.NotNil(t, err)
	assert.Equal(t, "Handler is not a function, but bool", err.Error())
}

func TestRouteMatch(t *testing.T) {
	routeA, err := r.Serve(router.Request{
		Field:     "fieldA",
		Arguments: json.RawMessage("{\"foo\":\"bar\"}"),
	})

	assert.Nil(t, routeA)
	assert.Equal(t, "Nothing here in route A: bar", err.Error())

	routeB, err := r.Serve(router.Request{
		Field:     "fieldB",
		Arguments: json.RawMessage("{\"bar\":\"foo\"}"),
	})

	assert.Nil(t, routeB)
	assert.Equal(t, "Nothing here in route B: foo", err.Error())
}

func TestRouteMiss(t *testing.T) {
	routeC, err := r.Serve(router.Request{
		Field:     "fieldC",
		Arguments: json.RawMessage(""),
	})

	assert.Nil(t, routeC)
	assert.Equal(t, "No handler for request found: fieldC", err.Error())
}

func TestRouteMatchWithInvalidPayload(t *testing.T) {
	routeA, err := r.Serve(router.Request{
		Field:     "fieldA",
		Arguments: json.RawMessage(""),
	})

	assert.Nil(t, routeA)
	assert.Equal(t, "unexpected end of JSON input", err.Error())
}
