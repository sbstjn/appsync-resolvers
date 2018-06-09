package router

import (
	"errors"
	"fmt"
	"reflect"
)

type validate func(h reflect.Type) error

var validateList = []validate{
	func(h reflect.Type) error {
		if kind := h.Kind(); kind != reflect.Func {
			return fmt.Errorf("Handler is not a function, but %s", kind)
		}

		return nil
	},
	func(h reflect.Type) error {
		if h.NumIn() > 1 {
			return errors.New("Handler must not have more than one argument, got ")
		}

		return nil
	},
	func(h reflect.Type) error {
		if h.NumIn() == 1 && h.In(0).Kind() != reflect.Struct {
			return errors.New("Handler argument must be struct")
		}

		return nil
	},
	func(h reflect.Type) error {
		if h.NumOut() > 2 {
			return errors.New("Handler must not have more than two return values")
		}

		return nil
	},
	func(h reflect.Type) error {
		if h.NumOut() < 1 {
			return errors.New("Handler must have at least one return value")
		}

		return nil
	},
	func(h reflect.Type) error {
		if last := h.Out(h.NumOut() - 1); last.String() != "error" {
			return errors.New("Last or only return value must be an error")
		}

		return nil
	},
}
