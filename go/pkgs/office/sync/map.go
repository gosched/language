package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// https://golang.org/pkg/sync/#Map
// https://colobu.com/2017/07/11/dive-into-sync-Map/

// m.Range(f func(key, value interface{}) bool)
// m.Load(key interface{}) (value interface{}, ok bool)
// m.LoadOrStore(key, value interface{}) (actual interface{}, loaded bool)
// m.Store(key, value interface{})
// m.Delete(key interface{})

// UserInfo .
type UserInfo struct {
	ID       int
	Email    string
	Username string
}

// HelloMap .
func HelloMap() {
	var m sync.Map
	var wg sync.WaitGroup
	var keys = rand.Perm(30)
	var values = rand.Perm(30)

	fmt.Println(keys)
	fmt.Println(values)
	fmt.Println()

	f := func(key, value interface{}) bool {
		//  If f returns false, range stops the iteration.
		if time.Now().Minute()%2 == 0 {
			return false
		}
		fmt.Println("show:", key, value)
		return true
	}

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			m.Store(keys[i], values[i])
			wg.Done()
		}(i)
	}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			key, value := keys[i]*keys[i], values[i]*values[i]
			fmt.Println("LoadOrStore:", key, value)
			m.LoadOrStore(key, value)
			wg.Done()
		}(i)
	}

	wg.Wait()

	fmt.Println("Range:")
	m.Range(f)
	fmt.Println()

	for i := 0; i < 20; i++ {
		m.Delete(keys[i])
	}

	fmt.Println("Range:")
	m.Range(f)
	fmt.Println()

	var dict = make(map[int]UserInfo)
	u1 := UserInfo{ID: 1, Email: "golang@gmail.com", Username: "golang"}
	u2 := UserInfo{ID: 2, Email: "javascript@gmail.com", Username: "javascript"}
	u3 := UserInfo{ID: 3, Email: "python@gmail.com", Username: "python"}
	dict[1] = u1
	dict[2] = u2

	m.Store("test", dict)
	test, _ := m.Load("test")

	fmt.Println("Range:")
	m.Range(f)
	fmt.Println()

	fmt.Println("Test map:")
	for k, v := range test.(interface{}).(map[int]UserInfo) {
		fmt.Println(k, v)
	}
	fmt.Println()

	dict[3] = u3

	fmt.Println("Range:")
	m.Range(f)
	fmt.Println()
}
