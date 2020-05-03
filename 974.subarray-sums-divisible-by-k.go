// Given an array A of integers, return the number of (contiguous, non-empty) subarrays that have a sum divisible by K.

// Example 1:

// Input: A = [4,5,0,-2,-3,1], K = 5
// Output: 7
// Explanation: There are 7 subarrays with a sum divisible by K = 5:
// [4, 5, 0, -2, -3, 1], [5], [5, 0], [5, 0, -2, -3], [0], [0, -2, -3], [-2, -3]

func subarraysDivByK(A []int, K int) int {
    currSum := 0
    dict := make(map[int]int)
    dict[0] = 1
    result := 0
    for _, val := range A {
        currSum += val
        remainder := currSum % K
        if remainder < 0 {
            remainder += K
        }
        if _, ok := dict[remainder]; ok {
            result += dict[remainder]
        }
        dict[remainder] += 1
    }
    fmt.Println(dict)
    return result
}
