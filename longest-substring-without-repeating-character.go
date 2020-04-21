func lengthOfLongestSubstring(s string) int {
    maxLength := 0
    start := 0
    dict := make(map[string]int)
    for index := range s {
        
        if _, ok := dict[string(s[index])]; ok  {
            start = max(start, dict[string(s[index])] + 1)
            // dict[string(s[index])] = index
        }
        dict[string(s[index])] = index
        maxLength = max(maxLength, index - start + 1)
    }
    return maxLength
}

func max(x1 int, x2 int) int {
    if x1 > x2 {
        return x1
    } else {
        return x2
    }
}
