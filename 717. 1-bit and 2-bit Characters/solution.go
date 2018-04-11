func isOneBitCharacter(bits []int) bool {
    nlimit := len(bits)
    v := false

    for k := 0; k < nlimit; k ++{
        if bits[k] == 1{
            v = false
            k ++
        } else{
            v = true
        }
    }
    return v

}