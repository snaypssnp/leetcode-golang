package main

import "fmt"

// https://leetcode.com/problems/two-sum/description/
func main() {
	var nums = []int{2, 7, 11, 15}

	fmt.Println(twoSum(nums, 9))

}

func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		currentNum := nums[i]
		prevNum := target - currentNum

		if _, ok := numsMap[prevNum]; ok {
			return []int{numsMap[prevNum], i}
		}

		numsMap[currentNum] = i
	}
	return []int{-1, 1}
}
