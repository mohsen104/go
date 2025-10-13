package utils

import "strings"

func LengthOfLastWord(s string) int {
	words := strings.Split(strings.TrimSpace(s), " ")
	return len(words[len(words)-1])
}
