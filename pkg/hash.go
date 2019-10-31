package pkg

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
	"io"
	"os"

	"github.com/cxmcc/tiger"
	"github.com/jzelinskie/whirlpool"
	"github.com/raja/argon2pw"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
)

var AlgoFlag string

func callHashFactory(input string, hashFactory func() hash.Hash) string {
	hash := hashFactory()
	io.WriteString(hash, input)
	bytes := fmt.Sprintf("%x", hash.Sum(nil))
	return bytes
}
func HashPassword(password string) (string, error) {
	switch AlgoFlag {
	case "argon2":
		h, err := argon2pw.GenerateSaltedHash(password)
		if err != nil {
			panic(err)
		}
		s := fmt.Sprintf("%s", h)
		return s, nil
	case "ntlm":
		bytes, _ := nthash(password)
		s := fmt.Sprintf("%s", bytes)
		return s, nil
	case "lanman":
		bytes := lanMan(password)
		s := fmt.Sprintf("%x", bytes)
		return s, nil
	case "bcrypt":
		bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
		return string(bytes), err
	case "md4":
		bytes := callHashFactory(password, md4.New)
		return string(bytes), nil
	case "md5":
		bytes := callHashFactory(password, md5.New)
		return string(bytes), nil
	case "ripemd160":
		bytes := callHashFactory(password, ripemd160.New)
		return string(bytes), nil
	case "sha1":
		bytes := callHashFactory(password, sha1.New)
		return string(bytes), nil
	case "sha2-224":
		bytes := callHashFactory(password, sha256.New224)
		return string(bytes), nil
	case "sha2-256":
		bytes := callHashFactory(password, sha256.New)
		return string(bytes), nil
	case "sha2-512":
		bytes := callHashFactory(password, sha512.New)
		return string(bytes), nil
	case "sha3-224":
		bytes := callHashFactory(password, sha3.New224)
		return string(bytes), nil
	case "sha3-256":
		bytes := callHashFactory(password, sha3.New256)
		return string(bytes), nil
	case "sha3-512":
		bytes := callHashFactory(password, sha3.New512)
		return string(bytes), nil
	case "tiger":
		h := tiger.New()
		io.WriteString(h, password)
		bytes := fmt.Sprintf("%x", h.Sum(nil))
		return string(bytes), nil
	case "uuid1":
		u1 := uuid.Must(uuid.NewV1(), nil)
		u2 := fmt.Sprintf("%s", u1)
		return string(u2), nil
	case "uuid4":
		u1 := uuid.Must(uuid.NewV4(), nil)
		u2 := fmt.Sprintf("%s", u1)
		return string(u2), nil
	case "whirlpool":
		w := whirlpool.New()
		text := []byte(password)
		w.Write(text)
		s := fmt.Sprintf("%x", w.Sum(nil))
		return s, nil

	default:
		fmt.Printf("You provided an invalid flag for the algorithm.\n")
		os.Exit(1)
		return "", nil

	}

}
