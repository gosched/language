package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

// UserInfo .
type UserInfo struct {
	ID        int
	Name      string
	LastName  string
	FirstName string
}

func init() {

}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", logging(handler0))
	mux.HandleFunc("/hello", handler0)
	mux.HandleFunc("/query", handler1)
	mux.HandleFunc("/login", handler2)
	mux.HandleFunc("/book", handler3)
	mux.HandleFunc("/books", handler4)
	mux.HandleFunc("/report", handler5)
	mux.HandleFunc("/redirect", redirect)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {

	}
}

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("URL.Path:", r.URL.Path)  // 當前網址
		log.Println("Referrer:", r.Referer()) // 從哪連過來的
		f(w, r)
	}
}

func handler0(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

// http://localhost:8080/query/?hello=world&hello=today&message=123
func handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.RawQuery)
	fmt.Println(r.URL.String())

	q := r.URL.Query()
	if q["hello"] != nil {
		hello := q["hello"][0]
		fmt.Println(hello)
	}

	values, err := url.ParseQuery(r.URL.EscapedPath())
	if err != nil {

	}

	message := values.Get("message")
	fmt.Println(message)
}

// application/x-www-form-urlencoded
func handler2(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	uid := r.Form["uid"]
	number := r.PostFormValue("number")
	numbers := r.PostForm["numbers"]
	fmt.Println(uid)
	fmt.Println(number)
	fmt.Println(numbers)
}

// multipart/form-data
func handler3(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20) // 32 MB
	f, header, err := r.FormFile("uploadfile")
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer f.Close()

	temp, err := os.Create("./storage/upload/" + header.Filename)
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer temp.Close()

	_, err = io.Copy(temp, f)
	if err != nil {
		log.Fatalln(err)
		return
	}
}

// multipart/form-data
func handler4(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	files := r.MultipartForm.File["files"]
	for _, file := range files {
		f, err := file.Open()
		if err != nil {
			log.Fatalln(err)
		}
		defer f.Close()

		temp, err := os.Create("./storage/upload/" + file.Filename)
		if err != nil {
			log.Fatalln(err)
		}
		defer temp.Close()

		_, err = io.Copy(temp, f)
		if err != nil {
			log.Fatalln(err)
			return
		}
	}
}

// http POST http://localhost:8080/report API-Token:123 name=gosched
// application/json
func handler5(w http.ResponseWriter, r *http.Request) {
	user := &UserInfo{}

	// token := r.Header.Get("API-Token")

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	err = json.Unmarshal(data, user)
	if err != nil {
		log.Println(err)
		return
	}

	b, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func redirect(w http.ResponseWriter, r *http.Request) {
	// http.Redirect(w, r, "/", 301) // 永久轉移
	http.Redirect(w, r, "/", 302) // 暫時轉移
}
