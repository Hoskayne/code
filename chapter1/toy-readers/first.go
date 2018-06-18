// implementation of read method
// from this stackoverflow answer: https://stackoverflow.com/a/28202063
// playground link: https://play.golang.org/p/9BbS54d8pb
// also took some snippets from link below
// https://medium.com/@matryer/golang-advent-calendar-day-seventeen-io-reader-in-depth-6f744bb4320b

package main

import (
	"fmt"
	"io"
	"io/ioutil"
)

type Reader struct {
	data []byte
}

func NewReader(toRead string) *Reader {
	return &Reader{[]byte(toRead)}
}

func (r Reader) eof() bool {
	return len(r.data) == 0
}

func (r *Reader) readByte() byte {
	//this function assumes that eof check was done before
	b := r.data[0]
	r.data = r.data[1:]
	return b
}

func (r *Reader) Read(p []byte) (n int, err error) {
	if r.eof() {
		err = io.EOF
		return
	}

	if l := len(p); l > 0 {
		for n < l {
			p[n] = r.readByte()
			n++
			if r.eof() {
				r.data = []byte{}
				break
			}
		}
	}
	return
} //now type Reader implements io.Reader interface

func main() {

	//now I can do cool things with it

	//reading directly
	r := NewReader("length?")
	p := make([]byte, 4)
	n, _ := r.Read(p)
	fmt.Printf("Length of p: ")
	fmt.Print(n)

	//read everything from Reader and get raw byte data
	thingy := NewReader("Hello hello goodbye goodbye")
	stuff, _ := ioutil.ReadAll(thingy)
	fmt.Printf("\n %s", stuff)

	//can also easily decode JSON and other cool things

	//some hypothetical function
	//func dumbHashAlgo (s string) (string error)
	//could become func dumbHashAlgo (r io.Reader) io.Reader
	//and people can use it with any type that implements io.Reader
	//e.g. strings (as in this example)
	//files, buffers, web requests, anything
}
