package main

import (
	"fmt"
	"time"
)

func producer(ch chan int, id int) {
for i := 1; i <= 5; i++ {
	 fmt.Printf("Producer %d producing %d\n", id, i)
	  ch <- i
		time.Sleep(time.Millisecond * 100)  // simulating work
}
}

func consumer(ch chan int, id int, done chan bool) {
for {
		val, ok := <-ch
	if !ok {
	 fmt.Printf("Consumer %d done, channel closed\n", id)
	done <- true
		return}
		fmt.Printf("Consumer %d consumed %d\n", id, val)
		time.Sleep(time.Millisecond * 150)  // simulating work
	}}

func ShowProducerConsumer() {
ch := make(chan int, 3)
done := make(chan bool)

	// starting 2 producers
go producer(ch, 1)
 go producer(ch, 2)

	// starting 3 consumers
	for i := 1; i <= 3; i++ {
		go consumer(ch, i, done)
	}

	// wait for producers to finish
	time.Sleep(time.Second * 2)
close(ch)  // closing channel to signal consumers no more data

	// wait for consumers to finish
	for i := 1; i <= 3; i++ {
		<-done}
	fmt.Println("All consumers finished")
}
