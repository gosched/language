package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// https://golang.org/pkg/net/http/
// https://golang.org/src/net/http/client.go

// http.Client

// The client must close the response body when finished with it

// CheckRedirect specifies the policy for handling redirects.
// If CheckRedirect is not nil, the client calls it before
// following an HTTP redirect. The arguments req and via are
// the upcoming request and the requests made already, oldest
// first. If CheckRedirect returns an error, the Client's Get
// method returns both the previous Response (with its Body
// closed) and CheckRedirect's error (wrapped in a url.Error)
// instead of issuing the Request req.
// As a special case, if CheckRedirect returns ErrUseLastResponse,
// then the most recent response is returned with its body
// unclosed, along with a nil error.

// If CheckRedirect is nil, the Client uses its default policy,
// which is to stop after 10 consecutive requests.

// 301, 302, 303
// shouldRedirect = true
// includeBody = false
// if reqMethod != "GET" && reqMethod != "HEAD" {
// 	redirectMethod = "GET"
// }

// 307, 308
// shouldRedirect = true
// includeBody = true
// if resp.Header.Get("Location") == "" {
// 	shouldRedirect = false
// 	break
// }
// if ireq.GetBody == nil && ireq.outgoingLength() != 0 {
// 	shouldRedirect = false
// }

// CheckRedirect: func(req *http.Request, via []*http.Request) error {
// 	return http.ErrUseLastResponse
// },

func main() {
	resp, err := http.Get("http://example.com/")
	if err != nil {

	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

	}
	fmt.Println(string(body))
}
