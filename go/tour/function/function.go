package main

import (
	"fmt"
	"reflect"
)

func showTypes(elements ...interface{}) []string {
	types := []string{}
	for _, element := range elements {
		switch element.(type) {
		case bool:
			types = append(types, "bool")
		case int:
			types = append(types, "int")
		// case int32:
		// 	types = append(types, "int32")
		case int64:
			types = append(types, "int64")
		case float64:
			types = append(types, "float64")
		case complex128:
			types = append(types, "complex128")
		case rune:
			types = append(types, "rune")
		case string:
			types = append(types, "string")
		case []int:
			types = append(types, "[]int")
		default:
			types = append(types, fmt.Sprint(reflect.TypeOf(element)))
		}
	}
	fmt.Println(types)
	return types
}

func sum(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println(nums, sum)
	return sum
}

func closure() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func hello(f func() string) {
	fmt.Println("hello", f())
}

func show(f func(x, y int) int) func(x, y int) int {
	return func(x, y int) int {
		fmt.Println(x, y)
		return f(x, y)
	}
}

func add(x, y int) int {
	return x + y
}

func f3(n int) {
	if n > 1 {
		n--
		f4(n)
	}
}

func f4(n int) {
	if n%2 == 0 {
		fmt.Println(n, "Even")
	} else {
		fmt.Println(n, "Odd")
	}
	f3(n)
}

func main() {
	// Variadic function
	showTypes([]int{1, 2, 3}, 3.14, 100, false, 'z', "golang", nil)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
	sum(1, 2)
	sum(1, 2, 3)

	// First-class function
	f1 := func() {
		fmt.Println("Hello, playground")
	}

	f1()

	// Anonymous function
	func() {
		fmt.Println("Hello, playground")
	}()

	// Result only
	r := func() string {
		return "Hello, playground"
	}()

	fmt.Println(r)

	// Closure
	nextInt := closure()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	newInts := closure()
	fmt.Println(newInts())

	// Callback
	world := func() string { return "world" }
	hello(world)

	// Decorator pattern
	fmt.Println(show(add)(1, 2))
	// show(add) == func(x, y int) int
	// show(add)(1, 2) == func(1, 2 int) int

	// Pipeline pattern

	// Indirect recursion
	f3(10)
}
