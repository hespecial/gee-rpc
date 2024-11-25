package main

import (
	"fmt"
	gee_rpc "gee-rpc"
	"log"
	"sync"
)

func main() {
	log.SetFlags(0)
	client, _ := gee_rpc.Dial("tcp", ":6666")
	defer func() {
		_ = client.Close()
	}()

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			args := fmt.Sprintf("geerpc req %d", i)
			var reply string
			if err := client.Call("Foo.Sum", args, &reply); err != nil {
				log.Fatal("call Foo.Sum error:", err)
			}
			log.Println("reply:", reply)
		}(i)
	}
	wg.Wait()
}
