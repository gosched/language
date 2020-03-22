package main

import (
	"net/http"

	"golang.org/x/net/http2"
)

func main() {
	var server http.Server

	server.Addr = ":8888"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello http2"))
	})

	http2.VerboseLogs = true
	http2.ConfigureServer(&server, &http2.Server{})

	server.ListenAndServe()
}
