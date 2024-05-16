/*
XOR is a bitwise operator which returns 0 if the bits are the same, and 1 otherwise. In textbooks the XOR operator is denoted by ⊕, but in most challenges and programming languages you will see the caret ^ used instead.

A	B	Output
0	0	0
0	1	1
1	0	1
1	1	0

For longer binary numbers we XOR bit by bit: 0110 ^ 1010 = 1100. We can XOR integers by first converting the integer from decimal to binary. We can XOR strings by first converting each character to the integer representing the Unicode character.

Given the string "label", XOR each character with the integer 13. Convert these integers back to a string and submit the flag as crypto{new_string}.

The Python pwntools library has a convenient xor() function that can XOR together data of different types and lengths. But first, you may want to implement your own function to solve this.
*/

package main

import "fmt"

func xorString(input string, key int32) string {
	// creates an empty slice of type int32 which the number of elements corresponds to the number of characters of the input string
	// if the input string is "label", it will have 5 elements
	xorResult := make([]int32, len(input))

	fmt.Println("XORing input string...")
	// variable i is the element position of the slice. It will iterate over L, A, B, E, L
	// variable char is the value of the element i.e. the actual character of the input string
	for i, char := range input {
		xorResult[i] = char ^ key
		fmt.Println(string(xorResult))
	}

	return string(xorResult)
}

func main() {
	input := "label"
	key := int32(13)

	output := xorString(input, key)
	fmt.Println("XOR'd result for label using 13 is:", output)
}
