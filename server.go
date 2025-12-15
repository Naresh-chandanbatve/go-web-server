package main

import (
	"bufio"
	"net"
	"strings"
	"io"
)

func handleConnection(conn net.Conn, cfg Config) {

	reader := bufio.NewReader(conn)

	reqLine, err := reader.ReadString('\n')
	if err != nil {
		conn.Close()
		return
	}

	method, path, _ := parseRequestLine(reqLine)

	reader = bufio.NewReader(
		io.MultiReader(strings.NewReader(reqLine), reader),
	)

	for prefix, backend := range cfg.Proxies {
		if strings.HasPrefix(path, prefix) {
			reverseProxy(conn, reader, backend)
			return
		}
	}

	if cfg.Root != "" {
		defer conn.Close()
		serveStatic(conn, method, path, "HTTP/1.1", cfg.Root)
		return
	}

	defer conn.Close() 
	writeResponse(conn, "404 Not Found", "Not Found")
	conn.Close()
}


