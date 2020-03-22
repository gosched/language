# HTTP

## Server

- SetKeepAlivesEnabled(v bool)
- ListenAndServe() error
- ListenAndServeTLS(certFile, keyFile string) error
- Serve(l net.Listener) error
- ServeTLS(l net.Listener, certFile, keyFile string) error
- Close() error
- RegisterOnShutdown(f func())
- Shutdown(ctx context.Context) error

## Request

```
func from(w http.ResponseWriter, r *http.Request) {
	strReferrer := r.Referer()
	fmt.Println(strReferrer)
}
```

## ResponseWriter

## Header

## Cookie
- SameSite

## ServeMux

- Handle(pattern string, handler Handler)
- HandleFunc(pattern string, handler func(ResponseWriter, *Request))

## Handler

- FileServer(root FileSystem) Handler
- NotFoundHandler() Handler
- RedirectHandler(url string, code int) Handler
- StripPrefix(prefix string, h Handler) Handler
- TimeoutHandler(h Handler, dt time.Duration, msg string) Handler

### FileServer

```go
package main

import (
	"net/http"
)

func main() {
	dir := http.Dir("/usr/share")
	http.ListenAndServe(":8080", http.FileServer(dir))
	// http://127.0.0.1:8080
}
```

```go
package main

import (
	"net/http"
)

func main() {
	filesHandler := http.FileServer(http.Dir("./public"))
	http.Handle("/static/", http.StripPrefix("/static/", filesHandler))
	http.ListenAndServe(":8080", nil)
	// http://127.0.0.1:8080/static/
}
```

```go
package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	filesHandler := http.FileServer(http.Dir("./public"))
	mux.Handle("/static/", http.StripPrefix("/static/", filesHandler))
	http.ListenAndServe(":8080", mux)
	// http://127.0.0.1:8080/static/
}
```

```go
package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()
	fs := http.FileServer(http.Dir("./public"))
	mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8080", mux)
	// http://127.0.0.1:8080/static/
}
```