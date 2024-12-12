func pickGifts(gifts []int, k int) int64 {
    
    sort.Ints(gifts)
    
    var res int64
    
    for i := len(gifts)-1;;{
        if k == 0 {
            break
        }
        gifts[i] = int(math.Floor(math.Sqrt(float64(gifts[i]))))
                       sort.Ints(gifts)
        k --
    }
    
                       for j := 0; j < len(gifts); j++{
                           res += int64(gifts[j])
                       }
                       
                       return res
    
}