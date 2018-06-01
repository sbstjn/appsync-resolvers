package router_test

import (
	"encoding/json"
	"errors"
	"os"
	"testing"

	router "github.com/sbstjn/appsync-router"
	"github.com/stretchr/testify/assert"
)

func handleRouteA(req json.RawMessage) (interface{}, error) {
	return nil, errors.New("Nothing here in route A")
}

func handleRouteB(req json.RawMessage) (interface{}, error) {
	return nil, errors.New("Nothing here in route B")
}

var (
	r = router.New()
)

func TestMain(m *testing.M) {
	r.Add("fieldA", handleRouteA)
	r.Add("fieldB", handleRouteB)

	os.Exit(m.Run())
}

func TestRouteMatch(t *testing.T) {
	routeA, err := r.Serve(router.Request{
		Field:     "fieldA",
		Arguments: json.RawMessage(""),
	})

	assert.Nil(t, routeA)
	assert.Equal(t, errors.New("Nothing here in route A"), err)

	routeB, err := r.Serve(router.Request{
		Field:     "fieldB",
		Arguments: json.RawMessage(""),
	})

	assert.Nil(t, routeB)
	assert.Equal(t, errors.New("Nothing here in route B"), err)
}

func TestRouteMiss(t *testing.T) {
	routeC, err := r.Serve(router.Request{
		Field:     "fieldC",
		Arguments: json.RawMessage(""),
	})

	assert.Nil(t, routeC)
	assert.Equal(t, errors.New("Unable to handle request: fieldC"), err)
}
