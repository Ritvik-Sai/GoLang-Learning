package main

import "fmt"

func DFS(graph map[int][]int, start int, visited map[int]bool) {
	if visited[start] {
		return}

	fmt.Print(start, " ")
visited[start] = true

for _, v := range graph[start] {
		DFS(graph, v, visited)}}

func ShowDFS() {
	graph := map[int][]int{
	1: {2, 3},
	2: {4},
	3: {4, 5},
	4: {},
	5: {1},   } 

	visited := make(map[int]bool)
  fmt.Print("DFS Traversal: ")
	DFS(graph, 1, visited)
  fmt.Println()
}
