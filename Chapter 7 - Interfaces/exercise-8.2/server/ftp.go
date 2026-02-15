package main

import (
	"bufio"
	"log"
	"strings"
)

type ftp struct{}

// Serve scans incoming requests for valid commands and routes them to handler functions.
func (f ftp) Serve(c *Conn) {
	c.respond(220)

	s := bufio.NewScanner(c.conn)
	for s.Scan() {
		input := strings.Fields(s.Text())
		if len(input) == 0 {
			continue
		}

		command, args := input[0], input[1]
		log.Printf("<< %s %v", command, args)

		switch command {
		case "CWD": // cd
			break
		case "LIST": // ls
			break
		}
	}
}
