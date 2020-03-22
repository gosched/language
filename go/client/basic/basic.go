package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"net/url"
	"os"
)

// https://golang.org/pkg/net/http/#Client

// https://colobu.com/2017/04/19/go-http-redirect/

// The Client's Transport typically has internal state (cached TCP connections),
// so Clients should be reused instead of created as needed.
// Clients are safe for concurrent use by multiple goroutines.

// When following redirects,
// the Client will forward all headers set on the initial Request
// except: Authorization, WWW-Authenticate, Cookie

// If Jar is nil, the initial cookies are forwarded without change.

// CheckRedirect, Timeout, Transport, Jar

// Head .
func Head() {
	// client := &http.Client{}
	// resp, err := client.Head()
	// defer resp.Body.Close()
}

// Get .
func Get() {
	// client := &http.Client{}
	// resp, err := client.Get()
	// defer resp.Body.Close()
}

// PostForm .
func PostForm() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		log.Println(r.FormValue("hello"))
	}))
	defer ts.Close()

	client := &http.Client{}

	u, err := url.Parse(ts.URL)
	if err != nil {
		log.Fatal(err)
	}

	target := u.String()
	body := url.Values{}
	body.Set("hello", "hello world")

	resp, err := client.PostForm(target, body)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	// target := u.String()
	// contentType := "application/x -www-form-urlencoded"
	// data := url.Values{}
	// body := strings.NewReader(data.Encode())
	// resp, err := client.Post(target, contentType, body)
	// if err != nil {
	// 	log.Println(err)
	// }
	// defer resp.Body.Close()
}

// Post .
func Post() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(32 << 20)
		if err != nil {
			log.Println(err)
			return
		}

		f, fh, err := r.FormFile("test")
		if err != nil {
			log.Println(err)
		}

		log.Println(fh.Header.Get("Content-Disposition"))
		log.Println(fh.Header.Get("Content-Type"))

		name := "./temp/temp.pdf"
		file, err := os.Create(name)
		if err != nil {
			log.Println(err)
		}

		_, err = io.Copy(file, f)
		if err != nil {
			log.Println(err)
		}

		w.WriteHeader(http.StatusCreated)
	}))
	defer ts.Close()

	client := &http.Client{}

	u, err := url.Parse(ts.URL)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open("./terms.pdf")
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	target := u.String()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("test", file.Name())
	if err != nil {
		log.Println(err)
		return
	}

	_, err = io.Copy(part, file)
	if err != nil {
		log.Println(err)
		return
	}

	err = writer.Close()
	if err != nil {
		log.Println(err)
		return
	}

	contentType := writer.FormDataContentType()

	resp, err := client.Post(target, contentType, body)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
}

// Do .
func Do() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(32 << 20)
		if err != nil {
			log.Println(err)
			return
		}

		f, fh, err := r.FormFile("test")
		if err != nil {
			log.Println(err)
		}

		log.Println(fh.Header.Get("Content-Disposition"))
		log.Println(fh.Header.Get("Content-Type"))

		name := "./temp/temp.pdf"
		file, err := os.Create(name)
		if err != nil {
			log.Println(err)
		}

		_, err = io.Copy(file, f)
		if err != nil {
			log.Println(err)
		}

		w.WriteHeader(http.StatusCreated)
	}))
	defer ts.Close()

	client := &http.Client{}

	u, err := url.Parse(ts.URL)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open("./terms.pdf")
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	method := "POST"
	target := u.String()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("test", file.Name())
	if err != nil {
		log.Println(err)
		return
	}

	_, err = io.Copy(part, file)
	if err != nil {
		log.Println(err)
		return
	}

	contentType := writer.FormDataContentType()

	err = writer.Close()
	if err != nil {
		log.Println(err)
		return
	}

	req, err := http.NewRequest(method, target, body)
	req.Header.Set("Content-Type", contentType)

	resp, err := client.Do(req)
	if err != nil {

	}
	defer resp.Body.Close()
}

