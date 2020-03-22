package basic

import (
	"fmt"
	"math/rand"
	"time"
)

func task1() <-chan int32 {
	r := make(chan int32)

	go func() {
		defer close(r)
		time.Sleep(time.Second * 3)
		r <- rand.Int31n(100)
	}()

	return r
}

func task2() <-chan int32 {
	r := make(chan int32)

	go func() {
		defer close(r)
		time.Sleep(time.Second * 3)
		r <- rand.Int31n(100)
	}()

	return r
}

func task3() <-chan int32 {
	r := make(chan int32)

	go func() {
		defer close(r)
		time.Sleep(time.Second * 3)
		r <- rand.Int31n(100)
	}()

	return r
}

// PromiseAll .
func PromiseAll() {
	ch1, ch2, ch3 := task1(), task2(), task3()
	a, b, c := <-ch1, <-ch2, <-ch3
	fmt.Println(a, b, c)
}
