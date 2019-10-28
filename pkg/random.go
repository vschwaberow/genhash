package pkg

import "crypto/rand"

func GenerateRandomBytes(rsize int) (rndbytes []byte, nrsize int) {
	key := make([]byte, rsize)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}
	nrsize = len(key)
	return key, nrsize
}
