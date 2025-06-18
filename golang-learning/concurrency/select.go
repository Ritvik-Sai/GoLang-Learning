package main

import (
	"fmt"
	"time"
)

func ShowSelect() {
	ch := make(chan string)

go func() {
		time.Sleep(2 * time.Second)
		ch <- "result from channel"
	}()

select {
case res := <-ch:
	 fmt.Println("Received:", res)
case <-time.After(1 * time.Second):
	 fmt.Println("Timeout: no response received in 1 second")
	}}
