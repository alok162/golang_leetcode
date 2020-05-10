// Given a string, find the first non-repeating character in it and return it's index. If it doesn't exist, return -1.

// Examples:

// s = "leetcode"
// return 0.

// s = "loveleetcode",
// return 2.

func firstUniqChar(s string) int {
    dict := make(map[string]int)
    for _, val := range s {
        dict[string(val)] += 1
    }
    for index, val := range s {
        if dict[string(val)] == 1 {
            return index
        }
    }
    return -1
}
