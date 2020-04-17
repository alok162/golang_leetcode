func longestValidParentheses(s string) int {
    left, right := 0, 0
    maxResult := 0
    for _, val := range s {
        if string(val) == "(" {
            left += 1
        } else if string(val) == ")" {
            right += 1
        }
        if left == right && maxResult < 2*left {
            maxResult = 2*left
        }
        if right > left {
            left, right = 0, 0
        }
    }
    left, right = 0, 0
    i := len(s)-1
    for i >= 0 {
        if string(s[i]) == "(" {
            left += 1
        } else if string(s[i]) == ")" {
            right += 1
        }
        if left == right && maxResult < 2* left {
            maxResult = 2 * left
        }
        if right < left {
            left, right = 0, 0
        }
        i -= 1
    }
    return maxResult
}
