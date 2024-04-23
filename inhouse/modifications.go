package inhouse

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Capitalize(toCap string) string {
	for i, v := range toCap {
		return string(unicode.ToUpper(v)) + toCap[i+1:]
	}
	return " "
}

func Low(toLow string) string {
	return strings.ToLower(toLow)
}

func Up(toUp string) string {
	return strings.ToUpper(toUp)
}

func Modify(s string) string {
	word := strings.Fields(s)
	for i := 0; i < len(word); i++ {
		var f func(string) string
		if strings.HasPrefix(word[i], "(cap") {
			f = Capitalize
		} else if strings.HasPrefix(word[i], "(low") {
			f = Low
		} else if strings.HasPrefix(word[i], "(up") {
			f = Up
		} else {
			continue
		}
		if strings.Contains(word[i], ",") {
			// Extract the number after the comma
			numStr := strings.Trim(word[i+1], ")")
			number, err := strconv.Atoi(numStr)
			if number > len(s[:i]) {
				fmt.Println("The words to be modified are less than specified")
			}
			if number <= i {
				// Apply the function 'f' to the preceding 'number' of words
				for j := i - number; j < i; j++ {
					word[j] = f(word[j])
				}
			}
			if err != nil {
				fmt.Println("Error converting to interger")
			}
			// Remove the current word
			word = append(word[:i], word[i+2:]...)
			i--
		} else {
			subString := []string{"(cap)", "(low)", "(up)"}
			for _, subStr := range subString {
				if strings.Contains(strings.ToLower(word[i]), subStr) {
					// Apply the function 'f' to the preceding word
					if i > 0 {
						word[i-1] = f(word[i-1])
					}
					// Remove the current word
					word = append(word[:i], word[i+1:]...)
					i--
				}
			}
		}
	}
	return strings.Join(word, " ")
}
