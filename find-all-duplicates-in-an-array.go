// Input:
// [4,3,2,7,8,2,3,1]

// Output:
// [2,3]

func findDuplicates(nums []int) []int {
    var result []int
    for index := range nums {
        curr := abs(nums[index])
        if nums[curr-1] < 0 {
            result = append(result, curr)
        } else {
            nums[curr-1] = -nums[curr-1]
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
