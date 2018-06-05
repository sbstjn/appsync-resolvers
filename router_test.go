package router_test

import (
	"encoding/json"
	"errors"
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

type Response struct {
	Name string `json:"name"`
}

type ParamsRouteEmpty struct {
}

func handleRouteA(args ParamsRouteA) (interface{}, error) {
	return nil, fmt.Errorf("Nothing here in route A: %s", args.Foo)
}

func handleRouteB(args ParamsRouteB) (interface{}, error) {
	return nil, fmt.Errorf("Nothing here in route B: %s", args.Bar)
}

func handleRouteEmpty(args ParamsRouteEmpty) (interface{}, error) {
	return Response{"Frank Ocean"}, nil
}

func handleRouteWithoutParam() (interface{}, error) {
	return nil, errors.New("Nothing here in empty route")
}

func handleRouteWithMultipleParams(lorem string, foo string) (interface{}, error) {
	return nil, errors.New("Nothing here in empty route")
}

func handleRouteWithThreeReturnValue() (interface{}, interface{}, interface{}) {
	return nil, nil, nil
}

func handleRouteWithOneReturnValue() error {
	return errors.New("Nothing here in route with one return value")
}

func handleRouteWithNoReturnValue() {

}

var (
	r = router.New()
)

func TestMain(m *testing.M) {
	r.Add("fieldA", handleRouteA)
	r.Add("fieldB", handleRouteB)
	r.Add("fieldEmpty", handleRouteEmpty)
	r.Add("fieldNoParam", handleRouteWithoutParam)

	os.Exit(m.Run())
}

func TestRouteMustBeFunction(t *testing.T) {
	err := r.Add("fieldD", true)

	assert.NotNil(t, err)
	assert.Equal(t, "Handler is not a function, but bool", err.Error())
}

func TestRouteMustReturnTwoParameters(t *testing.T) {
	errNoReturn := r.Add("fieldD", handleRouteWithNoReturnValue)

	assert.NotNil(t, errNoReturn)
	assert.Equal(t, "Router only supports handler with two return values", errNoReturn.Error())

	errOneReturn := r.Add("fieldD", handleRouteWithOneReturnValue)

	assert.NotNil(t, errOneReturn)
	assert.Equal(t, "Router only supports handler with two return values", errOneReturn.Error())

	errThreeReturn := r.Add("fieldD", handleRouteWithThreeReturnValue)

	assert.NotNil(t, errThreeReturn)
	assert.Equal(t, "Router only supports handler with two return values", errThreeReturn.Error())
}

func TestRouteMustNotHaveMultipleParams(t *testing.T) {
	errMultiple := r.Add("fieldMultipleParam", handleRouteWithMultipleParams)

	assert.NotNil(t, errMultiple)
	assert.Equal(t, "Router only supports handler with one or none parameter", errMultiple.Error())
}

func TestRouteMatchA(t *testing.T) {
	routeA, err := r.Serve(router.Request{
		Field:     "fieldA",
		Arguments: json.RawMessage("{\"foo\":\"bar\"}"),
	})

	assert.Nil(t, routeA)
	assert.Equal(t, "Nothing here in route A: bar", err.Error())
}

func TestRouteMatchB(t *testing.T) {
	routeB, err := r.Serve(router.Request{
		Field:     "fieldB",
		Arguments: json.RawMessage("{\"bar\":\"foo\"}"),
	})

	assert.Nil(t, routeB)
	assert.Equal(t, "Nothing here in route B: foo", err.Error())
}

func TestRouteMatchEmpty(t *testing.T) {
	res, err := r.Serve(router.Request{
		Field:     "fieldEmpty",
		Arguments: json.RawMessage("{}"),
	})

	assert.Nil(t, err)
	assert.Equal(t, "Frank Ocean", res.(Response).Name)
}

func TestRouteWithoutParam(t *testing.T) {
	res, err := r.Serve(router.Request{
		Field:     "fieldNoParam",
		Arguments: json.RawMessage("{}"),
	})

	assert.Nil(t, res)
	assert.Equal(t, "Nothing here in empty route", err.Error())
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
