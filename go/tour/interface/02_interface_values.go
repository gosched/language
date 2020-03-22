package main

import (
	"fmt"
	"math"
)

// Interface values

// Under the hood, interface values can be thought of as a tuple of a value and a concrete type:
// (value, type)

// An interface value holds a value of a specific underlying concrete type.
// Calling a method on an interface value executes the method of the same name on its underlying type.

// Interface values with nil underlying values
// If the concrete value inside the interface itself is nil,
// the method will be called with a nil receiver.

// A nil interface value holds neither value nor concrete type.
// Calling a method on a nil interface is a run-time error
// because there is no type inside the interface tuple to indicate which concrete method to call.

// I .
type I interface {
	M()
}

// T .
type T struct {
	S string
}

// M .
func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

// F .
type F float64

// M .
func (f F) M() {
	fmt.Println(f)
}

// HelloInterfaceValue .
func HelloInterfaceValue() {
	var i, j I
	fmt.Println(i == j)            // true
	fmt.Printf("(%v, %T)\n", i, i) // (<nil>, <nil>)
	// i.M()                       // runtime error

	var t *T
	i = t                          // 賦予 nil
	fmt.Printf("(%v, %T)\n", i, i) // (<nil>, *main.T)
	i.M()                          // 可自定義賦予 nil 的結果, 未定義會報錯

	i = &T{"Hello"}                // 賦值
	fmt.Printf("(%v, %T)\n", i, i) // (&{Hello}, *main.T)
	i.M()

	i = F(math.Pi)                 // 轉型再賦值
	fmt.Printf("(%v, %T)\n", i, i) // (3.141592653589793, main.F)
	i.M()
}
