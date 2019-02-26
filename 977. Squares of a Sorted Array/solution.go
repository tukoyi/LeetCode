//Runtime: 480 ms, faster than 100.00% of Go online submissions for Squares of a Sorted Array.
//Memory Usage: 8.6 MB, less than 18.60% of Go online submissions for Squares of a Sorted Array.
// https://leetcode.com/problems/squares-of-a-sorted-array/
// 注意len(A) < 2和index == 1时的边界条件


func sortedSquares(A []int) []int {
    if len(A) < 2 {
        return []int{squra(A[0])}
    }
    
    start := 0
    end   := len(A) - 1
    
    ret := make([]int, len(A))
    index := len(A) - 1
    
    for start < end {
        if squra(A[start]) > squra(A[end]) {
            ret[index] = squra(A[start])
            start ++
            
            if index == 1 {
                ret[0] = squra(A[end])
                break
            }
        } else {
            ret[index] = squra(A[end])
            end --
            
            if index == 1 {
                ret[0] = squra(A[start])
                break
            }
            
        }   
        
        index --
       
    }
    return ret
    
}

func squra(num int) int {
    return num * num
}