// Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.

// Example:

// Input: [0,1,0,3,12]
// Output: [1,3,12,0,0]
// Note:

// You must do this in-place without making a copy of the array.
// Minimize the total number of operations.


func moveZeroes(nums []int)  {
    zero := 0
    for index, _ := range nums {
        if nums[zero] == 0 && nums[index] != 0 {
            nums[zero], nums[index] = nums[index], nums[zero]
            zero ++
        } else if nums[zero] != 0 {
            zero ++
        }
    }
}
