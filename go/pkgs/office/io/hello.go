package main

import (
	"io"
	"log"
	"sync"
)

// HelloPipe .
func HelloPipe() {
	r1, w1 := io.Pipe()
	r2, w2 := io.Pipe()

	w := io.MultiWriter(w1, w2)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		temp := make([]byte, 256)
		_, err := r1.Read(temp)
		if err != nil {

		}
		log.Println(string(temp))
		wg.Done()
	}()

	go func() {
		temp := make([]byte, 256)
		_, err := r2.Read(temp)
		if err != nil {

		}
		log.Println(string(temp))
		wg.Done()
	}()

	w.Write([]byte("hello"))

	wg.Wait()

}

func main() {
	HelloPipe()
}
