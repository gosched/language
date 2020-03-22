package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

// https://golang.org/pkg/bytes/
// https://golang.org/pkg/io/
// https://golang.org/pkg/io/ioutil/
// https://golang.org/pkg/bufio/

// https://gobyexample.com/reading-files
// https://gobyexample.com/writing-files
// https://gobyexample.com/line-filters

// Read .
func Read() {

}

// MultiRead .
func MultiRead() {
	io.MultiReader()
}

// Write .
func Write() {
	f, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return
	}
	defer f.Close()
	f.WriteString("今天天氣很好\n")
	f.Sync()

	// var w bufio.Writer

	// var w = bufio.NewWriter(os.Stdout)
	// var w = bufio.NewWriterSize(os.Stdout, 128)

	// var w = bufio.NewWriter(f)
	var w = bufio.NewWriterSize(f, 128)

	// w.ReadFrom()

	w.Write([]byte("明天天氣更好"))
	w.WriteByte(byte('\n'))
	w.WriteString("測試")
	w.WriteRune('\n')

	fmt.Println(w.Size())
	fmt.Println(w.Available())
	fmt.Println(w.Buffered())
	fmt.Println()

	fmt.Printf("%#v", w)
	fmt.Println()

	w.Flush()
	fmt.Println()

	// fmt.Printf("%#v", w)
	// fmt.Println()

	// w.Reset()
}

// ReadAndWrite .
func ReadAndWrite() {
	// bufio.NewReadWriter()
}

// ReadLines .
func ReadLines() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// scanner := bufio.NewScanner(os.Stdin)

	// lines := []string{}
	for scanner.Scan() {
		// lines = append(lines, scanner.Text())
		log.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	// log.Println(lines)
}

func main() {

}
