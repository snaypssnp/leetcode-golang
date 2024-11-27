package main

import (
	"fmt"
)

// https://leetcode.com/problems/first-bad-version/description/
func main() {
	fmt.Println(firstBadVersion(10))
}

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

// Time: O(log n), Space: O(1)
func firstBadVersion(n int) (ans int) {
	ans = -1

	for left, right := 0, n; left <= right; {
		mid := left + (right-left)/2

		if isBadVersion(mid) {
			ans = mid

			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return
}

func isBadVersion(n int) bool {
	return n >= 7
}
