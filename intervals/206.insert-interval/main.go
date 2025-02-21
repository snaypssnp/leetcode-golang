package main

// https://leetcode.com/problems/insert-interval/
func insert(intervals [][]int, newInterval []int) (res [][]int) {
	l := len(intervals) - 1

	if len(intervals) == 0 || intervals[0][0] > newInterval[1] {
		return append([][]int{newInterval}, intervals...)
	}

	if intervals[l][1] < newInterval[0] {
		return append(intervals, newInterval)
	}

	for i := 0; i <= l; i++ {
		interval := intervals[i]

		if i > 0 && newInterval[0] > intervals[i-1][1] && newInterval[1] < interval[0] {
			res = append(res, newInterval)
		}

		if newInterval[0] > interval[1] || newInterval[1] < interval[0] {
			res = append(res, interval)
			continue
		}

		start := min(interval[0], newInterval[0])

		if i == l {
			return append(res, []int{start, max(interval[1], newInterval[1])})
		}

		j := i

		for j < l && newInterval[1] >= intervals[j+1][0] {
			j++
		}
		i = j

		if i == 0 {
			start = min(start, newInterval[0])
		}

		res = append(res, []int{start, max(intervals[j][1], newInterval[1])})

	}

	return
}
