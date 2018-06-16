package resolvers

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatorListError(t *testing.T) {
	list := validateList{
		func(handler reflect.Type) error {
			return errors.New("Failed")
		},
	}

	var item reflect.Type
	assert.Error(t, list.run(item))
}

func TestValidatorListValid(t *testing.T) {
	list := validateList{
		func(handler reflect.Type) error {
			return nil
		},
	}

	var item reflect.Type
	assert.Nil(t, list.run(item))
}
