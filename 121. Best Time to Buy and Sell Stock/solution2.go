func maxProfit(prices []int) int {
    var profit, minPrice int
    if len(prices) < 2{
        return 0
    }
    
    minPrice = prices[0]
    for i := 0; i < len(prices); i ++ {
        if prices[i] - minPrice > profit{
            profit = prices[i] - minPrice 
        }
        if prices[i] < minPrice{
            minPrice = prices[i]
        }
    }
    return profit
}