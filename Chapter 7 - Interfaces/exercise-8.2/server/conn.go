package main

import "net"

type Conn struct {
	conn    net.Conn
	rootDir string
	workDir string
}

func NewConn(conn net.Conn, rootDir string) *Conn {
	return &Conn{
		conn:    conn,
		rootDir: rootDir,
		workDir: "/",
	}
}
