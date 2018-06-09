package router

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidHandlerFunction(t *testing.T) {
	assert.Error(t, handler{true}.validate())
	assert.Error(t, handler{412}.validate())
	assert.Error(t, handler{""}.validate())

	assert.Error(t, handler{func() {}}.validate())
	assert.Error(t, handler{func() interface{} { return nil }}.validate())
	assert.Error(t, handler{func(args struct{}) {}}.validate())
	assert.Error(t, handler{func(args struct{}) interface{} { return nil }}.validate())
	assert.Error(t, handler{func(args struct{}) (interface{}, interface{}) { return nil, nil }}.validate())

	assert.Error(t, handler{func(args string) (interface{}, error) { return nil, nil }}.validate())

	assert.Error(t, handler{func(args struct{}, param struct{}) (interface{}, error) { return nil, nil }}.validate())

	assert.Error(t, handler{func(args struct{}) (interface{}, error, error) { return nil, nil, nil }}.validate())
}

func TestValidHandlerFunction(t *testing.T) {
	assert.Nil(t, handler{func(args struct{}) (interface{}, error) { return nil, nil }}.validate())
	assert.Nil(t, handler{func(args struct{}) error { return nil }}.validate())
	assert.Nil(t, handler{func() (interface{}, error) { return nil, nil }}.validate())
	assert.Nil(t, handler{func() error { return nil }}.validate())
}

func TestArgumentFailsOnInvalidJSON(t *testing.T) {
	message := json.RawMessage("{\"\"name\": \"example\"}")
	example := handler{func(args struct {
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
	example := handler{func(args struct {
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
	example := handler{func() error {
		return nil
	}}

	data, err := example.prepare(message)

	assert.Nil(t, err)
	assert.Nil(t, data)
}
