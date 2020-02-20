package main

func rob(nums []int) int {
	length := len(nums)

	if length == 0 {
		return 0
	}

	if length == 1 {
		return nums[0]
	}

	res := make([]int, length)
	for i, n := range nums {
		switch i {
		case 0, 1:
			res[i] = n
		case 2:
			res[i] = n + res[i-2]
		default:
			res[i] = max(n + res[i-2], n + res[i-3])
		}
	}

	return max(res[len(res)-1], res[len(res)-2])
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
