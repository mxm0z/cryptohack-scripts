/*
Below is a series of outputs where three random keys have been XOR'd together and with the flag.
Use the above properties to undo the encryption in the final line to obtain the flag.

KEY1 = a6c8b6733c9b22de7bc0253266a3867df55acde8635e19c73313
KEY2 ^ KEY1 = 37dcb292030faa90d07eec17e3b1c6d8daf94c35d4c9191a5e1e
KEY2 ^ KEY3 = c1545756687e7573db23aa1c3452a098b71a7fbf0fddddde5fc1
FLAG ^ KEY1 ^ KEY3 ^ KEY2 = 04ee9855208a2cd59091d04767ae47963170d1660df7f56f5faf

Before you XOR these objects, be sure to decode from hex to bytes.
*/

package main

import (
	"encoding/hex"
	"fmt"
)

func xorKeys(b ...[]byte) []byte {

	b_len := len(b[0])
	for _, m := range b {
		if len(m) != b_len {
			panic("length mismatch!")
		}
	}
	br := make([]byte, b_len)
	for i := range b[0] {
		br[i] = 0
		for _, m := range b {
			br[i] = br[i] ^ m[i]
		}
	}
	return br
}

func main() {
	k1, _ := hex.DecodeString("a6c8b6733c9b22de7bc0253266a3867df55acde8635e19c73313")
	k2_k1, _ := hex.DecodeString("37dcb292030faa90d07eec17e3b1c6d8daf94c35d4c9191a5e1e")
	k2_k3, _ := hex.DecodeString("c1545756687e7573db23aa1c3452a098b71a7fbf0fddddde5fc1")
	flag_k1_k3_k2, _ := hex.DecodeString("04ee9855208a2cd59091d04767ae47963170d1660df7f56f5faf")

	k2 := xorKeys(k1, k2_k1)
	k3 := xorKeys(k2, k2_k3)
	masterKey := xorKeys(k2_k1, k3)
	flag := xorKeys(flag_k1_k3_k2, masterKey)

	fmt.Println(string(flag))

	//k2 = k1 ^ k2_k1
	//k3 = k2 ^ k2_k3
	//masterKey = k1 ^ k2 ^ k3
	//flag = flag_k1_k3_k2 ^ masterKey
}
