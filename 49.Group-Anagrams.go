// Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
// Output:
// [
//   ["ate","eat","tea"],
//   ["nat","tan"],
//   ["bat"]
// ]

func groupAnagrams(strs []string) [][]string {
    var ans [][]string
    dict := make(map[string][]string)
    
    for _, str := range(strs) {
        tempArr := [26]int{}
        for _, char := range str {
            tempArr[char-'a'] ++
        }
        tempStr := ""
        for _, val := range tempArr {
            tempStr += strconv.Itoa(val)
        }
        dict[tempStr] = append(dict[tempStr], str)
    }
    for _, val := range dict {
        ans = append(ans, val)
    }
    return ans
}
