func maxProduct(nums []int) int {
    left_max := nums[0]
    right_max := nums[len(nums)-1]
    left := nums[0]
    for i:=1; i<len(nums); i++ {
        left *= nums[i]
        if left == 0 && nums[i] != 0 {
            left = nums[i]
        }
        if left > left_max {
            left_max = left
        }
    }
    right := nums[len(nums)-1]
    for i:=len(nums)-2; i>-1; i-- {
        right *= nums[i]
        if right == 0 && nums[i] != 0 {
            right = nums[i]
        }
        if right > right_max {
            right_max = right
        }
    }
    if right_max > left_max {
        return right_max
    } else {
        return left_max
    }
}
