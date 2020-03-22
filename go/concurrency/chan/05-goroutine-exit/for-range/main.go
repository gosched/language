package main

import (
	"fmt"
	"time"
)

func main() {
	in := make(chan int)

	go func(in <-chan int) {
		for x := range in {
			fmt.Printf("process %d\n", x)
		}
		fmt.Println("exit")
	}(in)

	for i := 0; i < 10; i++ {
		in <- i
	}

	close(in)

	time.Sleep(time.Second * 5)
}
