package main

import (
	"fmt"
	"time"
)

func ShowChannels2() {

fmt.Println("Demo: Unbuffered channel")
 ch := make(chan string)

	go func() {
		fmt.Println("sending hello")
	ch <- "hello"
	  fmt.Println("sent hello")}()

  time.Sleep(time.Second)

	fmt.Println("receiving")
	msg := <-ch
fmt.Println("received:", msg)
	fmt.Println("\nDemo: Buffered channel")

	ch2 := make(chan string, 2)
	ch2 <- "buffer1"
	ch2 <- "buffer2"

fmt.Println("sent two messages into buffered channel")
	fmt.Println("receiving", <-ch2)
fmt.Println("receiving", <-ch2)

}
