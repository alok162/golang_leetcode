// Given a set of candidate numbers (candidates) (without duplicates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target.

// The same repeated number may be chosen from candidates unlimited number of times.

// Note:

// All numbers (including target) will be positive integers.
// The solution set must not contain duplicate combinations.

// Example 1:

// Input: candidates = [2,3,6,7], target = 7,
// A solution set is:
// [
//   [7],
//   [2,2,3]
// ]

func combinationSum(candidates []int, target int) [][]int {
    var ans [][]int
    var combination []int
    sort.Ints(candidates)
    dfs(0, target, combination, candidates, &ans)
    return ans
}

func dfs(index int, target int, combination []int, candidates []int, ans *[][]int) {
    if target == 0 {
        *ans = append(*ans, append([]int{}, combination...))
        return
    }
    for i := index; i < len(candidates); i ++ {
        if candidates[i] > target {
            break
        }
        dfs(i, target-candidates[i], append(combination, candidates[i]), candidates, ans)
    }
}
