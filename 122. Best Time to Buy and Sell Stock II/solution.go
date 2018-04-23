package main 

import "fmt"

func maxProfit(prices []int) int {
    if len(prices) < 2{
        return 0
    }

    var profit int
    
    for i := 1; i < len(prices); i ++ {
        if prices[i] - prices[i -1] > 0{
            profit +=  prices[i] - prices[i -1]      
        }                   
    }
    return profit
}

func main() {
    a := []int{1,2,3,1,5}
    fmt.Println(maxProfit(a))
}