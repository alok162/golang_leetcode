func isValid(s string) bool {
    var st []string
    dict := map[string]string{"(": ")", "{": "}", "[": "]"}
    for _, val := range(s) {
        if string(val) == "(" || string(val) == "{" || string(val) == "[" {
            st = append(st, string(val))
        } else {
            if len(st) > 0 && dict[st[len(st)-1]] == string(val) {
                st = st[:len(st)-1]
            } else {
                return false
            }
        }
    }
    
    if len(st) == 0 {
        return true
    }
    return false
}
