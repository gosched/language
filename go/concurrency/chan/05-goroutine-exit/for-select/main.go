package main

import (
	"fmt"
	"time"
)

// 某個 chan 關閉
// 透過 return 結束 goroutine
func forSelect1() {
	in := make(chan int)

	go func(in <-chan int) {
		for {
			select {
			case x, ok := <-in:
				if !ok {
					fmt.Println("exit")
					return
				}
				fmt.Printf("process %d\n", x)
			}
		}
	}(in)

	for i := 0; i < 10; i++ {
		in <- i
	}

	close(in)

	time.Sleep(time.Second * 5)
}

// 某個 chan 關閉, 設置該 chan 為 nil
// 全部的 chan 都關閉了
// 透過 return 結束 goroutine
func forSelect2() {
	in1 := make(chan int)
	in2 := make(chan int)

	go func() {
		for {
			select {
			case x1, ok := <-in1:
				if !ok {
					in1 = nil
				}
				fmt.Printf("process %d\n", x1)
			case x2, ok := <-in2:
				if !ok {
					in2 = nil
				}
				fmt.Printf("process %d\n", x2)
			}
			if in1 == nil && in2 == nil {
				fmt.Println("exit")
				return
			}
		}
	}()

	for i := 0; i < 10; i++ {
		in1 <- i
	}

	for i := 0; i < 10; i++ {
		in2 <- i * i
	}

	close(in1)

	time.Sleep(time.Second * 2)

	close(in2)

	time.Sleep(time.Second * 2)
}

func main() {
	// forSelect1()
	forSelect2()
}
