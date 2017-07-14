// Leetcode #14: Longest Common Prefix
// Write a function to find the longest common prefix string amongst an array of
// strings.
//
// brute force one by one comparison from the left took me 16 min to write,
// including test case (52 + 16 lines including this comment) and error handling

package leetcode

func longestCommonPrefix(strs []string) string {
	// min length of any string
	minLen := func(strs []string) int {
		min := len(strs[0])
		for i := 1; i < len(strs); i++ {
			l := len(strs[i])
			if l < min {
				min = l
			}
		}
		return min
	}
	// brute force one by one comparison from the left
	isSame := func(strs []string, index int) bool {
		c := strs[0][index]
		for i := 1; i < len(strs); i++ {
			if c != strs[i][index] {
				return false
			}
		}
		return true
	}

	// Sanity checks
	if len(strs) == 0 {
		return ""
	}
	min := minLen(strs)
	if min == 0 {
		return ""
	}

	for i := 0; i < min; i++ {
		if !isSame(strs, i) {
			return strs[0][0:i]
		}
	}
	return strs[0][0:min]
}
