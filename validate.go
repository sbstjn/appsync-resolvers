package resolvers

import (
	"errors"
	"fmt"
	"reflect"
)

type validateFunc func(h reflect.Type) error
type validateList []validateFunc

func (v validateList) run(handler reflect.Type) error {
	for _, validator := range v {
		if err := validator(handler); err != nil {
			return err
		}
	}

	return nil
}

var validators = validateList{
	func(h reflect.Type) error {
		if kind := h.Kind(); kind != reflect.Func {
			return fmt.Errorf("Resolver is not a function, got %s", kind)
		}

		return nil
	},
	func(h reflect.Type) error {
		if num := h.NumIn(); num > 2 {
			return fmt.Errorf("Resolver must not have more than two argument, got %v", num)
		}

		return nil
	},
	func(h reflect.Type) error {
		if h.NumIn() == 1 && h.In(0).Kind() != reflect.Struct {
			return errors.New("Resolver argument must be struct")
		}

		return nil
	},
	func(h reflect.Type) error {
		if num := h.NumOut(); num > 2 {
			return fmt.Errorf("Resolver must not have more than two return values, got %v", num)
		}

		return nil
	},
	func(h reflect.Type) error {
		if h.NumOut() < 1 {
			return errors.New("Resolver must have at least one return value")
		}

		return nil
	},
	func(h reflect.Type) error {
		if last := h.Out(h.NumOut() - 1); last.String() != "error" {
			return errors.New("Last return value must be an error")
		}

		return nil
	},
}
