package main

// https://leetcode.com/problems/daily-temperatures/
func dailyTemperatures(temperatures []int) []int {
	var l = len(temperatures)
	var stack = []int{}
	var ans = make([]int, l)

	for i := 0; i < l; i++ {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
			n := len(stack) - 1
			topIndex := stack[n]
			ans[topIndex] = i - topIndex

			stack = stack[:n]
		}

		stack = append(stack, i)
	}

	return ans
}
