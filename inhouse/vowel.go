package inhouse

import (
	"strings"
)

func Vowel(s string) string {
	word := strings.Split(s, " ")

	vowels := []string{"a", "e", "i", "o", "u", "h"}
	for i := 0; i < len(word); i++ {
		// to lowercase for comparison
		if word[i] == "a" {
			// access the first letter in the next word
			firstLetter := strings.ToLower(string(word[i+1][0]))
			for _, vowel := range vowels {
				if firstLetter == vowel {
					word[i] = "an"
					break
				}
			}
		} else if word[i] == "A" {
			firstLetter := strings.ToLower(string(word[i+1][0]))
			for _, vowel := range vowels {
				if firstLetter == vowel {
					word[i] = "An"
					break
				}
			}
		}
	}
	return strings.Join(word, " ")
}
