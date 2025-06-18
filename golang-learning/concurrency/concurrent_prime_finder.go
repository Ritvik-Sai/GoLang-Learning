package main

import (
	"fmt"
	"sync"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false}
for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false}}
	return true
}

func checkPrimeWorker(nums []int, wg *sync.WaitGroup, primesChan chan int) {
	defer wg.Done()

	for _, num := range nums {
		 if isPrime(num) {
		primesChan <- num
		}}
}

func ShowConcurrentPrimeFinder() {
	nums := []int{}
for i := 1; i <= 50; i++ {
	    	nums = append(nums, i)
	}

 workerCount := 5
chunkSize := len(nums) / workerCount

var wg sync.WaitGroup
  primesChan := make(chan int, 10)

	for i := 0; i < workerCount; i++ {
	start := i * chunkSize
	end := start + chunkSize

		if i == workerCount-1 {
		end = len(nums)}
	wg.Add(1)
	go checkPrimeWorker(nums[start:end], &wg, primesChan)
	}

	go func() {
		wg.Wait()
		close(primesChan)}()

	fmt.Println("Prime numbers found:")
	for prime := range primesChan {
		fmt.Println(prime)}
}
