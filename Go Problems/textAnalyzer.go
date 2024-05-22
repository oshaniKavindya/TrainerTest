/*package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a text: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	fmt.Print("Enter a letter to count: ")
	letter, _ := reader.ReadString('\n')
	letter = strings.TrimSpace(letter)

	charCount := countCharacters(text)
	wordCount := countWords(text)
	letterCount := countSpecificLetter(text, letter)

	fmt.Printf("Number of characters: %d\n", charCount)
	fmt.Printf("Number of words: %d\n", wordCount)
	if len(letter) == 1 {
		fmt.Printf("Occurrences of '%s': %d\n", letter, letterCount)
	} else {
		fmt.Println("Invalid input for letter count. Please enter a single letter.")
	}
}

func countCharacters(text string) int {
	return len(text)
}

func countWords(text string) int {
	words := strings.Fields(text) //Fields is used to split the input text into a slice of words
	return len(words)
}

func countSpecificLetter(text, letter string) int {
	count := 0
	for _, char := range text {			//The range keyword in Go is used to iterate over elements in a variety of data structures, such as slices, arrays, maps, channels, and strings. 
		if strings.ToLower(string(char)) == strings.ToLower(letter) {
			count++
		}
	}
	return count
}
*/