package main

import "fmt"

func main() {
	fmt.Println(solveNQueens(4))
}

func solveNQueens(n int) [][]string {
	var result [][]string

	qss := placeQueen(n, 0, [][]int{})

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

func placeQueen(size int, rowPos int, acc [][]int) [][]int {
	if rowPos >= size {
		return acc
	}

	var nextRowPos = rowPos+1
	var qss [][]int

	if rowPos == 0 {
		for col := 0; col < size; col++ {
			qss = append(qss, []int{col})
		}
		return placeQueen(size, nextRowPos, qss)
	}

	for _, qPlacedCols := range acc {
		for col := 0; col < size; col++ {
			if canPlaceQueen(qPlacedCols, rowPos, col) {
				newQs := append([]int{}, qPlacedCols...)
				newQs = append(newQs, col)
				qss = append(qss, newQs)
			}
		}
	}

	return placeQueen(size, nextRowPos, qss)
}

func canPlaceQueen(cols []int, row int, col int) bool {
	for r, c := range cols {
		diff := row-r
		switch col {
		case c, c+diff, c-diff:
			return false
		}
	}
	return true
}
