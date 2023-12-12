package main

/*
To use this program you should remove two "condition" lines in a file $GOROOT/src/crypto/aes/cipher_asm.go:
	--- if !supportsAES {
          return newCipherGeneric(key)
	--- }

and change "case" line in a file $GOROOT/src/crypto/aes/cipher.go:

	func NewCipher(key []byte) (cipher.Block, error) {
	...
	--- case 16, 24, 32:
	+++ case 16, 24, 32, 128:
	...
	}

*/

import (
	"fmt"
	"encoding/hex"
	"crypto/sha512"
	"os"
	"hash"
	"strings"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"errors"
	"bufio"
 )

var pz_salt []byte
var pz_cipherBytes []byte
var pz_cipherBytesLen int



    // PZL12 msg
const msg = "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"

func TiamatDecodeCheck(hasher hash.Hash, pass string) string {
	r := sha512.Sum512([]byte(pass))
	
	for in := 0; in < 11512; in++ {
	r = sha512.Sum512(r[:])
	}
	pwd := make([]byte, hex.EncodedLen(64))
	hex.Encode(pwd, r[:])
	derivedKeyBytes := []byte{}
	bx := []byte{}
	for len(derivedKeyBytes) < 144 {
		if len(bx) > 0 {
			hasher.Write(bx)
		}
		hasher.Write(pwd)
		hasher.Write(pz_salt)
		bx = hasher.Sum(nil)
		hasher.Reset()

		for i := 1; i < 10000; i++ {
			hasher.Write(bx)
			bx = hasher.Sum(nil)
			hasher.Reset()
		}
		derivedKeyBytes = append(derivedKeyBytes, bx...)
	}
	block, err := aes.NewCipher(derivedKeyBytes[:128])
	if err != nil {
		panic(err)
	}
	var cp []byte = make([]byte, pz_cipherBytesLen)
	copy(cp, pz_cipherBytes)
	mode := cipher.NewCBCDecrypter(block, derivedKeyBytes[128:])
	mode.CryptBlocks(cp, cp)
	length := len(cp)
	unpadding := int(cp[length-1])
	endp := string(cp[:(length - unpadding)])

	const search = "\"kty\":\"RSA\""

	if x := strings.Contains(endp, search); x == true {
            return "1"
	} else {
	    return "0"
	}
}

func b64toBinary() {
	data, err := base64.StdEncoding.DecodeString(msg)
	if err != nil {
	panic(errors.New("base64 invalid"))
	}
	if string(data[:8]) != "Salted__" {
		panic(errors.New("Invalid data"))
	}
	pz_salt = data[8:16]
	pz_cipherBytes = data[16:]
	pz_cipherBytesLen = len(pz_cipherBytes)
}

func main() {
    b64toBinary()
    h := md5.New()

    scanner := bufio.NewScanner(os.Stdin)
    var result string
    for scanner.Scan() {
        arg := scanner.Text()
        result += TiamatDecodeCheck(h,arg)
    }

    fmt.Println(result)

}
