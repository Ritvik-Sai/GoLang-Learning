package main

import (
    "fmt"
    "time"
)

func say(s string){
    for i:=0; i<3; i++ {
      fmt.Println(s , ":", i)
     time.Sleep(100 * time.Millisecond)
    }
}

func ShowGoroutines() {
    fmt.Println("starting goroutines...")
   go say("goroutine 1")
    go say("goroutine 2")

   // waiting a bit for goroutines to finish
   	time.Sleep(500 * time.Millisecond)
    fmt.Println("done with goroutines")}
