package router

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidHandlerFunction(t *testing.T) {
	assert.Error(t, Handler{true}.validate())
	assert.Error(t, Handler{412}.validate())
	assert.Error(t, Handler{""}.validate())

	assert.Error(t, Handler{func() {}}.validate())
	assert.Error(t, Handler{func() interface{} { return nil }}.validate())
	assert.Error(t, Handler{func(args struct{}) {}}.validate())
	assert.Error(t, Handler{func(args struct{}) interface{} { return nil }}.validate())
	assert.Error(t, Handler{func(args struct{}) (interface{}, interface{}) { return nil, nil }}.validate())

	assert.Error(t, Handler{func(args string) (interface{}, error) { return nil, nil }}.validate())

	assert.Error(t, Handler{func(args struct{}, param struct{}) (interface{}, error) { return nil, nil }}.validate())

	assert.Error(t, Handler{func(args struct{}) (interface{}, error, error) { return nil, nil, nil }}.validate())
}

func TestValidHandlerFunction(t *testing.T) {
	assert.Nil(t, Handler{func(args struct{}) (interface{}, error) { return nil, nil }}.validate())
	assert.Nil(t, Handler{func(args struct{}) error { return nil }}.validate())
	assert.Nil(t, Handler{func() (interface{}, error) { return nil, nil }}.validate())
	assert.Nil(t, Handler{func() error { return nil }}.validate())
}

func TestArgumentFailsOnInvalidJSON(t *testing.T) {
	message := json.RawMessage("{\"\"name\": \"example\"}")
	example := Handler{func(args struct {
		Name string `json:"name"`
	}) error {
		return nil
	}}

	data, err := example.prepare(message)

	assert.Error(t, err)
	assert.Nil(t, data)
}

func TestArgumentWorkForValidJSON(t *testing.T) {
	message := json.RawMessage("{\"name\": \"example\"}")
	example := Handler{func(args struct {
		Name string `json:"name"`
	}) error {
		return nil
	}}

	data, err := example.prepare(message)

	assert.Nil(t, err)
	assert.NotNil(t, data)
}

func TestArgumentWorkWithoutHandlerArguments(t *testing.T) {
	message := json.RawMessage("{\"name\": \"example\"}")
	example := Handler{func() error {
		return nil
	}}

	data, err := example.prepare(message)

	assert.Nil(t, err)
	assert.Nil(t, data)
}
