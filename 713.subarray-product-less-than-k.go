// You are given an array of positive integers nums.

// Count and print the number of (contiguous) subarrays where the product of all the elements in the subarray is less than k.

// Example 1:
// Input: nums = [10, 5, 2, 6], k = 100
// Output: 8
// Explanation: The 8 subarrays that have product less than 100 are: [10], [5], [2], [6], [10, 5], [5, 2], [2, 6], [5, 2, 6].
// Note that [10, 5, 2] is not included as the product of 100 is not strictly less than k.

func numSubarrayProductLessThanK(nums []int, k int) int {
    left, right := 0, 0
    result := 0
    prod := 1
    for index, _ := range nums {
        prod *= nums[index]
        right = index
        for prod >= k && left < right {
            prod /= nums[left]
            left ++
        }
        if prod < k {
            result += right - left + 1
        }
        fmt.Println(right - left + 1)
    }
    return result
}
