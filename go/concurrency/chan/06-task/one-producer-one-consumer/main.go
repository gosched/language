package main

import (
	"fmt"
	"time"
)

func producer(n int) <-chan int {
	out := make(chan int)
	go func() {
		defer func() {
			close(out)
			fmt.Println("producer exit")
		}()

		for i := 1; i <= n; i++ {
			fmt.Printf("send %d\n", i)
			out <- i
			time.Sleep(time.Millisecond)
		}
	}()
	return out
}

func consumer(in <-chan int) <-chan struct{} {
	finish := make(chan struct{})

	go func() {
		defer func() {
			fmt.Println("consumer exit")
			finish <- struct{}{}
			close(finish)
		}()

		for {
			select {
			case x, ok := <-in:
				if !ok {
					return
				}
				fmt.Printf("process %d\n", x)
			}
		}
	}()

	return finish
}

func main() {
	items := producer(5)
	finish := consumer(items)
	<-finish
	fmt.Println("main exit")
}
