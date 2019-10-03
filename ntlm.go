package main

import (
	"unicode/utf16"

	"golang.org/x/crypto/md4"
)

func nthash(Password string) (string, error) {

	b := utf16le(Password)
	Password = string(b[:])
	bytes := callHashFactory(Password, md4.New)
	return string(bytes), nil
}

func utf16le(s string) []byte {
	codes := utf16.Encode([]rune(s))
	b := make([]byte, len(codes)*2)
	for i, r := range codes {
		b[i*2] = byte(r)
		b[i*2+1] = byte(r >> 8)
	}
	return b
}
