// Given an array of integers where 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.

// Find all the elements of [1, n] inclusive that do not appear in this array.

// Could you do it without extra space and in O(n) runtime? You may assume the returned list does not count as extra space.

// Example:

// Input:
// [4,3,2,7,8,2,3,1]

// Output:
// [5,6]



func findDisappearedNumbers(nums []int) []int {
    var result []int
    for index := range nums {
        curr := abs(nums[index])
        if nums[curr-1] > 0 {
            nums[curr-1] = -nums[curr-1]
        }
    }
    
    for index := range nums {
        if nums[index] > 0 {
            result = append(result, index+1)
        }
    }
    
    return result
}

func abs(n int) int {
    if n > 0 {
        return n
    } else {
        return -n
    }
}
