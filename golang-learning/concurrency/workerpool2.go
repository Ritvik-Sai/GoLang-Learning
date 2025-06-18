package main

import (
  "fmt"
  "time"
)

// func worker2(id int, jobs <-chan int, results chan<- int) {
//   for j := range jobs {
//     fmt.Println("worker", id, "started job", j)
//     time.Sleep(time.Second)
//     fmt.Println("worker", id, "finished job", j)
//     results <- j * 2
//   }
// }

func ShowWorkerPool2() {
  jobs := make(chan int, 5)
  results := make(chan int, 5)

  for w := 1; w <= 3; w++ {
    go worker(w, jobs, results)
  }

  go func() {
    for res := range results {
      fmt.Println("collected result:", res)
    }}()

  for j := 1; j <= 5; j++ {
    jobs <- j }
close(jobs)

  time.Sleep(time.Second * 6) // wait longer for collector to finish
  close(results)
}
