import "fmt"
func titleToNumber(s string) int {
    var result int
    i := len(s)-1
    curr := 1
    for i >= 0 {
        if i == len(s)-1 {
            result = (int(s[i]) - int('A')) + 1
        } else {
            result += curr * 26 * ((int(s[i]) - int('A')) + 1)
            curr *= 26
        }
        i -= 1
    }
    return result
}

