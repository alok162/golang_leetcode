func sortColors(nums []int)  {
    zero := 0
    one := 0
    high := len(nums)-1
    for one <=  high{
        if nums[one] == 0 {
            nums[zero], nums[one] = nums[one], nums[zero]
            zero += 1
            one += 1
        } else if nums[one] == 1{
            one += 1
        } else {
            nums[one], nums[high] = nums[high], nums[one]
            high -= 1
        }
    }
}
