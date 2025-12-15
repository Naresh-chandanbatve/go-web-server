package main

import (
	"fmt"
	"net"
	"strings"
	"bufio"
)

func parseRequestLine(line string) (string, string, string) {
	parts := strings.Fields(line)
	if len(parts) != 3 {
		return "", "", ""
	}
	return parts[0], parts[1], parts[2]
}

func writeResponse(conn net.Conn, status, body string) {
	resp := "HTTP/1.1 " + status + "\r\n" +
		"Content-Length: " + fmt.Sprint(len(body)) + "\r\n" +
		"Content-Type: text/plain\r\n" +
		"\r\n" +
		body

	conn.Write([]byte(resp))
}


func readRequest(reader *bufio.Reader) (string, []string, error) {
	var reqLine string
	var headers []string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return "", nil, err
		}

		if reqLine == "" {
			reqLine = strings.TrimSpace(line)
			continue
		}

		headers = append(headers, line)

		if line == "\r\n" {
			break
		}
	}
	return reqLine, headers, nil
}