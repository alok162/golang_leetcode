// Given a set of distinct integers, nums, return all possible subsets (the power set).

// Note: The solution set must not contain duplicate subsets.

// Example:

// Input: nums = [1,2,3]
// Output:
// [
//   [3],
//   [1],
//   [2],
//   [1,2,3],
//   [1,3],
//   [2,3],
//   [1,2],
//   []
// ]


func subsets(nums []int) [][]int {
    sort.Ints(nums)
    var ans [][]int
    var subset []int
    subsetsHelper(0, subset, &ans, nums)
    return ans
}

func subsetsHelper(index int, subset []int, ans *[][]int, nums []int) {
    *ans = append(*ans, append([]int{}, subset...))
    for i := index; i < len(nums); i ++ {
        subsetsHelper(i+1, append(subset, nums[i]), ans, nums)
    }
}
