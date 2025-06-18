package main

import (
	"fmt"
	"strconv"
)

// generate numbers
func generate(nums ...int) <-chan int {
	out := make(chan int)
go func() {
	for _, n := range nums {
			out <- n}
		close(out)}()
	return out
}

// square numbers
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
	for n := range in {
	out <- n * n
		}
		close(out)
	}()
return out
}

// convert numbers to string
func convertToString(in <-chan int) <-chan string {
out := make(chan string)
	go func() {
		for n := range in {
		out <- "Num: " + strconv.Itoa(n)
		}
		close(out)
	}()
return out
}

func ShowPipeline() {
	nums := []int{2, 3, 4, 5}

 numbers := generate(nums...)
squared := square(numbers)
strs := convertToString(squared)

	for s := range strs {
		fmt.Println(s)
	}
}
