package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	done := make(chan struct{})
	go func() { //fork point
		work()
		done <- struct{}{}
	}()
	<-done //chennel join point
	fmt.Println("elapsed: ", time.Since(now))
	fmt.Println("done waiting, main exits")

}

func work() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("printing some stuff")
}
