package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
)

// e, n - public
// d, n - private

func rsaKeygen(bits int) (n, e, d *big.Int, err error) {
	// Generate two large prime numbers p and q
	p, err := rand.Prime(rand.Reader, bits)
	if err != nil {
		return nil, nil, nil, err
	}

	q, err := rand.Prime(rand.Reader, bits)
	if err != nil {
		return nil, nil, nil, err
	}

	// n = p * q
	n = new(big.Int).Mul(p, q)

	// phi(n) = (p - 1) * (q - 1)
	phi := new(big.Int).Mul(new(big.Int).Sub(p, big.NewInt(1)), new(big.Int).Sub(q, big.NewInt(1)))

	// used default exp
	e = big.NewInt(65537)

	// compute exponent d = e ^ (-1) mod phi(n)
	d = new(big.Int)
	d.ModInverse(e, phi)

	return n, e, d, nil
}

func rsaEncrypt(plaintext, n, e *big.Int) *big.Int {
	ciphertext := new(big.Int)
	ciphertext.Exp(plaintext, e, n) // (plaintext ^ e) mod n
	return ciphertext
}

func rsaDecrypt(ciphertext, n, d *big.Int) *big.Int {
	plaintext := new(big.Int)
	plaintext.Exp(ciphertext, d, n) // (ciphertext ^ d) mod n
	return plaintext
}

func textToBigInt(text string) *big.Int {
	result := big.NewInt(0)
	for _, char := range text {
		result = result.Mul(result, big.NewInt(256)) // Assuming UTF-8 encoding
		result = result.Add(result, big.NewInt(int64(char)))
	}
	return result
}

func bigIntToText(number *big.Int) string {
	result := ""
	for number.Sign() > 0 {
		char := rune(number.Int64() % 256) // Assuming UTF-8 encoding
		result = string(char) + result
		number = number.Div(number, big.NewInt(256))
	}
	return result
}

func main() {
	bits := 1024
	n, e, d, err := rsaKeygen(bits)
	if err != nil {
		log.Fatal("Error generating keys: %w", err)
	}

	fmt.Println("Example with number: ")

	num := big.NewInt(999)
	fmt.Println("Plaintext: ", num.String())

	ciphertext := rsaEncrypt(num, n, e)
	fmt.Println("Encrypted: ", ciphertext.String())

	decrypted := rsaDecrypt(ciphertext, n, d)
	fmt.Println("Decrypted: ", decrypted.String())

	fmt.Println("\nExample with text: ")

	text := "Hello, RSA!"
	textBigInt := textToBigInt(text)
	fmt.Println("Plaintext: ", text)

	ciphertext = rsaEncrypt(textBigInt, n, e)
	fmt.Println("Encrypted: ", ciphertext.String())

	decrypted = rsaDecrypt(ciphertext, n, d)
	fmt.Println("Decrypted: ", bigIntToText(decrypted))
}
