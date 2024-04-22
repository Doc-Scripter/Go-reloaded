package inhouse

import (
	"fmt"
	"strconv"
	"strings"
)

func Hex(s string) string {
	words := strings.Fields(s)
	for i := 0; i < len(words); i++ {
		// Check and remove "(hex)" and adjust the index.
		if words[i] == "(hex)" && i > 0 {
			hex := words[i-1]
			decimalNum, err := strconv.ParseInt(hex, 16, 64)
			if err == nil {
				words[i-1] = fmt.Sprintf("%d", decimalNum)
				words = append(words[:i], words[i+1:]...)
				i--
			}
			if err != nil {
				fmt.Println("Error converting to string")
			}
		}
	}
	return strings.Join(words, " ")
}
