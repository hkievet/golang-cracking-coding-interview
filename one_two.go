package interview

func IsPermutation(s1, s2 string) bool {
	runes1, runes2 := []rune(s1), []rune(s2)
	// They aren't permutations if they are different lengths
	if len(runes1) != len(runes2) {
		return false
	}
	for _, r := range runes1 {
		found := false
		for i, check := range runes2 {
			if r == check {
				// remove the value from runes2 so it cannot count as a match in future iterations
				runes2 = append(runes2[:i], runes2[i+1:]...)
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
