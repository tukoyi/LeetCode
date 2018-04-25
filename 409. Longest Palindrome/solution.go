func longestPalindrome(s string) int {
    var res int
    myMap := make(map[rune]int)
    
    lists := []rune(s)
    if len(lists) < 2{
    	return 1
    }
    //映射到map[str]cnt
    for _, v := range lists{
        if _, ok := myMap[v]; ok{
            myMap[v] += 1
        } else {
            myMap[v] = 1
        }
    }
    
    //统计字符个数大于2的字符数目
    for _, v := range myMap {
        if v / 2 >= 1 {
            res = res + v - v % 2
        }
        
    }
    
    //没有两个相同字符比如“AB”情况下，也算能产生1个字符的回文，下面的条件没有改成res != 0 && res < len([]rune(s))
    if  res < len([]rune(s)){
        res += 1
    }
    fmt.Println(myMap)
    return res
}