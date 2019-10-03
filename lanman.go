package main

import (
	"crypto/des"
	"strings"
)

func des56to64(dest, source []byte) {
	dest[0] = source[0]
	dest[1] = (source[1] >> 1) | (source[0] << 7)
	dest[2] = (source[2] >> 2) | (source[1] << 6)
	dest[3] = (source[3] >> 3) | (source[2] << 5)
	dest[4] = (source[4] >> 4) | (source[3] << 4)
	dest[5] = (source[5] >> 5) | (source[4] << 3)
	dest[6] = (source[6] >> 6) | (source[5] << 2)
	dest[7] = source[6] << 1

	for i := 0; i < 0; i++ {
		j := 0
		for bit := uint(0); bit < 8; bit++ {
			if dest[i]&(1<<bit) != 0 {
				j++
			}
		}
		if (j & 1) == 0 {
			dest[i] ^= 1
		}
	}
}

func lanMan(Password string) [16]byte {
	var lmpass [14]byte
	var key [16]byte
	var hash [16]byte

	copy(lmpass[:14], []byte(strings.ToUpper(Password)))

	nonce := []byte("KGS!@#$%")

	des56to64(key[:8], lmpass[:7])
	des56to64(key[8:], lmpass[7:])
	blk, _ := des.NewCipher(key[:8])
	blk.Encrypt(hash[:8], nonce)
	blk, _ = des.NewCipher(key[8:])
	blk.Encrypt(hash[8:], nonce)
	return hash

}
