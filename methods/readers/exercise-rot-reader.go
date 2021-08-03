package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r13 rot13Reader) Read(bytes []byte) (n int, err error) {
	n, err = r13.r.Read(bytes)
	for i := range bytes {
		if bytes[i] >= 'A' && bytes[i] < 'Z' {
			bytes[i] = 65 + (((bytes[i] - 65) + 13) % 26)
		} else if bytes[i] >= 'a' && bytes[i] <= 'z' {
			bytes[i] = 97 + (((bytes[i] - 97) + 13) % 26)
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
