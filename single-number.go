// Given a non-empty array of integers, every element appears twice except for one. Find that single one.

// Input: [2,2,1]
// Output: 1

func singleNumber(nums []int) int {
    result := 0
    for _, val := range nums {
        result ^= val
    }
    return result
}
