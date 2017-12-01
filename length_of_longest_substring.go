package leetcode

// LengthOfLongestSubstring returns number of bytes (not runes)
func LengthOfLongestSubstring(s string) int {
	max := 0
	for j := 0; j < len(s); j++ {
		m := make(map[byte]bool)
		for i := j; i < len(s); i++ {
			c := s[i]
			if m[c] {
				break
			}
			m[c] = true
		}
		if len(m) > max {
			max = len(m)
		}
	}
	return max
}
