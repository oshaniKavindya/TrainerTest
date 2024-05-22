/*package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// User-defined parameters
	length := 8
	uppercase := true
	lowercase := true
	digits := true
	specialChars := true

	// Generate the random password
	password := generateRandomPassword(length, uppercase, lowercase, digits, specialChars)

	// Print the generated password
	fmt.Println("Generated Password:", password)
}

func generateRandomPassword(length int, uppercase, lowercase, digits, specialChars bool) string {
	var charset string
	if uppercase {
		charset += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if lowercase {
		charset += "abcdefghijklmnopqrstuvwxyz"
	}
	if digits {
		charset += "0123456789"
	}
	if specialChars {
		charset += "!@#$%^&*-_=+,.<>/?"
	}

	password := make([]byte, length)
	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}

	return string(password)
}*/
