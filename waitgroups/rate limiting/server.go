package main

import (
	"log"
	"net"
	"sync/atomic"
	"time"
)

func main() {

	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("could not create listener: %v", err)
	}
	var connections int32
	for {
		conn, err := li.Accept() //create listener at port 8080
		if err != nil {
			continue
		}
		connections++ //increment count of connections
		//routine to serve connection
		go func() {
			defer func() { //function to close connection and decrement count
				_ = conn.Close()
				atomic.AddInt32(&connections, -1)
			}()
			//serve maximum 3 request in parallel
			if atomic.LoadInt32(&connections) > 3 {
				return
			}

			time.Sleep(time.Second)
			_, err := conn.Write([]byte("success")) //write byte slice to connection
			if err != nil {
				log.Fatalf("could not write to connection: %v", err)
			}
		}()
	}
}
