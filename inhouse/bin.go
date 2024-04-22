package inhouse

import (
	"fmt"
	"strconv"
	"strings"
)

func Bin(s string) string {
	// Use Fields to split the string
	word := strings.Fields(s)
	for i := 0; i < len(word); i++ {
		// Check for "(bin)"
		if word[i] == "(bin)" && i > 0 {
			bin := word[i-1]
			decimalNum, err := strconv.ParseInt(bin, 2, 64)
			if err == nil {
				word[i-1] = fmt.Sprintf("%d", decimalNum)
				// Remove the "(bin)"
				word = append(word[:i], word[i+1:]...)
				// Adjust the index after removal.
				i--
			}
		}
	}
	return strings.Join(word, " ") // Join the words with spaces.
}
