package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/mitchellh/mapstructure"
)

// Type conversion

// Item .
type Item struct {
	ID int `json:"id"`
}

// StructToJSON .
func StructToJSON() {
	item := Item{
		ID: 1,
	}

	jsonBytes, err := json.Marshal(item)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jsonBytes))
}

// JSONToStruct .
func JSONToStruct() {
	jsonStr := `
		{
			"id":3
		}
	`
	item := Item{}
	json.Unmarshal([]byte(jsonStr), &item)
	fmt.Printf("%+v\n", item)
}

// MapToJSON .
func MapToJSON() {
	m := map[string]interface{}{"id": 5}

	jsonBytes, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jsonBytes))
}

// JSONToMap .
func JSONToMap() {
	jsonStr := `
		{
			"id":7
		}
	`

	var m map[string]interface{}

	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", m)
}

// MapToStruct .
func MapToStruct() {
	m := make(map[string]interface{})
	m["ID"] = 9

	var item Item
	err := mapstructure.Decode(m, &item)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", item)
}

// StructToMap .
func StructToMap() {
	item := Item{
		ID: 11,
	}

	t := reflect.TypeOf(item)
	v := reflect.ValueOf(item)

	var m = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		m[t.Field(i).Name] = v.Field(i).Interface()
	}

	fmt.Printf("%+v\n", m)
}
