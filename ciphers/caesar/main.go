package main

import (
	"bufio"
	"fmt"
	"os"
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
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println("Enter shift: ")
	var shift int
	fmt.Scanf("%d", &shift)

	encrypted := encrypt(text, shift)
	decrypted := decrypt(encrypted, shift)

	fmt.Println("Original text: ", text)
	fmt.Println("Encrypted text: ", encrypted)
	fmt.Println("Decrypted text: ", decrypted)
}
