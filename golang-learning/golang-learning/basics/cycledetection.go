package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

func DetectCycle(head *Node) bool {
	slow := head
	 fast := head

	for fast != nil && fast.next != nil {
	 slow = slow.next
		fast = fast.next.next

		if slow == fast {
	return true
		}}
	return false
}

func ShowCycleDetection() {
	// create list: 1->2->3->4->5->3 (cycle)
 n1 := &Node{val:1}
n2 := &Node{val:2}
  n3 := &Node{val:3}
n4 := &Node{val:4}
n5 := &Node{val:5}

 n1.next = n2
n2.next = n3
n3.next = n4
 n4.next = n5
n5.next = n3 // cycle here

	fmt.Println("Cycle detected?", DetectCycle(n1))
}
