package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type countingWriter struct {
	w     io.Writer
	count int64
}

func (cw *countingWriter) Write(p []byte) (int, error) {
	n, err := cw.w.Write(p)
	cw.count += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := &countingWriter{w: w}
	return cw, &cw.count
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	w, count := CountingWriter(os.Stdout)
	for scanner.Scan() {
		out := append(scanner.Bytes(), '\n')
		if string(out) == "quit\n" {
			break
		}
		w.Write(out)
		fmt.Printf("Count now at: %v\n", strconv.FormatInt(*count, 10))
	}
}
