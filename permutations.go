// Given a collection of distinct integers, return all possible permutations.

// Input: [1,2,3]
// Output:
// [
//   [1,2,3],
//   [1,3,2],
//   [2,1,3],
//   [2,3,1],
//   [3,1,2],
//   [3,2,1]
// ]

func permute(nums []int) [][]int {
    var ans [][]int
    var index int
    var dfs func(index int) 
    
    dfs = func(index int) {
         if index == len(nums) {
            temp := make([]int, len(nums))
            copy(temp, nums)
            ans = append(ans, temp)
            return
        }
        for i := index; i < len(nums); i ++ {
            nums[index], nums[i] = nums[i], nums[index]
            dfs(index+1)
            nums[index], nums[i] = nums[i], nums[index]
        }
    }
    
    dfs(index)
    return ans
}
