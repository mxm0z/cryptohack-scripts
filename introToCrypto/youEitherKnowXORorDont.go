/*
You either know, XOR you don't

I've encrypted the flag with my secret key, you'll never be able to guess it.

Remember the flag format and how it might help you in this challenge!

0e0b213f26041e480b26217f27342e175d0e070a3c5b103e2526217f27342e175d0e077e263451150104
*/

package main

import (
	"encoding/hex"
	"fmt"
)

// xorBytes XORs two byte slices and returns the result.
func xorBytes(a, b []byte) []byte {
	minLength := len(a)
	if len(b) < minLength {
		minLength = len(b)
	}
	result := make([]byte, minLength)
	for i := 0; i < minLength; i++ {
		result[i] = a[i] ^ b[i]
	}
	return result
}

// repeatKey repeats the key until it matches the length of the data.
func repeatKey(key []byte, length int) []byte {
	repeated := make([]byte, length)
	for i := 0; i < length; i++ {
		repeated[i] = key[i%len(key)]
	}
	return repeated
}

func main() {
	// Encrypted message in hex
	encryptedHex := "0e0b213f26041e480b26217f27342e175d0e070a3c5b103e2526217f27342e175d0e077e263451150104"
	encryptedBytes, _ := hex.DecodeString(encryptedHex)

	// Known start of the flag
	knownBeginPlaintext := "crypto{"
	knownBeginBytes := []byte(knownBeginPlaintext)

	// Derive the beginning of the key
	partialKey := xorBytes(knownBeginBytes, encryptedBytes[:len(knownBeginBytes)])
	fmt.Printf("Derived partial key: %s\n", string(partialKey))

	// Workaround to confirm if the derived partial key is correct
	fmt.Println("Complete or repeat the correct partial key:")
	fmt.Scan(&partialKey)

	// Assume the key repeats; extend the key
	fullKey := repeatKey(partialKey, len(encryptedBytes))

	// Decrypt the entire message
	decryptedBytes := xorBytes(encryptedBytes, fullKey)
	decryptedMessage := string(decryptedBytes)

	// Print the decrypted message
	fmt.Println("Decrypted message:", decryptedMessage)
}
