# Elfscrow Decryption - solution in Go

## Intro

This small repo contains my go code for solving the Elfscro file decryption challenge of the KringleCon II as part of SANS Holiday Hack Challenge in 2019.

## Usage

```go
go run des-go.go
```

This will grab the encrypted pdf file from the same directory and run DES-CBC decrypt on it in a loop. 

On every iteration of the loop counter a specific UNIX timestamp between 7-9 pm December 6 2019 (1575658800 - 1575666000) is fed into the key generator that will be used to decrypt. 

If the resulting plain-text starts with %PDF then we save it and exit, knowing that we found the decrypted file contents that translate to a valid pdf. 

The file was encrypted on 12/06/2019 @ 8:20pm (UTC) - 1575663650.

## Credits

For the original skeleton code for DES-CBC implementation in GO:
https://stackoverflow.com/questions/41579325/golang-how-do-i-decrypt-with-des-cbc-and-pkcs7
