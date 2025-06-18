package main

import(
 "fmt"
 "time"
)

func Ping(pings chan<- string, msg string){

pings <- msg}

func Pong(pings <-chan string, pongs chan<- string){
	msg := <- pings
	 pongs <- msg
}

func ShowChannels(){

pings := make(chan string , 1)
	pongs := make(chan string , 1)
	 Ping(pings , "hello")
   Pong(pings , pongs)

	fmt.Println( "message is", <-pongs)
	time.Sleep(time.Millisecond * 500)  // waits
}
