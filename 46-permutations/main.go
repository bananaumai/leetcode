package main

import "fmt"

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}

	if len(nums) == 0 {
		return nil
	}

	var result [][]int
	for i, n := range nums {
		var rest []int
		rest = append(append(rest, nums[:i]...), nums[i+1:]...)
		for _, ns := range permute(rest) {
			result = append(result, append([]int{n}, ns...))
		}
	}

	return result
}

func main() {
	arr := []int{1,2,3}
	fmt.Println(permute(arr))
}
