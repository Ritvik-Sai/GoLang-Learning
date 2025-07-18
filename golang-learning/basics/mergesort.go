package main

import "fmt"

func merge(left, right []int) []int {
	result := []int{}
i,j := 0, 0

	for i < len(left) && j < len(right) {
	if left[i] < right[j] {
		 result = append(result, left[i])
			i++
	}else {
		   result = append(result, right[j])
			j++
		}
	}

 result = append(result, left[i:]...)
result = append(result, right[j:]...)
  	return result
}

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr}

	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
right := MergeSort(arr[mid:])
	return merge(left, right)
}

func ShowMergeSort() {
  arr := []int{8, 3, 7, 4, 9, 1}
	fmt.Println("Before sort:", arr)
	 sorted := MergeSort(arr)
	fmt.Println("After sort:", sorted)
}
