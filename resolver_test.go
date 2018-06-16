package router

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidResolverFunction(t *testing.T) {
	assert.Error(t, resolver{true}.validate())
	assert.Error(t, resolver{412}.validate())
	assert.Error(t, resolver{""}.validate())

	assert.Error(t, resolver{func() {}}.validate())
	assert.Error(t, resolver{func() interface{} { return nil }}.validate())
	assert.Error(t, resolver{func(args struct{}) {}}.validate())
	assert.Error(t, resolver{func(args struct{}) interface{} { return nil }}.validate())
	assert.Error(t, resolver{func(args struct{}) (interface{}, interface{}) { return nil, nil }}.validate())

	assert.Error(t, resolver{func(args string) (interface{}, error) { return nil, nil }}.validate())

	assert.Error(t, resolver{func(args struct{}, param struct{}) (interface{}, error) { return nil, nil }}.validate())

	assert.Error(t, resolver{func(args struct{}) (interface{}, error, error) { return nil, nil, nil }}.validate())
}

func TestValidResolverFunction(t *testing.T) {
	assert.Nil(t, resolver{func(args struct{}) (interface{}, error) { return nil, nil }}.validate())
	assert.Nil(t, resolver{func(args struct{}) error { return nil }}.validate())
	assert.Nil(t, resolver{func() (interface{}, error) { return nil, nil }}.validate())
	assert.Nil(t, resolver{func() error { return nil }}.validate())
}

func TestArgumentFailsOnInvalidJSON(t *testing.T) {
	message := json.RawMessage("{\"\"name\": \"example\"}")
	example := resolver{func(args struct {
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
	example := resolver{func(args struct {
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
	example := resolver{func() error {
		return nil
	}}

	data, err := example.prepare(message)

	assert.Nil(t, err)
	assert.Nil(t, data)
}
