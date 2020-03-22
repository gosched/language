package explore

import (
	"fmt"
	"sync"
)

func hello() {
	fmt.Println("Only once")
}

// HelloOnce .
func HelloOnce() {
	var once sync.Once

	done := make(chan bool)

	for i := 0; i < 10; i++ {
		go func() {
			once.Do(hello)
			done <- true
		}()
	}

	for i := 0; i < 10; i++ {
		<-done
	}
}
