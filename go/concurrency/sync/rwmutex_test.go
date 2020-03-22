package explore

import (
	"sync"
	"testing"
)

// BenchmarkAccount-8   	  240930	      4712 ns/op	       0 B/op	       0 allocs/op
// ok  	concurrency/sync/rwmutex	1.596s

func BenchmarkAccountRW(b *testing.B) {
	a := Account{name: "gosched", amount: 0}

	var wg sync.WaitGroup

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			for i := 0; i < b.N; i++ {
				a.Balance()
			}
			wg.Done()
		}()
	}

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < b.N/10000; j++ {
				a.Deposit(1000)
			}
			wg.Done()
		}()
	}

	wg.Wait()
}
