package main

import "sort"

// https://leetcode.com/problems/hand-of-straights/
func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}

	m := map[int]int{}

	for _, num := range hand {
		m[num]++
	}

	l, i := len(m), 0

	queue := make([][2]int, l)

	for value, frequency := range m {
		queue[i] = [2]int{value, frequency}
		i++
	}

	sort.Slice(queue, func(i, j int) bool {
		return queue[i][0] > queue[j][0]
	})

	for len(queue) > 0 {
		arr := [][2]int{}

		for i := 0; i < groupSize && len(queue) > 0; i++ {
			l := len(queue) - 1
			numFrequency := queue[l]

			if len(arr) > 0 && numFrequency[0]-arr[len(arr)-1][0] > 1 {
				return false
			}

			queue = queue[:l]

			numFrequency[1]--
			arr = append(arr, numFrequency)
		}

		if len(arr) != groupSize {
			return false
		}

		for i := len(arr) - 1; i >= 0; i-- {
			numFrequency := arr[i]

			if numFrequency[1] > 0 {
				queue = append(queue, numFrequency)
			}
		}
	}

	return true
}
