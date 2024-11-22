package main

import (
	"encoding/json"
	"fmt"
	geerpc "gee-rpc"
	"gee-rpc/codec"
	"log"
	"net"
)

func main() {
	conn, _ := net.Dial("tcp", ":6666")
	defer func() {
		_ = conn.Close()
	}()

	_ = json.NewEncoder(conn).Encode(geerpc.DefaultOption)
	cc := codec.NewGobCodec(conn)

	for i := 0; i < 5; i++ {
		h := &codec.Header{
			ServiceMethod: "Service.Func",
			Seq:           uint64(i),
		}
		_ = cc.Write(h, fmt.Sprintf("geerpc request seq: %d", h.Seq))
		_ = cc.ReadHeader(h)
		var reply string
		_ = cc.ReadBody(&reply)
		log.Println("reply:", reply)
	}
}
