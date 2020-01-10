package main

import "fmt"

func xorQueries(arr []int, queries [][]int) []int {
	out := make([]int, len(queries))
	for i, q := range queries {
		s := arr[q[0] : q[1]+1]
		res := 0
		for _, n := range s {
			res = res ^ n
		}
		out[i] = res
	}
	return out
}

func main() {
	arr := []int{4, 8, 2, 10}
	queries := [][]int{{2, 3}, {1, 3}, {0, 0}, {0, 3}}
	fmt.Println(xorQueries(arr, queries))
}
