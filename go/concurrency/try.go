package main

import (
	"fmt"
	"sync"
)

// Print100 .
func Print100() {
	var wg sync.WaitGroup

	wg.Add(2)

	control := make(chan int)

	go func() {
		for i := 1; i <= 100; i++ {
			control <- i
			if i%2 != 0 {
				fmt.Println(i)
			}
		}
		wg.Done()
	}()

	go func() {
		for i := 1; i <= 100; i++ {
			<-control
			if i%2 == 0 {
				fmt.Println(i)
			}
		}
		wg.Done()
	}()

	wg.Wait()

	close(control)
}

func main() {

}
