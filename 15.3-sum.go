// Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

// Note:

// The solution set must not contain duplicate triplets.



// Given array nums = [-1, 0, 1, 2, -1, -4],

// A solution set is:
// [
//   [-1, 0, 1],
//   [-1, -1, 2]
// ]

func threeSum(nums []int) [][]int {
    ans := [][]int{}
    sort.Ints(nums)
    for i := 0; i < len(nums)-2; i ++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        j, k := i+1, len(nums)-1
        for j < k {
            sum := nums[i] + nums[j] + nums[k]
            switch {
                case sum > 0:
                    k --
                case sum < 0:
                    j ++
                default:
                    ans = append(ans, []int{nums[i], nums[j],nums[k]})
                    j, k = removeDuplicate(nums, j, k)
            }
        }
    }
    return ans
}

func removeDuplicate(nums []int, left, right int) (int, int) {
    for left < right {
        switch {
        case nums[left] == nums[left+1]:
            left ++
        case nums[right] == nums[right-1]:
            right --
        default:
            left ++
            right --
            return left, right
        }
    }
   
    return left, right
}
