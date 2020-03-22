package main

import "fmt"

// The interface type that specifies zero methods
// is known as the empty interface

// An empty interface may hold values of any type.
// Every type implements at least zero methods.

// Empty interfaces are used by code that handles values of unknown type.

// HelloInterfaceEmpty .
func HelloInterfaceEmpty() {
	var i interface{}
	fmt.Printf("(%v, %T)\n", i, i)

	i = 42
	fmt.Printf("(%v, %T)\n", i, i)

	i = "hello"
	fmt.Printf("(%v, %T)\n", i, i)
}
