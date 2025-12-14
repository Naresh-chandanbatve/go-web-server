package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
)

func serveStatic(conn net.Conn, method, path, version, root string) {
	if method != "GET" {
		writeResponse(conn, "405 Method Not Allowed", "Only GET supported")
		return
	}

	if path == "/" {
		path = "/index.html"
	}

	filePath := filepath.Join(root, path)

	data, err := os.ReadFile(filePath)
	if err != nil {
		writeResponse(conn, "404 Not Found", "File not found")
		return
	}

	headers := fmt.Sprintf(
		"%s 200 OK\r\nContent-Length: %d\r\nContent-Type: text/html\r\n\r\n",
		version,
		len(data),
	)

	conn.Write([]byte(headers))
	conn.Write(data)
}
