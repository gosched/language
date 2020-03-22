package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// https://golang.org/pkg/net/
// https://golang.org/pkg/net/http/

// https://colobu.com/2015/10/09/Linux-Signals/

// CreateServer .
func CreateServer(mux http.Handler) *http.Server {
	var err error

	server := &http.Server{
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		Handler:      mux,
	}

	server.RegisterOnShutdown(func() {

	})

	// server.SetKeepAlivesEnabled(true)

	if err != nil {
		log.Println(err)
	}

	return server
}

// RunServer .
func RunServer(server *http.Server) {
	var err error
	defer log.Println(err)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGTSTP)
	go func() {
		s := <-signals
		log.Println(s)

		// syscall.SIGUSR1
		// syscall.SIGUSR2
		// syscall.SIGTERM
		// syscall.SIGTSTP

		if s.String() == "" {
			err = server.Close()
		}

		if s.String() == "" {
			ctx := context.Background()
			err = server.Shutdown(ctx)
			if err != nil {
				log.Println(err)
			}
		}
	}()

	redirect := func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r,
			"https://"+r.Host+r.URL.String(),
			http.StatusMovedPermanently)
	}

	go http.ListenAndServe(":80", http.HandlerFunc(redirect))

	err = server.ListenAndServeTLS("cert.pem", "key.pem")

	// l, err := net.Listen("tcp", ":8080")
	// err = server.Serve(l)

	// l, err := net.Listen("tcp", ":443")
	// server.ServeTLS(l, "cert.pem", "key.pem")

	log.Println("exit")
}
