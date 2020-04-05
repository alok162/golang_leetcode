import 	"strconv"
func letterCombinations(digits string) []string {
    dict := map[int64]string{1:"", 2:"abc", 3:"def", 4:"ghi", 5:"jkl", 6:"mno",                                                7:"pqrs", 8:"tuv", 9:"wxyz"}
    var output []string
    var results []string
    if len(digits) == 0 {return results}
    util(dict, digits, &output, 0, &results)
    return results
}
func util(dict map[int64]string, digits string, output *[]string, k int, results *[]string) {
    if len(digits) == k {
        s := ""
        for _, val := range (*output) {
            s += val
        }
        (*results) = append((*results), s)
        return
    } 
    num, _ := strconv.ParseInt(string(digits[k]), 10, 32)
    for _, val := range dict[num] {
        (*output) = append((*output), string(val))

        util(dict, digits, output, k+1, results)
        x := len((*output))
        (*output) = (*output)[:x-1]
    }
}
