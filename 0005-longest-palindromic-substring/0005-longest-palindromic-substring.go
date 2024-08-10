func longestPalindrome(s string) string {
	var max string

	if len(s) == 1 {
		return s
	}
	if len(s) == 0 {
		return ""
	}
	for i := 0; i < len(s); i++ {
		pal1 := exp(s, i, i)
		pal2 := exp(s, i, i+1)

		if len(pal1) > len(max) {
			max = pal1
		}

		if len(pal2) > len(max) {
			max = pal2
		}
	}
	return max
}

func exp(s string, left int, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}
