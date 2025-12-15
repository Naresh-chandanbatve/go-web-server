package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
)


func reverseProxy(client net.Conn, reader *bufio.Reader, backend string) {
	defer client.Close()
	log.Println("Proxying request to backend:", backend)

	reqLine, headers, err := readRequest(reader)
	if err != nil {
		log.Println("read request error:", err)
		writeResponse(client, "400 Bad Request", "Bad Request")
		return
	}

	contentLen := 0
	for _, h := range headers {
		if strings.HasPrefix(strings.ToLower(h), "content-length:") {
			v := strings.TrimSpace(strings.SplitN(h, ":", 2)[1])
			contentLen, _ = strconv.Atoi(v)
		}
	}

	body := []byte{}
	if contentLen > 0 {
		body = make([]byte, contentLen)
		_, err := io.ReadFull(reader, body)
		if err != nil {
			log.Println("read body error:", err)
			writeResponse(client, "400 Bad Request", "Bad Request")
			return
		}
	}

	for i, h := range headers {
		if strings.HasPrefix(strings.ToLower(h), "host:") {
			host := strings.Split(backend, ":")[0]
			headers[i] = "Host: " + host + "\r\n"
		}
	}
	headers = append(headers, "Connection: close\r\n")

	fullRequest := reqLine + "\r\n" + strings.Join(headers, "") + string(body)

	backendConn, err := net.Dial("tcp", backend)
	if err != nil {
		log.Println("backend dial error:", err)
		writeResponse(client, "502 Bad Gateway", "Backend down")
		return
	}
	defer backendConn.Close()

	_, err = backendConn.Write([]byte(fullRequest))
	if err != nil {
		log.Println("write to backend error:", err)
		return
	}

	_, err = io.Copy(client, backendConn)
	if err != nil {
		log.Println("copy back error:", err)
	}
}
