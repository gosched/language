package main

import (
	"fmt"
	"strconv"
)

// https://en.wikipedia.org/wiki/Data_type
// https://en.wikipedia.org/wiki/Type_inference
// https://en.wikipedia.org/wiki/Type_conversion

// User .
type User struct{}

// Users .
type Users *[]User

func typePrintf(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func typeAssertion() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// f = i.(float64) // panic
	// fmt.Println(f)
}

func typeSwitch(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

// IntAndString .
func IntAndString() {
	number, err := strconv.Atoi("1")
	if err != nil {

	}
	str := strconv.Itoa(number)
	fmt.Println(str)
}

func main() {
	StructToJSON()
	JSONToStruct()

	MapToJSON()
	JSONToMap()

	MapToStruct()
	StructToMap()
}
