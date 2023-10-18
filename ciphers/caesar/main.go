package main

import (
	"fmt"
	"strings"
)

func encrypt(text string, shift int) string {
	var result strings.Builder
	for _, char := range text {
		if char >= 'A' && char <= 'Z' {
			shifted := 'A' + (char-'A'+rune(shift))%26
			result.WriteString(string(shifted))
		} else if char >= 'a' && char <= 'z' {
			shifted := 'a' + (char-'a'+rune(shift))%26
			result.WriteString(string(shifted))
		} else {
			result.WriteString(string(char))
		}
	}
	return result.String()
}

func decrypt(text string, shift int) string {
	return encrypt(text, 26-shift)
}

func main() {
	text := "Hello, world!"
	shift := 3

	encrypted := encrypt(text, shift)
	decrypted := decrypt(encrypted, shift)

	fmt.Println("Original text: ", text)
	fmt.Println("Encrypted text: ", encrypted)
	fmt.Println("Decrypted text: ", decrypted)
}
