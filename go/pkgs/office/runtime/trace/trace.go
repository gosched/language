package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

func test1() {
	for i := 0; i < 10000; i++ {
		fmt.Println("test1")
	}
}

func test2() {
	for i := 0; i < 10000; i++ {
		fmt.Println("test2")
	}
}

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	test1()
	test2()
}
