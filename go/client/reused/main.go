package main

// https://golang.org/pkg/net/http/#Client

// The Client's Transport typically has internal state (cached TCP connections),
// so Clients should be reused instead of created as needed.
// Clients are safe for concurrent use by multiple goroutines.

// When following redirects,
// the Client will forward all headers set on the initial Request
// except: Authorization, WWW-Authenticate, Cookie

// If Jar is nil, the initial cookies are forwarded without change.

func main() {

}
