package june_challenge

func isSubsequence(s string, t string) bool {
	text, pattern := []byte(t), []byte(s)
	if len(pattern) == 0 {
		return true
	}
	p := 0
	for i := range text {
		if text[i] == pattern[p] {
			p++
		}
		if p == len(s) {
			return true
		}
	}
	return false
}
