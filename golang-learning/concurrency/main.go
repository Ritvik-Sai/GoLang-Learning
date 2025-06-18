package main

import "fmt"

func main() {

   fmt.Println("Start concurrency test ==")
   ShowGoroutines()

   	fmt.Println("\n runn worker pool test")
	ShowWorkerPool()

	 fmt.Println("\n run worker pool 2 test")
    ShowWorkerPool2()

	fmt.Println("\n runchannels test ")
ShowChannels()

	fmt.Println("\nrun channels 2 test")
	 ShowChannels2()

	 fmt.Println("\n run mutex test ")
	ShowMutex()

	fmt.Println("\n run select test")
ShowSelect()

}
