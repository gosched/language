package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

func a() {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			fmt.Println(i * j)
		}
	}
}

func main() {
	f, err := os.Create("cpu-profile.prof")
	if err != nil {
		log.Fatal(err)
	}
	f.Close()

	//pprof.WriteHeapProfile(f)
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	for i := 0; i < 10; i++ {
		go a()
	}

}

/*
	go build .
	go tool pprof ./discontinuous cpu-profile.prof
*/
