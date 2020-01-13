package main

import "fmt"

func xorQueries(arr []int, queries [][]int) []int {
	calc := make([]int, len(arr))
	for i, e := range arr {
		if i == 0 {
			calc[i] = e
		} else {
			calc[i] = calc[i-1] ^ e
		}
	}

	out := make([]int, len(queries))
	for i, q := range queries {
		if q[0] > 0 {
			out[i] = calc[q[0]-1] ^ calc[q[1]]
		} else {
			out[i] = calc[q[1]]
		}
	}
	return out
}

func main() {
	arr := []int{15, 8, 8, 8, 15}
	queries := [][]int{{2, 2}, {3, 3}}
	fmt.Println(xorQueries(arr, queries))
}
