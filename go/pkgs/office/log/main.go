package main

import (
	"fmt"
	"io"
	"log"
	"os"
	_ "reflect"
	"runtime"
)

var (
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	errFile, err := os.OpenFile("errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("can not open errors.log", err)
	}

	Warning = log.New(os.Stdout, "Warning: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(os.Stderr, errFile), "Error: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	Warning.Println("警告警告 日暮途遠")

	Foo()
}

// Foo .
func Foo() {
	fmt.Printf("我是: %s, 誰調用我: %s\n", printMyName(), printCallerName())
	Bar()
}

// Bar .
func Bar() {
	fmt.Printf("我是: %s, 誰調用我: %s\n", printMyName(), printCallerName())
}

func printMyName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

func printCallerName() string {
	pc, _, _, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name()
}
