package main

import (
    "crypto/cipher"
    "crypto/des"
    "fmt"
    "encoding/hex"
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

// func to generate the key for given seed
func generateKeyFromSeed(seed int) []byte {
    key_b := seed
    var key []byte
    for i:=0; i<8; i+=1 {
        key_b = key_b * 0x343fd + 0x269ec3
        tmp_val := fmt.Sprintf("%X", key_b >> 0x10 & 0x7fff)
        if len(tmp_val) < 2 {
            tmp_val = "0" + tmp_val
        }
        val := tmp_val[len(tmp_val)-2:]
        val_hex := make([]byte, hex.DecodedLen(len(val)))
        hex.Decode(val_hex, []byte(val))
        key = append(key, val_hex[0])
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
    start := 1575663700 //1575666000
    end   := 1575658800

    // null IV
    IV := []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
    
    // loop through all timestamps
    for time:= start; time > end; time -= 1 {    
        // Update terminal output
        fmt.Printf("Seed: %d              \r", time)

        // generate key from seed time
        key := generateKeyFromSeed(time)
        plainText, _ := decryptDESCBC(key, IV, fileToDecrypt)
        
        // if first 4 characters match %PDF
        if string(plainText[:4]) == "%PDF" {
            fmt.Printf("Key: %x - Time: %d - FILE: %s\n", key, time, string(plainText[:4]))
            ioutil.WriteFile("ElfUResearchLabsSuperSledOMaticQuickStartGuideV1.2.pdf", plainText, 0644)
            break
        }    
    }
}
