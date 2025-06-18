package main

import "fmt"

func BinarySearch(arr []int, target int) int {
	low := 0
high:= len(arr) -1

	for low <= high {
	mid := (low + high) / 2

	 if arr[mid] == target {
			return mid
	} else if arr[mid] < target {
			low = mid +1
	 	}else {
			high = mid -1
		}}
return -1
}

func ShowBinarySearch() {
	arr := []int{1, 3, 5, 7, 9, 11}
	target := 7

pos := BinarySearch(arr, target)
   if pos != -1 {
		fmt.Println("found target at index:", pos)
}else {
	fmt.Println("target not found")
	}
}
