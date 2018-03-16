package interview

func EscapeSpaces(s string) string {
	runes := []rune(s)
	newRunes := []rune{}
	for _, r := range runes {
		if r == ' ' {
			newRunes = append(newRunes, []rune("%20")...)
		} else {
			newRunes = append(newRunes, r)
		}
	}
	return string(newRunes)
}
