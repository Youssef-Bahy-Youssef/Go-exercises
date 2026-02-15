package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"time"
)

var port = flag.Int("port", 8000, "port")

func main() {
	flag.Parse()
	ln, err := net.Listen("tcp", "localhost:"+fmt.Sprint(*port))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		go handleConn(conn)

		time.Sleep(time.Hour)
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)

	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)

	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), time.Second)
	}
	io.Copy(os.Stdout, c)
	c.Close()
}
