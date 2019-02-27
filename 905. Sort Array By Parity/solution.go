//Runtime: 68 ms, faster than 100.00% of Go online submissions for Sort Array By Parity.
//Memory Usage: 8 MB, less than 86.11% of Go online submissions for Sort Array By Parity.
//Next challenges:

// 这里用到了golang的切片拼接特性。不考虑该特性，可使用首尾两边逼近交换的思想

func sortArrayByParity(A []int) []int {
    var oddArr, evenArr, ret  []int
    for _, v := range A {
        if v % 2 == 0 {
            evenArr = append(evenArr, v)
        } else {
            oddArr = append(oddArr, v)
        }
    }
    
    ret = append(evenArr, oddArr...)
    return ret
}