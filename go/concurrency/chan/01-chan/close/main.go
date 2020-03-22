package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Player .
type Player struct {
	username     string
	successRatio int // a number between [0, 100]
}

func play(p *Player, table chan int) {
	for {
		ball, ok := <-table

		if !ok {
			// channel closed
			// the current player has won
			fmt.Printf("%s win.\n", p.username)
			return
		}

		r := rand.Intn(100)
		if r > p.successRatio {
			// failed to catch
			fmt.Printf("%s lose.\n", p.username)
			close(table)
			return
		}

		fmt.Printf("%d %+10s successfully catches the ball\n", ball, p.username)

		time.Sleep(time.Millisecond * 200)

		ball++

		table <- ball
	}
}

func main() {
	var wg sync.WaitGroup
	table := make(chan int)
	ball := 1

	wg.Add(2)
	go func() {
		play(&Player{
			username:     "getwd",
			successRatio: 80,
		}, table)
		wg.Done()
	}()

	go func() {
		time.Sleep(time.Millisecond)
		play(&Player{
			username:     "gosched",
			successRatio: 70,
		}, table)
		wg.Done()
	}()

	table <- ball

	wg.Wait()

	fmt.Println("Game Over")
}
