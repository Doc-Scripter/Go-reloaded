package main

import (
	"fmt"
	"os"
	"strings"

	"go-reloaded/inhouse"
)

func main() {
	inputFile := "sample.txt"
	outputFile := "result.txt"

	// Read the input file
	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Convert the data to a string
	textData := string(data)

	lines := strings.Split(textData, "\n")

	// Initialize a slice to hold the processed lines
	var processedLines []string

	// Process each line
	for _, line := range lines {
		// Initialize a variable to keep track of the previous iteration's result
		var previousResult string

		// Apply corrections until the result can't be changed further
		for {
			// Apply all modifications
			Result := inhouse.Punctuation(line)
			Result = inhouse.Vowel(Result)
			Result = inhouse.Hex(Result)
			Result = inhouse.Bin(Result)
			Result = inhouse.Modify(Result)

			// Check if the result has changed from the previous iteration
			if Result == previousResult {
				// No more corrections can be made, exit the loop
				break
			}

			// Update the previousResult and line for the next iteration
			previousResult = Result
			line = Result
		}

		// Add the processed line to the slice
		processedLines = append(processedLines, previousResult)
	}

	// Join the processed lines back into a single string
	finalResult := strings.Join(processedLines, "\n")

	// Write the final result to the output file
	err = os.WriteFile(outputFile, []byte(finalResult), 0o644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
}
