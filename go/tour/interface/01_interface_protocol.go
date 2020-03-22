package main

import "fmt"

// https://tour.golang.org/methods/9
// https://en.wikipedia.org/wiki/Protocol_(object-oriented_programming)

// An interface type is defined as a set of method signatures.
// A value of interface type can hold any value that implements those methods.

// Interfaces are implemented implicitly
// A type implements an interface by implementing its methods.
// There is no explicit declaration of intent, no "implements" keyword.

// interface 是數個方法聲明的集合
// 關心能實現什麼, 不關心如何實現

// 類型實現了 interface 聲明的所有方法
// 就實現了該 interface, 不需顯示聲明
// 類型可以實現多個 interface

// 若要實現 interface
// 參數的數量及類型要一致, 參數的名字可以不同
// 回傳值的數量及類型要一致, 回傳值的名字可以不同

// interface 通常以 er 作為後綴

// Tester .
type Tester interface {
	test(x int) bool
}

// TempTest .
type TempTest struct{}

func (t TempTest) test(i int) (b bool) {
	// fmt.Printf("(%v, %T)\n", t, t)
	b = i%2 == 0
	return b
}

// AnotherTest .
type AnotherTest struct{}

func (t *AnotherTest) test(i int) (b bool) {
	// fmt.Printf("(%v, %T)\n", t, t)
	b = i%2 != 0
	return b
}

// User .
type User struct {
	Username string
	Age      int
}

// Overriding
func (u User) String() string {
	return fmt.Sprintf("%v (%v years)", u.Username, u.Age)
}

// HelloInterfaceProtocol .
func HelloInterfaceProtocol() {
	var t Tester

	temp := TempTest{}
	t = temp // TempTest implements test
	result1 := t.test(100)
	println(result1)

	another := &AnotherTest{}
	t = another // AnotherTest implements test
	result2 := t.test(100)
	println(result2)

	// AnotherTest does not implement Tester
	// test method has pointer receiver
	// t = *another

	// Overriding
	// &{hello 100} -> hello (100 years)
	user := &User{Username: "hello", Age: 100}
	fmt.Println(user)
}
