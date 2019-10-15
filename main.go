package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"hash"
	"io"
	"os"
	"strings"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
)

var algoFlag string

func callHashFactory(input string, hashFactory func() hash.Hash) string {
	hash := hashFactory()
	io.WriteString(hash, input)
	bytes := fmt.Sprintf("%x", hash.Sum(nil))
	return bytes
}

func hashPassword(password string) (string, error) {
	switch algoFlag {
	case "NTLM":
		bytes, _ := nthash(password)
		s := fmt.Sprintf("%s", bytes)
		return s, nil
	case "LANMAN":
		bytes := lanMan(password)
		s := fmt.Sprintf("%x", bytes)
		return s, nil
	case "BCRYPT":
		bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
		return string(bytes), err
	case "MD4":
		bytes := callHashFactory(password, md4.New)
		return string(bytes), nil
	case "MD5":
		bytes := callHashFactory(password, md5.New)
		return string(bytes), nil
	case "RIPEMD160":
		bytes := callHashFactory(password, ripemd160.New)
		return string(bytes), nil
	case "SHA1":
		bytes := callHashFactory(password, sha1.New)
		return string(bytes), nil
	case "SHA2-224":
		bytes := callHashFactory(password, sha256.New224)
		return string(bytes), nil
	case "SHA2-256":
		bytes := callHashFactory(password, sha256.New)
		return string(bytes), nil
	case "SHA2-512":
		bytes := callHashFactory(password, sha512.New)
		return string(bytes), nil
	case "SHA3-224":
		bytes := callHashFactory(password, sha3.New224)
		return string(bytes), nil
	case "SHA3-256":
		bytes := callHashFactory(password, sha3.New256)
		return string(bytes), nil
	case "SHA3-512":
		bytes := callHashFactory(password, sha3.New512)
		return string(bytes), nil
	case "UUID1":
		u1 := uuid.Must(uuid.NewV1())
		u2 := fmt.Sprintf("%s", u1)
		return string(u2), nil
	case "UUID4":
		u1 := uuid.Must(uuid.NewV4())
		u2 := fmt.Sprintf("%s", u1)
		return string(u2), nil

	default:
		fmt.Printf("You provided an invalid flag for the algorithm.\n")
		os.Exit(1)
		return "", nil

	}

}

func main() {

	algo := flag.String("a", "", "Provide algorithm lanman ntlm md4 md5 bcrypt ripemd160 sha1 sha2-224 sha2-256 sha2-512 sha3-224 sha3-256 sha3-512 uuid1 uuid4")
	flag.Parse()

	value := *algo
	algoFlag = strings.ToUpper(value)

	if flag.Arg(0) == "" {
		switch algoFlag {
		case "UUID1":
			u1 := uuid.Must(uuid.NewV1())
			fmt.Printf("%s\n", u1)
			return
		case "UUID4":
			u1 := uuid.Must(uuid.NewV4())
			fmt.Printf("%s\n", u1)
			return
		}
		sc := bufio.NewScanner(os.Stdin)

		for sc.Scan() {
			password := sc.Text()
			hash, _ := hashPassword(password)
			fmt.Println(hash)
		}

		if err := sc.Err(); err != nil {
			fmt.Println("failed to read input.")
		}
	} else {

		if flag.Arg(0) != "" {
			password := flag.Arg(0)
			hash, _ := hashPassword(password)
			fmt.Println(hash)
		}
	}

}
