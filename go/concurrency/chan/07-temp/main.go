package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

// tasks1
func tasks1() {
	outChan := make(chan int)
	errChan := make(chan error)
	finishChan := make(chan struct{})

	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(outChan chan<- int, errChan chan<- error, val int, wg *sync.WaitGroup) {
			defer wg.Done()
			time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
			if val == 30 {
				errChan <- errors.New("error in " + strconv.Itoa(val))
			}
			outChan <- val
		}(outChan, errChan, i, &wg)
	}

	go func() {
		wg.Wait()
		fmt.Println("all tasks tried")
		close(finishChan)
	}()

	var count int

Loop:
	for {
		select {
		case val := <-outChan:
			log.Println("finished:", val)
			count++
		case err := <-errChan:
			log.Println("error:", err)
			break Loop // Stop the follow-up task when an error occurs
		case <-finishChan:
			break Loop
		case <-time.After(100000 * time.Millisecond):
			break Loop
		}
	}
	fmt.Println(count)
}

// tasks2
func tasks2() {
	outChan := make(chan int)
	errIDChan := make(chan int)
	finishChan := make(chan struct{})

	numberOfTasks := 99
	wg := sync.WaitGroup{}
	wg.Add(numberOfTasks + 1)
	for i := 0; i < 100; i++ {
		go func(outChan chan<- int, errIDChan chan<- int, id int, wg *sync.WaitGroup) {
			defer wg.Done()
			time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
			value := rand.Intn(30)
			if value%2 == 0 {
				errIDChan <- id
			}
			outChan <- value
		}(outChan, errIDChan, i, &wg)
	}

	go func() {
		wg.Wait()
		fmt.Println("all tasks tried")
		close(finishChan)
	}()

	var errorIDs = []int{}
Loop:
	for {
		select {
		case val := <-outChan:
			log.Println("finished:", val)
		case err := <-errIDChan:
			errorIDs = append(errorIDs, err)
		case <-finishChan:
			break Loop
		case <-time.After(100000 * time.Millisecond):
			break Loop
		}
	}

	fmt.Println("number of tasks", numberOfTasks)
	fmt.Println("number of failed tasks", len(errorIDs))
	fmt.Println("number of successful tasks", numberOfTasks-len(errorIDs))

	if len(errorIDs) != 0 {
		for _, id := range errorIDs {
			// retry or error handling
			_ = id
		}
	}
}

func main() {
	// tasks1()
	// tasks2()
}
