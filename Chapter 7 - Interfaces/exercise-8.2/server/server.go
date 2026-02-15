package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"path/filepath"
)

var port = flag.Int("port", 8000, "port")
var rootDir = flag.String("rootDir", "public", "roo directory")

func main() {
	flag.Parse()
	// fmt.Println("Port:", *port)

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
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	absPath, err := filepath.Abs(rootDir)
	if err != nil {
		log.Fatal(err)
	}
	ftp.Serve(ftp.NewConn(c, absPath))
}
