package main

// https://leetcode.com/problems/can-place-flowers/
func canPlaceFlowers(flowerbed []int, n int) bool {
	var counter int

	l := len(flowerbed) - 1

	for i := 0; i <= l; i++ {
		if flowerbed[i] == 0 {
			if (i == 0 || flowerbed[i-1] == 0) && (i == l || flowerbed[i+1] == 0) {
				flowerbed[i] = 1
				counter++
			}
		}

		if counter >= n {
			return true
		}
	}

	return false
}
