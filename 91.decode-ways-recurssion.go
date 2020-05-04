// A message containing letters from A-Z is being encoded to numbers using the following mapping:

// 'A' -> 1
// 'B' -> 2
// ...
// 'Z' -> 26

// Recurssive solution

func numDecodings(s string) int {
    result := 0
    if string(s[0]) == "0" {
        return 0
    }
    return numDecodingsUtil(s, len(s), result)
}

func numDecodingsUtil(s string, n int, result int) int {
    if n == 0 || n == 1 {
        return 1
    } else {
        result = 0
        if string(s[n-1]) > "0" {
            result += numDecodingsUtil(s, n-1, result)
        }
        if string(s[n-2]) == "1" || (string(s[n-2]) == "2" && string(s[n-1]) < "7") {
            result += numDecodingsUtil(s, n-2, result)
        }
        return result
    }
}
