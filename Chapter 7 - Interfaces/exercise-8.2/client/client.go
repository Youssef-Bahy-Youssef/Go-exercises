package main

import (
	"flag"
	"fmt"
	"net"
)

var port = flag.Int("port", 8000, "port number of the server")

func main() {
	flag.Parse()

	conn, err := net.Dial("tcp", "localhost"+fmt.Sprint(*port))
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	defer conn.Close()
	// make action

}
