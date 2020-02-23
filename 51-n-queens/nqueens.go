package main

import "fmt"

func main() {
	fmt.Println(solveNQueens(4))
}

func solveNQueens(n int) [][]string {
	var result [][]string

	for _, qs := range nQueens(n, n-1) {
		var ss []string
		for _, q := range qs {
			s := ""
			for i := 0; i < n; i++ {
				if q == i {
					s = fmt.Sprintf("%sQ", s)
				} else {
					s = fmt.Sprintf("%s.", s)
				}
			}
			ss = append(ss, s)
		}
		result = append(result, ss)
	}

	return result
}

func nQueens(n int, m int) [][]int {
	var qss [][]int

	if m == 0 {
		for q := 0; q < n; q++ {
			qss = append(qss, []int{q})
		}
		return qss
	}

	for _, qs := range nQueens(n, m-1) {
		for newQ := 0; newQ < n; newQ++ {
			addable := true
			for k, q := range qs {
				diff := m-k
				switch newQ {
				case q, q+diff, q-diff:
					addable = false
				}
			}
			if addable {
				newQs := append([]int{}, qs...)
				newQs = append(newQs, newQ)
				qss = append(qss, newQs)
			}
		}
	}

	return qss
}
