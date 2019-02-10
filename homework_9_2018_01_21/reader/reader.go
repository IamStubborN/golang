package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct{ r io.Reader }

func (t rot13Reader) Read(b []byte) (int, error) {
	n, e := t.r.Read(b)
	for i := 0; i < n; i++ {
		switch {
		case b[i] >= 110:
			b[i] -= 13
		case b[i] >= 97:
			b[i] += 13
		case b[i] >= 78:
			b[i] -= 13
		case b[i] >= 65:
			b[i] += 13
		}
	}
	return n, e
}

func main() {
	s := strings.NewReader("Penpx pbqr")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
