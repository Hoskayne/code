//implementing Read method using copy
//from this stackoverflow answer: https://stackoverflow.com/a/40643767
//playground link: https://play.golang.org/p/8QTECCkies

package main

import (
	"io"
	"io/ioutil"
	"log"
)

type Reader struct {
	data      []byte
	readIndex int64
}

func NewReader(toRead string) *Reader {
	return &Reader{data: []byte(toRead)}
}

func (r *Reader) Read(p []byte) (n int, err error) {
	if r.readIndex >= int64(len(r.data)) {
		err = io.EOF
		return
	}

	n = copy(p, r.data[r.readIndex:])
	r.readIndex += int64(n)
	return
}

func main() {
	M := NewReader("test")
	stuff, _ := ioutil.ReadAll(M)
	log.Printf("%s", stuff)

	M2 := NewReader("abcde")
	buf := make([]byte, 2)
	for i := 0; i < 4; i++ {
		n, err := M2.Read(buf)
		log.Println(n, err, buf, string(buf))
	}
}
