// Package di provides a simple dependency injection.
package di

import (
	"fmt"
	"reflect"
	"sync"
)

var (
	// once is used to ensure that the injector is only created once
	once sync.Once

	// injectorInst is the singleton instance of the injector
	injectorInst *injector
)

// injector is the dependency injector
type injector struct {
	// dependencies is a map of registered dependencies
	dependencies sync.Map
}

// instance returns the singleton instance of the injector
func instance() *injector {
	once.Do(func() {
		injectorInst = &injector{
			dependencies: sync.Map{},
		}
	})

	return injectorInst
}

// Provide registers a dependency with the injector.
func Provide[T any](t T) {
	in := instance()
	in.dependencies.Store(key(t), t)
}

// Get retrieves a dependency from the injector.
func Get[T any]() (T, error) {
	var (
		t  T
		in = instance()
		it = reflect.TypeOf((*T)(nil)).Elem()
	)

	if it.Kind() == reflect.Interface {
		value, found := in.findImplementation(it)
		if !found {
			return t, fmt.Errorf("di: no dependency found for interface %s", it.Name())
		}

		return value.(T), nil
	}

	value, found := in.dependencies.Load(key(t))
	if !found {
		return t, fmt.Errorf("di: no dependency found for type %s", it.Name())
	}

	return value.(T), nil
}

// findImplementation searches for an implementation of the given interface type in the registered dependencies
func (in *injector) findImplementation(t reflect.Type) (value any, found bool) {
	in.dependencies.Range(func(_, v any) bool {
		if reflect.TypeOf(v).Implements(t) {
			value = v
			found = true
			return false
		}

		return true
	})

	return value, found
}

// key generates a unique key for the given type
func key[T any](t T) string {
	typ := reflect.TypeOf(t)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	return typ.PkgPath() + typ.Name()
}

// Reset clears all registered dependencies.
func Reset() {
	in := instance()
	in.dependencies = sync.Map{}
}
