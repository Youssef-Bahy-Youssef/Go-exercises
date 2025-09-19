package main

import (
	"bufio"
	"fmt"
	"strings"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*(c)++
	}
	err := scanner.Err()
	return int(*c), err
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		*(c)++
	}
	err := scanner.Err()
	return int(*c), err
}

func main() {
	// WordCounter demo
	var name = "Youssef Bahy Youssef"
	var wc WordCounter
	wc.Write([]byte(name))
	fmt.Println("Word count:", wc) // Word count: 3

	wc = 0 // reset the counter
	fmt.Fprintf(&wc, "hello My name:, %s", name)
	fmt.Println("Word count:", wc) // Word count: 6

	// LineCounter demo
	var paragraph = "Youssef\nYoussef\nYoussef\n"
	var lc LineCounter

	lc.Write([]byte(paragraph))
	fmt.Println("Line count:", lc) // Line count: 3

	lc = 0 // reset the counter
	fmt.Fprintf(&lc, paragraph)
	fmt.Println("Line count:", lc) // Line count: 3
}
