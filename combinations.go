// Input: n = 4, k = 2
// Output:
// [
//   [2,4],
//   [3,4],
//   [2,3],
//   [1,2],
//   [1,3],
//   [1,4],
// ]

func combine(n int, k int) [][]int {
    var ans [][]int
    var path []int
    dfs(1, path, &ans, k, n)
    return ans
}

func dfs(index int, path []int, ans *[][]int, k int, n int) {
    if len(path) == k {
        // temp := make([]int, len(path))
        // copy(temp, path)
        *ans = append(*ans, append([]int{}, path...))
        return
    }
    for i := index; i <= n; i ++ {
        dfs(i+1, append(path, i), ans, k, n)
    }
}
