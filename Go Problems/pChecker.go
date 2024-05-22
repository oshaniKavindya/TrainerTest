/*package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	strength := checkPasswordStrength(password)

	fmt.Printf( strength)
}

func checkPasswordStrength(password string) string {
	var hasUpper, hasLower, hasNumber bool
	var strength string

	if len(password) < 8 {
		return "password must be at least 8 characters long"
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		}
	}

	if hasUpper && hasLower && hasNumber {
		strength = "Your Password is a strong one"
	} else if hasUpper || hasLower || hasNumber {
		strength = "password should contain at least one uppercase letter, one lowercase letter, and one number"
	} else {
		strength = "password should contain at least one uppercase letter, one lowercase letter, and one number"
	}

	return strength
}*/