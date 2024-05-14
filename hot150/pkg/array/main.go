
package array

import "sort"

func removeDuplicates(nums []int) int {
	i, r := 0, 0
	// 指针i表示当前要处理的数字
	for i < len(nums) {
		j := 0
		// 删除连续的重复项，i+1+j从i的下一个开始
		for i+j+1 < len(nums) {
			if nums[i+j] == nums[i+j+1] {
				j++
			} else {
				break
			}
		}
		nums[r] = nums[i]
		r++
		// 连续出现多次的，保留两次，此处多保留一次
		if j > 0 {
			nums[r] = nums[i]
			r++
		}
		i = i + j + 1
	}
	return r
}

// IntToRoman 整数转罗马数字
// https://leetcode.cn/problems/integer-to-roman/description/
func IntToRoman(num int) (ans string) {
	// hash存储1,4,5,9...1000
	hash := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}
	set := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	// for这个hash，从最大开始减num
	for _, k := range set {
		v := hash[k]
		// for到num小于val
		for {
			if num >= k {
				num = num - k
				ans = ans + v
			} else {
				break
			}
		}
	}
	return
}

// Convert Z字形变换
// https://leetcode.cn/problems/zigzag-conversion/?envType=study-plan-v2&envId=top-interview-150
func Convert(s string, numRows int) (ans string) {
	// 如果只有一个数组，直接返回
	if numRows == 1 {
		ans = s
		return
	}
	// string数组，每一行一个
	strs := make([]string, numRows)
	// 高度计数hi，先从0开始+到numRows,然后减到0（减的过程数组的）
	pos := 0
	// 表示高度递增
	startToEnd := true
	// for 字符串
	for i := range s {
		strs[pos] = strs[pos] + string(s[i])

		if startToEnd {
			pos++
		} else {
			pos--
		}
		// 如果递增到最后一位，或者递减到第一位，改变方向
		if (startToEnd && pos == numRows-1) || (!startToEnd && pos == 0) {
			startToEnd = !startToEnd
		}
	}
	// 拼字符串
	for _, str := range strs {
		ans = ans + str
	}
	return
}

// MaxArea https://leetcode.cn/problems/container-with-most-water/description/?envType=study-plan-v2&envId=top-interview-150
// 盛水醉的容器
// 证明移动大的不行，因为两端取的最小高度
func MaxArea(height []int) (ans int) {
	// 左右指针
	l, r := 0, len(height)-1
	for r != l {
		// 计算容量
		ans = max(ans, (r-l)*min(height[r], height[l]))
		// 移动小的
		if height[r] > height[l] {
			l++
		} else {
			r--
		}
	}
	return
}

// https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/submissions/531827518/?envType=study-plan-v2&envId=top-interview-150
func twoSum(nums []int, target int) []int {
	// 前后，计算和，大移动后面，小移动前面
	l, r := 0, len(nums)-1
	for l != r {
		if nums[l]+nums[r] > target {
			r--
		} else if nums[l]+nums[r] < target {
			l++
		} else {
			return []int{l + 1, r + 1}
		}
	}
	return []int{-1, -1}
}

// ThreeSum https://leetcode.cn/problems/3sum/description/?envType=study-plan-v2&envId=top-interview-150
// 三数之和
func ThreeSum(nums []int) (ans [][]int) {
	// 排序
	sort.Ints(nums)
	// 先确定一个数
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			return
		}
		// 第一位数，如果和前面的数相同，结果已经包含在ans中
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		sum := 0 - nums[i]
		// 双指针
		l, r := i+1, len(nums)-1
		for l < r {
			if nums[l]+nums[r] > sum {
				r--
			} else if nums[l]+nums[r] < sum {
				l++
			} else {
				ans = append(ans, []int{nums[i], nums[l], nums[r]})
				// 减少重复的
				// 左侧当前数字是l-1（已经+1了）
				for l++; l < r && nums[l] == nums[l-1]; l++ {
				}
				// 右侧当前数字是r+1（已经-1了）
				for r--; l < r && nums[r] == nums[r+1]; r-- {
				}
			}
		}
	}
	return
}
