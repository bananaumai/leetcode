package main

import "fmt"

func main() {
	fmt.Println(solveNQueens(4))
}

func solveNQueens(n int) [][]string {
	var result [][]string

	qss := nQueensRec(n, n-1)
	//qss := nQueensRecCPS(n, 0, [][]int{})
	//qss := nQueensLoop(n)

	for _, qs := range qss {
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

func nQueensRec(n int, m int) [][]int {
	var qss [][]int

	if m == 0 {
		for q := 0; q < n; q++ {
			qss = append(qss, []int{q})
		}
		return qss
	}

	for _, qs := range nQueensRec(n, m-1) {
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

func nQueensRecCPS(n int, m int, acc [][]int) [][]int {
	if m == n {
		return acc
	}

	var qss [][]int
	var next = m+1

	if m == 0 {
		for q := 0; q < n; q++ {
			qss = append(qss, []int{q})
		}
		return nQueensRecCPS(n, next, qss)
	}

	for _, qs := range acc {
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

	return nQueensRecCPS(n, next, qss)
}

func nQueensLoop(n int) [][]int {
	var qss [][]int

	for m := 0; m < n; m++ {
		if m == 0 {
			for q := 0; q < n; q++ {
				qss = append(qss, []int{q})
			}
			continue
		}

		var newQss [][]int

		for _, qs := range qss {
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
					newQss = append(newQss, newQs)
				}
			}
		}

		qss = newQss
	}

	return qss
}
