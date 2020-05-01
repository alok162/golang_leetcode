// Input: [7,1,5,3,6,4]
// Output: 5
// Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
//              Not 7-1 = 6, as selling price needs to be larger than buying price.

func maxProfit(prices []int) int {
    if len(prices) == 1 || len(prices) == 0 {
        return 0
    }
    currDiff := prices[1]-prices[0]
    currSum := currDiff
    result := currSum
    for i := 1; i < len(prices)-1; i ++ {
        currDiff = prices[i+1]-prices[i]
        if currSum < 0 {
            currSum = currDiff
        } else {
            currSum += currDiff
        }
        if currSum > result {
            result = currSum
        }
    }
    if result >= 0 {
        return result
    }
    return 0
}
