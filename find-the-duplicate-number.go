// Given an array nums containing n + 1 integers where each integer is between 1 and n (inclusive), prove that at least one duplicate number must exist. Assume that there is only one duplicate number, find the duplicate one.

// Note:

// You must not modify the array (assume the array is read only).
// You must use only constant, O(1) extra space.
// Your runtime complexity should be less than O(n2).
// There is only one duplicate number in the array, but it could be repeated more than once.

// Input: [1,3,4,2,2]
// Output: 2

func findDuplicate(nums []int) int {
    sort.Ints(nums)
    for i:=1; i < len(nums); i++ {
        if nums[i-1] == nums[i] {
            return nums[i]
        }
    }
    return 0
}
