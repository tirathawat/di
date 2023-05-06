# Package `di` - Simple Dependency Injection

The `di` package provides a simple dependency injection mechanism in Go. It allows you to register and retrieve dependencies in a flexible and decoupled manner.

## Installation

To install the `di` package, use the following command:

```
go get github.com/tirathawat/di
```

## Usage

Import the `di` package into your Go code:

```go
import "github.com/tirathawat/di"
```

### Registering Dependencies

To register a dependency, use the `Provide` function:

```go
di.Provide(dependency)
```

The `Provide` function accepts any value of any type as a dependency and registers it with the injector.

### Retrieving Dependencies

To retrieve a dependency from the injector, use the `Get` function:

```go
dependency, err := di.Get[DependencyType]()
```

The `Get` function retrieves a dependency of the specified type from the injector. It returns the dependency and an error if the dependency is not found.

### Interface Dependencies

If you need to retrieve a dependency based on an interface type, use the `Get` function with the interface type:

```go
dependency, err := di.Get[InterfaceType]()
```

The `Get` function will search for an implementation of the interface in the registered dependencies and return it.

### Resetting Dependencies

To clear all registered dependencies, use the `Reset` function:

```go
di.Reset()
```

The `Reset` function clears all dependencies registered in the injector, allowing you to start fresh.

## Example

Here's an example demonstrating the usage of the `di` package:

```go
package main

import (
	"fmt"

	"github.com/tirathawat/di"
)

type Database interface {
	Query(query string) string
}

type MySQLDatabase struct{}

func (d *MySQLDatabase) Query(query string) string {
	return "Result from MySQL: " + query
}

func main() {
	// Register dependencies
	di.Provide(&MySQLDatabase{})

	// Retrieve dependencies
	mysqlDB, err := di.Get[Database]()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Use the dependencies
	result := mysqlDB.Query("SELECT * FROM users")
	fmt.Println(result)

	// Reset dependencies
	di.Reset()
}
```

## Conclusion

The `di` package offers a simple dependency injection mechanism for your Go applications. By registering and retrieving dependencies through the package's functions, you can achieve decoupling and flexibility in your code.
