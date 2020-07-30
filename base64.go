// Another common encoding scheme is Base64, which allows us to represent binary data as an ASCII string using 64 characters. One character of a Base64 string encodes 6 bits, and so 4 characters of Base64 encodes three 8-bit bytes.
// Base64 is most commonly used online, where binary data such as images can be easy included into html or css files.
// Take the below hex string, decode it into bytes and then encode it into Base64.
// 72bca9b68fc16ac7beeb8f849dca1d8a783e8acf9679bf9269f7bf

package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	hexString := "72bca9b68fc16ac7beeb8f849dca1d8a783e8acf9679bf9269f7bf"
	hexByte, _ := hex.DecodeString(hexString)
	hex64 := base64.StdEncoding.EncodeToString([]byte(hexByte))
	fmt.Println(string(hex64))
	//decodedByte, _ := base64.StdEncoding.DecodeString(encodedBase64)
	//fmt.Println(string(decodedByte))
}
