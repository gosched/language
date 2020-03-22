package main

import (
	"net/http"
	"net/http/pprof"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)

	http.ListenAndServe("127.0.0.1:8080", mux)
}

/*
	go run test.go

	use chrome open
	http://localhost:8080/debug/pprof/

	go build .
	./continuous
	open new tab in terminal
	cd to continuous binary file's path
	go tool pprof ./continuous http://localhost:8080/debug/pprof/profile

*/
