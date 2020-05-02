// Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

// Example:

// Input: [-2,1,-3,4,-1,2,1,-5,4],
// Output: 6
// Explanation: [4,-1,2,1] has the largest sum = 6.

func maxSubArray(nums []int) int {
    currSum := 0
    if len(nums) == 0 {
        return -2147483648
    }
    ans := nums[0]
    for i := 0; i < len(nums); i ++ {
        currSum += nums[i]
        if ans < currSum {
            ans = currSum
        }
        if currSum < 0 {
            currSum = 0
        }
    }
    return ans
}
