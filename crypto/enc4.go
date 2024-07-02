//  https://cryptohack.org/courses/intro/enc4/

/**
Cryptosystems like RSA works on numbers, but messages are made up of characters. How should we convert our messages into numbers so that mathematical operations can be applied?
The most common way is to take the ordinal bytes of the message, convert them into hexadecimal, and concatenate.
This can be interpreted as a base-16/hexadecimal number, and also represented in base-10/decimal.
**/
// You can edit this code!
// Click here and start typing.
package main

import (
	"encoding/hex"

	"github.com/kenshaw/baseconv"

	"fmt"
)

func main() {

	str := []byte("HELLO")
	dst := make([]byte, hex.EncodedLen(len(str)))
	hex.Encode(dst, str)
	fmt.Printf("%s\n", dst)

	hexStr := string(dst)
	fmt.Println(hexStr)
	valDec, _ := baseconv.DecodeHexToDec(hexStr)
	fmt.Println(valDec)
	encodedStr := hex.EncodeToString([]byte(str))
	fmt.Println(encodedStr)

	strToDecode := "11515195063862318899931685488813747395775516287289682636499965282714637259206269"

	hexStr, _ = baseconv.EncodeHexFromDec(strToDecode)
	fmt.Println(hexStr)
	decoded, _ := hex.DecodeString(hexStr)
	fmt.Printf("\n%s", decoded)

}
