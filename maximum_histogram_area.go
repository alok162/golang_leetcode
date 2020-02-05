func largestRectangleArea(heights []int) int {
    var st []int
    maxArea := 0
    index := 0
    for index < len(heights)  {
        if len(st) == 0 || heights[st[len(st)-1]] < heights[index] {
            st = append(st, index)
            index += 1
        } else {
            top := st[len(st)-1]
            st = st[:len(st)-1]
            if len(st) > 0 {
                maxArea = max(maxArea, heights[top]*(index-st[len(st)-1]-1))
            } else { 
                maxArea = max(maxArea, heights[top]*index)
            }
        }
    }
    for len(st) != 0 {
        top := st[len(st)-1]
        st = st[:len(st)-1]
        if len(st) > 0 {
            maxArea = max(maxArea, heights[top]*(index-st[len(st)-1]-1))
            } else { 
                maxArea = max(maxArea, heights[top]*index)
            }
    }
    return maxArea
}

func max(data1 int, data2 int) int {
    if data1 >= data2 {
        return data1
    }
    return data2
}
