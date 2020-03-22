package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"time"
)

func f1(ctx context.Context, args ...string) {

}

func f2(ctx context.Context, args ...string) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second * 15)
		log.Println("from f1")
	}))
	defer ts.Close()

	req, err := http.NewRequestWithContext(ctx, "GET", ts.URL, nil)
	if err != nil {

	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}

func f3(ctx context.Context, args ...string) {
	tick := time.NewTicker(time.Second)
	ok := make(chan bool)

	go func() {
		println("go func")
	}()

loop:
	for {
		select {
		case <-tick.C:
			fmt.Println(time.Now())

		case <-ctx.Done():
			fmt.Println("cancel or timeout")
			tick.Stop()
			break loop

		case <-ok:
			fmt.Println("ok")
			tick.Stop()
			break loop
		}
	}
}

func main() {
	// ctx, cancel := context.WithCancel(context.Background())

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	// d, _ := time.ParseInLocation("2006-01-02 15:04:05", "2025-01-01 00:08:00", time.Local)
	// ctx, cancel := context.WithDeadline(context.Background(), d)

	// ctx := context.WithValue(context.Background(),,)
}
