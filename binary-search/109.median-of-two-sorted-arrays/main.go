package main

// https://leetcode.com/problems/median-of-two-sorted-arrays/
func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	l := l1 + l2
	nums := make([]int, l)

	for i, j, k := 0, 0, 0; k < l; k++ {
		if i >= l1 {
			nums[k] = nums2[j]
			j++
		} else if j >= l2 {
			nums[k] = nums1[i]
			i++
		} else if nums1[i] < nums2[j] {
			nums[k] = nums1[i]
			i++
		} else {
			nums[k] = nums2[j]
			j++
		}
	}

	mid1, mid2 := (l-1)/2, l/2

	return float64(nums[mid1]+nums[mid2]) / 2.0
}
