package explore

import "testing"

func TestTime(t *testing.T) {

}

func TestHelloTimer(t *testing.T) {
	HelloTimer()
}

func TestHelloTicker(t *testing.T) {
	HelloTicker()
}

// go test -v -timeout 60s -run TestHelloTicker
