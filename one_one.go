package interview

func HasUniqueCharacters(s string) bool {
	runes := []rune(s)
	for i, r := range runes {
		for j, check := range runes {
			if check == r && i != j {
				return false
			}
		}
	}
	return true
}
