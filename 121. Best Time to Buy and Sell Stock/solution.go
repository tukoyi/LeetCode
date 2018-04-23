package main 

import "fmt"

func maxProfit(prices []int) int {
    var profit, buyDay, soldDay int
    if len(prices) < 2{
        return 0
    }
    
    for buyDay = 0; buyDay < len(prices) - 1; buyDay ++ {
        for soldDay = buyDay + 1; soldDay < len(prices); soldDay ++ {
            if prices[soldDay] - prices[buyDay] > profit{
                profit = prices[soldDay] - prices[buyDay]
            }
        }
    }
    return profit
}

func main() {
    a := []int{7,1,5,3,6,4}
    fmt.Println(maxProfit(a))
}