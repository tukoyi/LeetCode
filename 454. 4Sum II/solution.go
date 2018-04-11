//beat 37%，时间复杂度n2
func fourSumCount(A []int, B []int, C []int, D []int) int {
    var cnt int
    sum := make(map[int]int)
    for i := 0; i < len(A); i ++{
        for j := 0; j < len(A); j ++{
            if _, ok := sum[A[i] + B[j]]; ok{
                sum[A[i] + B[j]] += 1
            } else {
                sum[A[i] + B[j]] = 1
            }
        }
    }
    fmt.Println(sum)
    for i := 0; i < len(C); i ++{
        for j := 0; j < len(D); j ++{    
            if _, ok := sum[- 1 * C[i] - D[j]]; ok{
                // 这里cnt是cnt + map[k]
                cnt = cnt + sum[- 1 * C[i] - D[j]]
            }
        }
    }    
    return cnt
}