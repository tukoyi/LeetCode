func maxProfit(prices []int) int {
    if len(prices) < 2{
        return 0
    }

    var profit, valley, peak int
    valley, peak = prices[0], prices[0]
    for i := 0; i < len(prices) - 1;  {
        for i + 1 <= len(prices) - 1 && prices[i] >= prices[i + 1]{  
            
            i ++
        // valley = prices[i]不能写在这里   
        }
        valley = prices[i]
        
        for i + 1 <= len(prices) - 1 && prices[i] <= prices[i + 1]{
            
             i ++
        //peak = prices[i]不能写在这里     
        }
        peak = prices[i]
        fmt.Println(valley, peak)
        profit += peak - valley
        
    }
    return profit
}