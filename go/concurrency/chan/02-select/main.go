package main

import (
	"fmt"
	"time"
)

// https://golang.org/ref/spec#Channel_types
// https://golang.org/ref/spec#Select_statements
// https://colobu.com/2016/04/14/Golang-Channels/
// https://blog.wu-boy.com/2019/05/handle-multiple-channel-in-15-minutes/

// If one or more of the communications can proceed, a single one that can proceed
// is chosen via a uniform pseudo-random selection.
// Otherwise, if there is a default case, that case is chosen.
// If there is no default case, the "select" statement blocks
// until at least one of the communications can proceed.

// select 沒有 default 會 blocking 等待 case 發生
// 若 blocking 又收不到 data 會造成 deadlock
func select1() {
	ch := make(chan int, 1)
	// ch <- 1
	select {
	case <-ch:
		fmt.Println("hello world")
	}
}

// 沒有 default, 沒有 timeout, 有多個 case 可選擇時
// select 會隨機選一個 case 執行
func select2() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		ch <- 1
		select {
		case <-ch:
			fmt.Println("random 01")
		case <-ch:
			fmt.Println("random 02")
		}
	}
}

// 有 default, 沒有 case 發生, 會執行 default
func select3() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		// ch <- 1
		select {
		case <-ch:
			fmt.Println("receive")
		default:
			fmt.Println("default")
		}
	}
}

// 有 default, 有 case 發生, 不會執行 default
func select4() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		ch <- 1
		select {
		case <-ch:
			fmt.Println("receive")
		default:
			fmt.Println("default")
		}
	}
}

// timeout
func select5() {
	ch := make(chan int, 1)
	defer close(ch)

	// ch <- 1

	select {
	case x, ok := <-ch:
		fmt.Println("receive", x, ok)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 01")
	}
}

// timeout
func select6() {
	timeout := make(chan bool, 1)
	defer close(timeout)

	ch := make(chan int, 1)
	defer close(ch)

	go func() {
		time.Sleep(3 * time.Second)
		timeout <- true
	}()

	// ch <- 1

	select {
	case x, ok := <-ch:
		fmt.Println("receive", x, ok)
	case <-timeout:
		fmt.Println("timeout 01")
	case <-time.After(time.Second * 5):
		fmt.Println("timeout 02")
	}
}

func main() {
	// select1()
	// select2()
	// select3()
	// select4()
	// select5()
	select6()
}
