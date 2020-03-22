package explore

import (
	"encoding/json"
	"fmt"
)

// User .
type User struct {
	Userame string      `json:"username"`
	Age     json.Number `json:"age"`
}

// JSONNumber .
func JSONNumber() {
	data1 := []byte(`{
		"username": "google",
		"age": "99"
	}`)

	data2 := []byte(`{
		"username": "apple",
		"age": 100
	}`)

	var u1 User
	if err := json.Unmarshal(data1, &u1); err != nil {
		panic(err)
	}
	fmt.Println(u1)

	var u2 User
	if err := json.Unmarshal(data2, &u2); err != nil {
		panic(err)
	}
	fmt.Println(u2)
}
