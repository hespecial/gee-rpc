package main

import (
	geerpc "gee-rpc"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":6666")
	if err != nil {
		log.Fatal("network error:", err)
	}
	log.Println("rpc server start on:", listener.Addr())
	geerpc.Accept(listener)
}
