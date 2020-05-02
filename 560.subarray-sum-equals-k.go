// Given an array of integers and an integer k, you need to find the total number of continuous subarrays whose sum equals to k.

// Example 1:
// Input:nums = [1,1,1], k = 2
// Output: 2

func subarraySum(nums []int, k int) int {
    dict := make(map[int]int)
    sum := 0
    result := 0
    for _, val := range nums {
        sum += val
        if sum == k {
            result += 1
        }
        if _, ok := dict[sum-k]; ok {
            result += dict[sum-k]
        }
        dict[sum] += 1
    }
    return result
}
