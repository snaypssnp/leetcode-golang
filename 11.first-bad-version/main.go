package main

import "fmt"

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
func firstBadVersion(n int) int {
	l := 0
	r := n
	for l < r {
		mid := (r + l) / 2
		if isBadVersion(mid) {
			if isBadVersion(mid-1) == false {
				return mid
			}

			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return l
}

func isBadVersion(n int) bool {
	return n >= 7
}
