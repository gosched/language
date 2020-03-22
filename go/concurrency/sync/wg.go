package explore

import (
	"fmt"
	"net/http"
	"sync"
)

// GetAllStatus .
func GetAllStatus() {
	var wg sync.WaitGroup
	results := make(chan string)
	urls := []string{
		"https://www.golang.org/",
		"https://www.google.com/",
	}
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			resp, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()
			results <- resp.Status
		}(url)
	}
	for result := range results {
		fmt.Println(result)
	}
	wg.Wait()
	close(results)
}
