func firstMissingPositive(nums []int) int {
    n := len(nums)
    for index, _ := range nums {
        if nums[index] <= 0 {
            nums[index] = n + 1
        }
    }
    for index, _ := range nums {
        val := abs(nums[index])
        if val <= n && nums[val-1] > 0 {
            nums[val-1] = -nums[val-1]
        }
    }
    for i:=1; i<=n; i++ {
        if nums[i-1] > 0 {
            return i
        }
    }
    return n+1
}

func abs(num int) int {
    if num > 0 {
        return num
    } else {
        return -num
    }
}
