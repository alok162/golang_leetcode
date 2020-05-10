// Given an array, rotate the array to the right by k steps, where k is non-negative.

// Follow up:

// Try to come up as many solutions as you can, there are at least 3 different ways to solve this problem.
// Could you do it in-place with O(1) extra space?

// Input: nums = [1,2,3,4,5,6,7], k = 3
// Output: [5,6,7,1,2,3,4]

func rotate(nums []int, k int)  {
    k = k % len(nums)
    reverse(&nums, 0, len(nums)-1)
    reverse(&nums, 0, k-1)
    reverse(&nums, k, len(nums)-1)
}

func reverse(nums *[]int, start int, end int) {
    i, j := start, end
    for i < j {
        (*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
        i ++
        j --
    }
}
