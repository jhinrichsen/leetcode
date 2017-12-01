package leetcode

import "strings"

func isValid(s string) bool {
	var pars = []string{"()", "[]", "{}"}
	// Lets not get carried away with PUSHs and POPs
	for changed := true; changed; {
		before := s
		for _, par := range pars {
			s = strings.Replace(s, par, "", -1)
		}
		changed = before != s
	}
	return len(s) == 0
}
