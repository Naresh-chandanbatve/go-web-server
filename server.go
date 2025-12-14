package main

import (
	"bufio"
	"net"
	"strings"
)

func handleConnection(conn net.Conn, cfg Config) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	reqLine, err := reader.ReadString('\n')
	if err != nil {
		return
	}

	method, path, version := parseRequestLine(reqLine)

	// Skip headers
	for {
		line, _ := reader.ReadString('\n')
		if line == "\r\n" {
			break
		}
	}



	serveStatic(conn, method, path, version, cfg.Root)
}
