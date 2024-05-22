/*package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// Define the input and output file paths
	inputFilePath := "input.txt"
	outputFilePath := "output.txt"

	// Read data from the input file
	data, err := ioutil.ReadFile(inputFilePath)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	// Count the number of lines in the input data
	numLines := countLines(string(data))

	// Write the number of lines to the output file
	err = writeResult(outputFilePath, "Number of lines", numLines)
	fmt.Println()
	if err != nil {
		fmt.Println("Error writing output file:", err)
		return
	}

	// Ask the user if they want to find specific words
	var findWords string
	fmt.Println("Do you want to find specific words in the text? (yes/no): ")
	fmt.Scanln(&findWords)

	if strings.ToLower(findWords) == "yes" {
		// Ask the user to input the words they want to find
		var wordsToFind string
		fmt.Print("Enter the words you want to find (separated by comma): ")
		fmt.Scanln(&wordsToFind)

		// Split the input words by comma and trim spaces
		searchWords := strings.Split(strings.TrimSpace(wordsToFind), ",")

		// Count occurrences of each word in the text
		wordCounts := make(map[string]int)
		for _, word := range searchWords {
			wordCounts[strings.TrimSpace(word)] = strings.Count(string(data), word)
		}

		// Write the word counts to the output file
		for word, count := range wordCounts {
			err := writeResult(outputFilePath, word, count)
			if err != nil {
				fmt.Println("Error writing output file:", err)
				return
			}
		}
		fmt.Println("Word counts written to output file:", len(wordCounts))
	}
}

// countLines counts the number of lines in the given text.
func countLines(text string) int {
	lines := strings.Split(text, "\n")
	return len(lines)
}

// writeResult writes the result to the output file.
func writeResult(filePath, resultName string, resultValue int) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the result to the file
	_, err = fmt.Fprintf(file, "%s: %d\n", resultName, resultValue)
	return err
}
*/