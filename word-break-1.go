// Input: s = "leetcode", wordDict = ["leet", "code"]

func wordBreak(s string, wordDict []string) bool {
    dp := make([]bool, len(s)+1)
    dp[0] = true
    dict := make(map[string]bool)
    for _, val := range wordDict {
        dict[val] = true
    }
    for i := 1; i < len(s) + 1; i++ {
        for j := 0; j < i; j++ {
            if dp[j] == true && dict[s[j:i]] {
                dp[i] = true
            }
        }
    }
    return dp[len(dp)-1]
}
