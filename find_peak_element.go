
func findPeakElement(nums []int) int {
    low, high := 0, len(nums)-1
    var util func(nums []int, low int, high int) int
    
        util = func(nums []int, low int, high int) int {
            if low >= high {
                return low
            }
            if low + 1 == high {
                if nums[low] > nums[high] {
                    return low
                } else {
                    return high
                }
            }
            mid := (low + high)/2
            if mid>0 && mid<high && nums[mid - 1]<=nums[mid] && nums[mid+1]<=nums[mid]{
                return mid
            }
            if mid > 0 && nums[mid-1] > nums[mid] {
                return util(nums, low, mid-1)
            } else  {
                return util(nums, mid+1, high)
            }
        }
    
    return util(nums, low, high)
}
