package main

import (
	"crypto/cipher"
	"crypto/des"
	"encoding/hex"
	"fmt"
	"io/ioutil"
)

// func to decrypt ciphertext with given key and IV vector
func decryptDESCBC(key, IV, cipherText []byte) ([]byte, error) {
	block, _ := des.NewCipher(key)
	blockMode := cipher.NewCBCDecrypter(block, IV)
	origData := make([]byte, len(cipherText))
	blockMode.CryptBlocks(origData, cipherText)
	return origData, nil
}

// basic linear congruential generator
func lcg(a, c, m, seed uint32) func() uint32 {
    r := seed
    return func() uint32 {
        r = (a*r + c) % m
        return r
    }
}

// microsoft generator has extra division step
func msg(seed uint32) func() uint32 {
    g := lcg(214013, 2531011, 1<<31, seed)
    return func() uint32 {
        return g() / (1 << 16)
    }
}

// LCG implt with rosetta, modified to work with this objective
func generateKeyFromSeed(seed uint32) []byte {
    var key []byte
    msf := msg(seed)
    for i := 0; i < 8; i++ {
        randHex := fmt.Sprintf("%x", msf()) // >> 0x10 & 0x7fff)
    		if len(randHex) < 2 {
    			randHex = "0" + randHex
    		}
    		valHex2Bytes := make([]byte, hex.DecodedLen(len(randHex[len(randHex)-2:])))
    		hex.Decode(valHex2Bytes, []byte(randHex[len(randHex)-2:]))
    		key = append(key, valHex2Bytes[0])
    }
    return key
}

func main() {
	// load file into []byte variable
	fileToDecrypt, err := ioutil.ReadFile("ElfUResearchLabsSuperSledOMaticQuickStartGuideV1.2.pdf.enc")
	if err != nil {
		panic(err)
	}

	// given start and end timestamps in UNIX epoch
	var start uint32 = 1575666000 //
	var end   uint32 = 1575658800

	// null IV
	IV := []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

	// loop through all timestamps
	for time := start; time > end; time -= 1 {
		// Update terminal output
		fmt.Printf("Seed: %d              \r", time)

		// generate key from seed time
		key := generateKeyFromSeed(time)
		plainText, _ := decryptDESCBC(key, IV, fileToDecrypt)

		// if first 4 characters match %PDF
		if string(plainText[:4]) == `%PDF` {
			fmt.Printf("Key-Used: %x - TimeStamp: %d - First-4-Chars: %s...\n", key, time, string(plainText[:4]))
			ioutil.WriteFile("ElfUResearchLabsSuperSledOMaticQuickStartGuideV1.2.pdf", plainText, 0644)
			break
		}
	}
}
