func lengthOfLongestSubstring(s string) int {
  
    var start int
    var length int
    indexMapper := make(map[rune]int)
    str := []rune(s)
    
    for i, ch := range str {
        if index, ok:= indexMapper[ch]; ok && index >= start {
            start = index+1
        }
        indexMapper[ch] = i
        if (i+1-start) > length {
            length = (i+1-start)
    }
    }
    
    return length
}