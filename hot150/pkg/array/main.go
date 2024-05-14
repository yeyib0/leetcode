
func removeDuplicates(nums []int) int {
    i, r := 0, 0
    // 指针i表示当前要处理的数字
    for i < len(nums) {
        j := 0
        // 删除连续的重复项，i+1+j从i的下一个开始
        for i+j+1 < len(nums) {
            if (nums[i+j] == nums[i+j+1]) {
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
