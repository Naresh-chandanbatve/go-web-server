package main

import (
	"fmt"
	"net"
)

func main() {
	cfg := LoadConfig("config.conf")

	ln, err := net.Listen("tcp", ":"+cfg.ListenPort)
	if err != nil {
		panic(err)
	}

	fmt.Println("Mini NGINX listening on", cfg.ListenPort)

	for {
		conn, _ := ln.Accept()
		go handleConnection(conn, cfg)
	}
}
