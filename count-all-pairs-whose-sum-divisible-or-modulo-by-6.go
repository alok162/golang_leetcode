import "fmt"
func numPairsDivisibleBy60(time []int) int {
    dict := make(map[int]int)
    for _ , key := range(time) {
        dict[key%60] += 1
    }
    result := 0
    fmt.Println(dict)
    for key, val := range(dict) {
        _, ok := dict[60-key]
        if ok || key == 0 {
            if key == 0 || 60-key == key {
                result += (val*(val-1))/2
            } else {
                result += val*dict[60-key]
                dict[key] = 0
                dict[60-key] = 0
            }
        }
    }
    return result
}
