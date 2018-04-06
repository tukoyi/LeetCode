func convert(s string, numRows int) string {
    l := len(s)
    slices := []byte(s)
    var result []byte 
    var k int
    
    if numRows==1 || l <= numRows{
        return s
    }
    
    for i := 0; i < numRows; i ++{
        for j := i; j < l; j += 2*(numRows - 1){
            result =  append(result, slices[j])
            
            if i != 0 && i != numRows - 1{
                k = j + 2*(numRows - 1) - 2 * i
                if k < l{
                    result =  append(result, slices[k])
                }
                
            }
            
        }
    }
    return string(result)
}