package sol

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	charFreq := make([]int, 26)
	for pos := range s {
		charFreq[s[pos]-'a']++
		charFreq[t[pos]-'a']--
	}
	for _, val := range charFreq {
		if val != 0 {
			return false
		}
	}
	return true
}
