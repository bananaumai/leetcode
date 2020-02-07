package main

var memo = make(map[string]int)

func minInsertions2(s string) int {
	if len(s) <= 1 {
		return 0
	}

	if cnt, ok := memo[s]; ok {
		return cnt
	}

	var substrs []string
	for i := 0; i < len(s); i++ {
		c := s[i]
		for j := len(s)-1; j > i; j-- {
			if c == s[j] {
				substr := s[i+1:j]
				substrs = append(substrs, substr)
			}
		}
	}

	strLen := len(s)
	min := strLen - 1

	if len(substrs) == 0 {
		memo[s] = min
		return min
	}

	for _, substr := range substrs {
		charsToAdd := strLen - len(substr) - 2

		count := minInsertions2(substr)

		count += charsToAdd

		if count == 0 {
			min = charsToAdd
			break
		}

		if count < min {
			min = count
		}
	}

	memo[s] = min
	return min
}
