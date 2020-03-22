package explore

import (
	"fmt"
	"sync"
)

func producer(nums ...int) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		for _, n := range nums {
			output <- n
		}
	}()
	return output
}

func square(input <-chan int) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		for n := range input {
			output <- n * n
		}
	}()
	return output
}

func merge(consumer ...<-chan int) <-chan int {
	output := make(chan int, 100)
	// output := make(chan int)

	var wg sync.WaitGroup

	collect := func(input <-chan int) {
		defer wg.Done()
		for n := range input {
			output <- n
		}
	}

	wg.Add(len(consumer))

	// fan in
	for _, c := range consumer {
		go collect(c)
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}

// Fan .
func Fan() {
	// items := rand.Perm(20)
	// source := producer(items...)

	source := producer(0, 1, 2, 3)

	// fan out
	c1 := square(source)
	c2 := square(source)
	c3 := square(source)

	// consumer
	for result := range merge(c1, c2, c3) {
		fmt.Printf("%3d ", result)
	}

	fmt.Println()
}
