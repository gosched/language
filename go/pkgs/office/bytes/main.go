package main

import (
	"bytes"
	"fmt"
)

// https://golang.org/pkg/bytes/
// https://golang.org/pkg/io/
// https://golang.org/pkg/io/ioutil/
// https://golang.org/pkg/bufio/

// type Buffer
//     func NewBuffer(buf []byte) *Buffer
//     func NewBufferString(s string) *Buffer
//     func (b *Buffer) Bytes() []byte
//     func (b *Buffer) Cap() int
//     func (b *Buffer) Grow(n int)
//     func (b *Buffer) Len() int
//     func (b *Buffer) Next(n int) []byte
//     func (b *Buffer) Read(p []byte) (n int, err error)
//     func (b *Buffer) ReadByte() (byte, error)
//     func (b *Buffer) ReadBytes(delim byte) (line []byte, err error)
//     func (b *Buffer) ReadFrom(r io.Reader) (n int64, err error)
//     func (b *Buffer) ReadRune() (r rune, size int, err error)
//     func (b *Buffer) ReadString(delim byte) (line string, err error)
//     func (b *Buffer) Reset()
//     func (b *Buffer) String() string
//     func (b *Buffer) Truncate(n int)
//     func (b *Buffer) UnreadByte() error
//     func (b *Buffer) UnreadRune() error
//     func (b *Buffer) Write(p []byte) (n int, err error)
//     func (b *Buffer) WriteByte(c byte) error
//     func (b *Buffer) WriteRune(r rune) (n int, err error)
//     func (b *Buffer) WriteString(s string) (n int, err error)
//     func (b *Buffer) WriteTo(w io.Writer) (n int64, err error)

func main() {
	var b bytes.Buffer
	fmt.Printf("%+v\n", b)
	fmt.Printf("%#v\n", b)
	fmt.Println()

	b.Write([]byte("空氣\n"))
	fmt.Printf("%+v\n", b)
	fmt.Printf("%#v\n", b)
	fmt.Println()

	var str string = "abc\n"
	b.WriteString(str)
	fmt.Printf("%+v\n", b)
	fmt.Printf("%#v\n", b)
	fmt.Println()

	fmt.Fprint(&b, "Site": https://golang.org/\n")
	fmt.Printf("%+v\n", b)
	fmt.Printf("%#v\n", b)
	fmt.Println()

	// var temp = [7]byte{}
	// n, err := b.Read(temp[:])
	// fmt.Println(n, err)
	// fmt.Printf("%+v\n", b)
	// fmt.Printf("%#v\n", b)
	// fmt.Printf("%+v\n", temp)
	// fmt.Printf("%#v\n", temp)
	// fmt.Println()

	// b.WriteTo(os.Stdout)
	// fmt.Printf("%+v\n", b)
	// fmt.Printf("%#v\n", b)
	// fmt.Println()
}
