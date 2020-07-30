// Cryptosystems like RSA works on numbers, but messages are made up of characters. How should we convert our messages into numbers so that mathematical operations can be applied?
// The most common way is to take the ordinal bytes of the message, convert them into hexadecimal, and concatenate. This can be interpreted as a base-16 number, and also represented in base-10.
// To illustrate:
// message: HELLO
// ascii bytes: [72, 69, 76, 76, 79]
// hex bytes: [0x48, 0x45, 0x4c, 0x4c, 0x4f]
// base-16: 0x48454c4c4f
// base-10: 310400273487
// (48454C4C4F)₁₆ = ((4 × 16⁹) + (8 × 16⁸) + (4 × 16⁷) + (5 × 16⁶) + (4 × 16⁵) + (12 × 16⁴) + (4 × 16³) + (12 × 16²) + (4 × 16¹) + (15 × 16⁰))₁₀ = (310400273487)₁₀
// Python's PyCryptodome library implements this with the methods Crypto.Util.number.bytes_to_long and Crypto.Util.number.long_to_bytes.
// Convert the following integer back into a message:
// base-10 -- 11515195063862318899931685488813747395775516287289682636499965282714637259206269
// base-16 -- 63727970746f7b336e633064316e365f346c6c5f3768335f7734795f6430776e7d
// 63 72 79 70 74 6f 7b 33 6e 63 30 64 31 6e 36 5f 34 6c 6c 5f 37 68 33 5f 77 34 79 5f 64 30 77 6e 7d
// https://gchq.github.io/CyberChef/

package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	hexadecimal := "63727970746f7b336e633064316e365f346c6c5f3768335f7734795f6430776e7d"
	data, _ := hex.DecodeString(hexadecimal)
	fmt.Println(string(data))
}
