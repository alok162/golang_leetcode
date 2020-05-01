// Find all possible combinations of k numbers that add up to a number n, given that only numbers from 1 to 9 can be used and each combination should be a unique set of numbers.

// Note:

// All numbers will be positive integers.
// The solution set must not contain duplicate combinations.

// Input: k = 3, n = 9
// Output: [[1,2,6], [1,3,5], [2,3,4]]


func combinationSum3(k int, n int) [][]int {
    var combination []int
    var ans [][]int
    dfs(1, n, k, combination, &ans)
    return ans
}

func dfs(index int, n int, k int, combination []int, ans *[][]int) {
    if n == 0 && len(combination) == k {
        *ans = append(*ans, append([]int{}, combination...))
        return
    }
    for i := index; i < 10; i ++ {
        if len(combination) > k {
            break
        }
        dfs(i+1, n-i, k, append(combination, i), ans)
    }
}
