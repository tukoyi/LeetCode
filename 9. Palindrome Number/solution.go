func isPalindrome(x int) bool {
    listx := []rune(strconv.Itoa(x))
    for i := 0; i < len(listx) / 2; i ++ {
        if listx[i] != listx[len(listx) - 1 - i]{
            return false
        }    
    }
    return true
}