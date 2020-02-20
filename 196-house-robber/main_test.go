package main

import "testing"

func TestRob(t *testing.T) {
	type test struct {
		input []int
		expect int
	}

	 tests := []test{
		{[]int{2,1,1,2}, 4},
		{[]int{1,2,3,1}, 4},
	}

	for _, tst := range tests {
		actual := rob(tst.input)
		if tst.expect != actual {
			t.Errorf("input: %v, expect: %d, actual: %d", tst.input, tst.expect, actual)
		}
	}
}
