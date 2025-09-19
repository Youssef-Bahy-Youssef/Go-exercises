package main

import (
	"bufio"
	"fmt"
	"strings"
)

// generic counter type
type Counter int

func (c *Counter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*(c)++
	}
	err := scanner.Err()
	if err != nil {
		fmt.Println("Error:", err)
	}
	return int(*c), err
}

func (c *Counter) writeHelper(p []byte, split bufio.SplitFunc) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(split)
	for scanner.Scan() {
		*(c)++
	}
	err := scanner.Err()
	return int(*c), err
}

// ByteCounter implements io.Writer to count bytes
type ByteCounter Counter

func (c *ByteCounter) Write(p []byte) (int, error) {
	return (*Counter)(c).writeHelper(p, bufio.ScanBytes)
}

// WordCounter implements io.Writer to count words
type WordCounter Counter

func (c *WordCounter) Write(p []byte) (int, error) {
	return (*Counter)(c).writeHelper(p, bufio.ScanWords)
}

// LineCounter implements io.Writer to count lines
type LineCounter Counter

func (c *LineCounter) Write(p []byte) (int, error) {
	return (*Counter)(c).writeHelper(p, bufio.ScanLines)
}

func main() {
	// ByteCounter demo
	var bc ByteCounter
	fmt.Fprintf(&bc, "Youssef Bahy")
	fmt.Println("Byte count:", bc) // Byte count: 12

	// WordCounter demo
	var wc WordCounter
	fmt.Fprintf(&wc, "Hello my name is %s", "Youssef Bahy")
	fmt.Println("Word count:", wc) // Word count: 6

	// LineCounter demo
	var lc LineCounter
	fmt.Fprintf(&lc, "line1\nline2\nline3\n")
	fmt.Println("Line count:", lc) // Line count: 3
}
