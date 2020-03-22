package explore

import (
	"fmt"
	"time"
)

func producer(n int) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		for i := 0; i < n; i++ {
			output <- i
		}
	}()
	return output
}

func square(inCh <-chan int) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		for n := range inCh {
			output <- n * n
			time.Sleep(time.Second)
		}
	}()
	return output
}

// Basic .
func Basic() {
	intput := producer(4)
	output := square(intput)

	for v := range output {
		fmt.Println(v)
	}
}
