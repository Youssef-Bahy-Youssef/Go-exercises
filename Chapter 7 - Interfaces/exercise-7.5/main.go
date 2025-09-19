package main

import (
	"fmt"
	"io"
	"strings"
)

// limitReader wraps an io.Reader and stops after n bytes
type limitReader struct {
	r    io.Reader
	n    int64
	read int64
}

func (l *limitReader) Read(p []byte) (int, error) {
	if l.read >= l.n {
		return 0, io.EOF
	}

	// Limit the slice to read only remaining bytes
	if int64(len(p)) > l.n-l.read {
		p = p[:l.n-l.read]
	}
	n, err := l.r.Read(p)
	l.read += int64(n)
	return n, err
}

// LimitReader returns a new Reader that reads from r but stops after n bytes
func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitReader{r: r, n: n}
}

func main() {
	name := "Youssef Bahy Youssef"
	l := LimitReader(strings.NewReader(name), int64(15))

	buf := make([]byte, 8)

	n, err := l.Read(buf)
	fmt.Println("n =", n, "data =", string(buf[:n]), "err =", err)

	n, err = l.Read(buf)
	fmt.Println("n =", n, "data =", string(buf[:n]), "err =", err)

	n, err = l.Read(buf)
	fmt.Println("n =", n, "data =", string(buf[:n]), "err =", err)
}
