/*
For the next few challenges, you'll use what you've just learned to solve some more XOR puzzles.

I've hidden some data using XOR with a single byte, but that byte is a secret. Don't forget to decode from hex first.

73626960647f6b206821204f21254f7d694f7624662065622127234f726927756d
*/

package main

import (
	"encoding/hex"
	"fmt"
	"strings"
)

func xorWithByte(data []byte, key byte) []byte {
	result := make([]byte, len(data))
	for i, b := range data {
		result[i] = b ^ key
	}
	return result
}

func main() {
	initialHex, _ := hex.DecodeString("73626960647f6b206821204f21254f7d694f7624662065622127234f726927756d")

	for i := 0; i < 256; i++ {
		result := xorWithByte(initialHex, byte(i))
		if strings.HasPrefix(string(result), "crypto") {
			fmt.Println(string(result))
		}
	}
}
