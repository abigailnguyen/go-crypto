//  https://cryptohack.org/courses/intro/enc4/

/*
Cryptosystems like RSA works on numbers, but messages are made up of characters. How should we convert our messages into numbers so that mathematical operations can be applied?
The most common way is to take the ordinal bytes of the message, convert them into hexadecimal, and concatenate.
This can be interpreted as a base-16/hexadecimal number, and also represented in base-10/decimal.
*/
package crypto

import (
	"encoding/hex"
	"fmt"

	"github.com/kenshaw/baseconv"
)

func Encrypted4() {
	str := []byte("HELLO")
	// make a byte array that is same length of the string to be encoded. Hex encode the string
	dst := make([]byte, hex.EncodedLen(len(str)))
	hex.Encode(dst, str)
	fmt.Printf("%s\n", dst)
	hexStr := string(dst)

	/// this is a shorter way to encode the the string to hex
	encodedStr := hex.EncodeToString([]byte(str))
	fmt.Println(encodedStr)

	// use baseconv library to convert the encoded hexstring to Decimal
	fmt.Println(hexStr)
	valDec, _ := baseconv.DecodeHexToDec(hexStr)
	fmt.Println(valDec)

	// this is the string that was encoded to bytes => hex => decimal
	strToDecode := "11515195063862318899931685488813747395775516287289682636499965282714637259206269"

	// convert the decimal string => hex => bytes => string
	hexStr, _ = baseconv.EncodeHexFromDec(strToDecode)
	fmt.Println(hexStr)
	decoded, _ := hex.DecodeString(hexStr)
	fmt.Printf("\n%s", decoded)
}
