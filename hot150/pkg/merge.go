package pkg

import "sort"

func merge(intervals [][]int) (ans [][]int) {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0]-intervals[j][0] < 0
	})

	for i, interval := range intervals {
		if i == 0 {
			ans = append(ans, interval)
			continue
		}
		if last := len(ans) - 1; ans[last][1] >= interval[0] {
			ans[last][1] = max(interval[1], ans[last][1])
		} else {
			ans = append(ans, interval)
		}
	}
	return
}
