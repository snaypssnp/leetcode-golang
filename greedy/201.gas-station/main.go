package main

// https://leetcode.com/problems/gas-station/
func canCompleteCircuit(gas []int, cost []int) int {
	if sum(gas) < sum(cost) {
		return -1
	}

	tank := 0
	l := len(gas)
	start := 0

	for i := 0; i < l; i++ {
		tank += gas[i] - cost[i]

		if tank < 0 {
			tank = 0
			start = i + 1
		}
	}

	return start
}

func sum(arr []int) (res int) {
	for _, n := range arr {
		res += n
	}

	return
}
