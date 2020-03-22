package main

import (
	"log"
	"net/http"
)

// LogRedirects .
type LogRedirects struct {
	Transport http.RoundTripper
}

// RoundTrip .
func (l LogRedirects) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	t := l.Transport
	if t == nil {
		t = http.DefaultTransport
	}
	resp, err = t.RoundTrip(req)
	if err != nil {
		return
	}
	switch resp.StatusCode {
	case http.StatusMovedPermanently, http.StatusFound, http.StatusSeeOther, http.StatusTemporaryRedirect:
		log.Println("Request for", req.URL, "redirected with status", resp.StatusCode)
	}
	return
}
