package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool, 1)
	go func() {
		time.Sleep(1 * time.Second)
		<-c
	}()
	start := time.Now()
	c <- true // blocks until other goroutine receives
	fmt.Printf("send took %v\n", time.Since(start))

}
