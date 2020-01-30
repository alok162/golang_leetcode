import "fmt"

func findAnagrams(s string, p string) []int {
    k := len(p)-1
    start := 0
    var res []int
     if len(s) < len(p) {
        return res
    }
    dict := make(map[string]int)
    temp := make(map[string]int)
    for _, char := range p {
        if temp[string(char)] == 0 {
            temp[string(char)] = 1
        } else {
            temp[string(char)] += 1
        }
    }
    i := 0
    for i < k {
        if dict[string(s[i])] == 0 {
            dict[string(s[i])] = 1
        } else {
            dict[string(s[i])] += 1
        }
        i += 1
    }
    for i < len(s) {
         if dict[string(s[i])] == 0 {
            dict[string(s[i])] = 1
        } else {
            dict[string(s[i])] += 1
        }
        if count(dict, temp) == true {
            res = append(res, start)
        }
        if dict[string(s[start])] == 0 {
            delete(dict, string(s[start]))
        } else {
            dict[string(s[start])] -= 1
        }
        start += 1
        i += 1
    }
    return res
}

func count(dict map[string]int, temp map[string]int) bool {
    for key := range dict {
        if dict[string(key)] != temp[string(key)] {
            return false
        }
    }
    return true
}
