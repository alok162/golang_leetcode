
// Given a sorted array nums, remove the duplicates in-place such that each element appear only once and return the new length.

// Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.

// Example 1:

// Given nums = [1,1,2],

// Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively.

// It doesn't matter what you leave beyond the returned length.


func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    i, j := 0, 0
    for j < len(nums)-1 {
        if nums[j] != nums[j+1] {
            nums[i] = nums[j]
            i ++
        }
        j ++
    }
    nums[i] = nums[len(nums)-1]
    return i+1
}
