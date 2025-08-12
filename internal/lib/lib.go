package lib

import (
	"strings"
)

func TitleCase(str string) string {
	if str == "" {
		return ""
	}
	words := strings.Split(strings.ToLower(strings.ReplaceAll(str, "_", " ")), " ")
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(string(word[0])) + word[1:]
		}
	}
	return strings.Join(words, " ")
}