// WithContext .
func WithContext() {
	// ctx := context.Background()
	// req, err := http.NewRequestWithContext(ctx, method, target, body)
	// client.Do(req)
}

// WithTimeout .
func WithTimeout() {
	// client := &http.Client{
	// 	Timeout: time.Second * 5,
	// }
}

// WithoutCheckRedirect .
func WithoutCheckRedirect() {
	client := &http.Client{}

	resp, err := client.Get("https://github.com/settings/profile")
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	finalURL := resp.Request.URL.String()
	log.Println(finalURL)

	// log.Println(resp.Header)

	// data, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Println(err)
	// }

	// log.Println(string(data))
}

// WithCheckRedirect .
// the next request is not sent
// and the most recent response is returned with its body unclosed.
func WithCheckRedirect() {
	checkRedirect := func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	client := &http.Client{
		CheckRedirect: checkRedirect,
	}

	resp, err := client.Get("https://github.com/settings/profile")
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	finalURL := resp.Request.URL.String()
	log.Println(finalURL)

	// log.Println(resp.Header)

	// data, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Println(err)
	// }

	// log.Println(string(data))
}

// WithoutCookieJar .
func WithoutCookieJar() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if cookie, err := r.Cookie("Test"); err != nil {
			http.SetCookie(w, &http.Cookie{Name: "Test", Value: "Cookie Value 1"})
		} else {
			cookie.Value = "Cookie Value 2"
			http.SetCookie(w, cookie)
		}
	}))
	defer ts.Close()

	client := &http.Client{}

	u, err := url.Parse(ts.URL)
	if err != nil {
		log.Fatal(err)
	}

	resp1, err := client.Get(u.String())
	if err != nil {
		log.Println(err)
	}
	defer resp1.Body.Close()
	log.Println("After 1st request:")
	log.Printf("%v\n", resp1.Header)

	resp2, err := client.Get(u.String())
	if err != nil {
		log.Println(err)
	}
	defer resp2.Body.Close()
	log.Println("After 2nd request:")
	log.Printf("%v\n", resp2.Header)
}

// WithCookieJar .
func WithCookieJar() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if cookie, err := r.Cookie("Test"); err != nil {
			http.SetCookie(w, &http.Cookie{Name: "Test", Value: "Cookie Value 1"})
		} else {
			cookie.Value = "Cookie Value 2"
			http.SetCookie(w, cookie)
		}
	}))
	defer ts.Close()

	u, err := url.Parse(ts.URL)
	if err != nil {
		log.Fatal(err)
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{
		Jar: jar,
	}

	resp1, err := client.Get(u.String())
	if err != nil {
		log.Println(err)
	}
	defer resp1.Body.Close()
	log.Println("After 1st request:")
	log.Printf("%v\n", resp1.Header)

	resp2, err := client.Get(u.String())
	if err != nil {
		log.Println(err)
	}
	defer resp2.Body.Close()
	log.Println("After 2nd request:")
	log.Printf("%v\n", resp2.Header)

	// if _, err = client.Get(u.String()); err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println("After 1st request:")
	// for _, cookie := range jar.Cookies(u) {
	// 	log.Printf("  %s: %s\n", cookie.Name, cookie.Value)
	// }

	// if _, err = client.Get(u.String()); err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println("After 2nd request:")
	// for _, cookie := range jar.Cookies(u) {
	// 	log.Printf("  %s: %s\n", cookie.Name, cookie.Value)
	// }
}

// WithTransport .
func WithTransport() {
	// client := &http.Client{
	// 	Transport: LogRedirects{},
	// }
}

// CloseIdleConnections .
func CloseIdleConnections() {
	// client := &http.Client{}
	// client.CloseIdleConnections()
	// t := time.NewTicker(time.Hour)
	// for {
	// 	select {
	// 	case <-t.C:
	// 		client.CloseIdleConnections()
	// 	}
	// }
}
