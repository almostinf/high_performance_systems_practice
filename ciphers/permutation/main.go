package main

import "fmt"

func encrypt(text string, shift int) string {
	if shift > len(text) {
		return text
	}

	parts := make([][]byte, 0, len(text)/shift+1)
	for i := 0; i < len(text); i++ {
		if i+shift < len(text) {
			parts = append(parts, []byte(text[i:i+shift]))
		} else {
			parts = append(parts, []byte(text[i:]))
		}
		i += shift - 1
	}

	for i := range parts {
		fmt.Println(string(parts[i]))
	}

	encrypted := make([]byte, 0, len(text))

	curSymbPos := 0
	for {
		exit := true
		for _, part := range parts {
			if curSymbPos < len(part) {
				encrypted = append(encrypted, part[curSymbPos])
				exit = false
			}
		}

		if exit {
			break
		}

		curSymbPos++
	}

	return string(encrypted)
}

func decrypt(text string, shift int) string {
	if shift > len(text) {
		return text
	}

	var bigPart int
	var numBigPart int
	if len(text)%shift == 0 {
		bigPart = len(text) / shift
		numBigPart = len(text) / shift
	} else {
		bigPart = len(text)/shift + 1
		numBigPart = len(text) % shift
	}

	parts := make([][]byte, 0, bigPart)
	for i := 0; i < len(text); i++ {
		var curPart int
		if numBigPart != 0 {
			curPart = bigPart
			numBigPart--
		} else {
			curPart = bigPart - 1
		}

		if i+curPart < len(text) {
			parts = append(parts, []byte(text[i:i+curPart]))
		} else {
			parts = append(parts, []byte(text[i:]))
		}

		i += curPart - 1
	}

	decrypted := make([]byte, 0, len(text))

	curSymbPos := 0
	for {
		exit := true
		for _, part := range parts {
			if curSymbPos < len(part) {
				decrypted = append(decrypted, part[curSymbPos])
				exit = false
			}
		}

		if exit {
			break
		}

		curSymbPos++
	}

	return string(decrypted)
}

func main() {
	text := "abcdefgih"
	encrypted := encrypt(text, 3)
	decrypted := decrypt(encrypted, 3)
	fmt.Println(encrypted)
	fmt.Println(decrypted)
}
