package di_test

import (
	"testing"

	"github.com/tirathawat/di"
)

type DependencyInterface interface {
	Foo() string
}

type Dependency struct{}

func (d *Dependency) Foo() string {
	return "foo"
}

type AnotherDependencyInterface interface {
	Bar() string
}

type AnotherDependency struct{}

func (d *AnotherDependency) Bar() string {
	return "bar"
}

func TestDependencyValue(t *testing.T) {
	di.Reset()
	d := &Dependency{}
	di.Provide(d)

	value, _ := di.Get[DependencyInterface]()

	if value != d {
		t.Errorf("Expected dependency value %v, but got %v", d, value)
	}
}

func TestErrorWhenImplementationDoesNotExist(t *testing.T) {
	di.Reset()

	_, err := di.Get[DependencyInterface]()

	if err == nil {
		t.Errorf("Expected error, but got none")
	}
}

func TestErrorWhenTypeDoesNotExist(t *testing.T) {
	di.Reset()

	_, err := di.Get[int]()

	if err == nil {
		t.Errorf("Expected error, but got none")
	}
}

func TestFindImplementation(t *testing.T) {
	di.Reset()
	d1 := &Dependency{}
	di.Provide(d1)

	value1, _ := di.Get[DependencyInterface]()

	if value1 != d1 {
		t.Errorf("Expected dependency value %v, but got %v", d1, value1)
	}
}

func TestSameInstanceForSameInterface(t *testing.T) {
	di.Reset()
	d1 := &Dependency{}
	di.Provide(d1)

	value1, _ := di.Get[DependencyInterface]()
	value2, _ := di.Get[DependencyInterface]()

	if value1 != value2 {
		t.Errorf("Expected the same instance for the same interface, but got different instances")
	}
}

func TestDifferentInstancesForDifferentInterfaces(t *testing.T) {
	di.Reset()
	di.Provide(&Dependency{})
	di.Provide(&AnotherDependency{})

	value1, _ := di.Get[AnotherDependencyInterface]()

	if _, ok := value1.(DependencyInterface); ok {
		t.Errorf("Expected different instances for different interfaces, but got the same instance")
	}
}

func TestIntValue(t *testing.T) {
	di.Reset()
	di.Provide(1)

	value, _ := di.Get[int]()

	if value != 1 {
		t.Errorf("Expected int value %v, but got %v", 1, value)
	}
}
