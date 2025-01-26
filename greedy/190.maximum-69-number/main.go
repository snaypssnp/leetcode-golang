package main

import "strconv"

func maximum69Number(num int) int {
	nums := []byte(strconv.Itoa(num))

	for i := 0; i < len(nums); i++ {

		if nums[i] == '6' {
			nums[i] = '9'
			break
		}
	}

	res, _ := strconv.Atoi(string(nums))

	return res
}
