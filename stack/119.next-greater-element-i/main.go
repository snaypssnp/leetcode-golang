package main

// https://leetcode.com/problems/next-greater-element-i/
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	stack := []int{}

	for _, v := range nums2 {
		for len(stack) > 0 && stack[len(stack)-1] < v {
			m[stack[len(stack)-1]] = v

			stack = stack[:len(stack)-1]
		}

		stack = append(stack, v)
	}

	for i, v1 := range nums1 {
		if v2, ok := m[v1]; ok {
			nums1[i] = v2
		} else {
			nums1[i] = -1
		}
	}

	return nums1
}
