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

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/md4"
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
	case "BCRYPT":
		bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
		return string(bytes), err
	case "MD4":
		bytes := callHashFactory(password, md4.New)
		return string(bytes), nil
	case "MD5":
		bytes := callHashFactory(password, md5.New)
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
	default:
		return "", nil

	}

}

func main() {

	algo := flag.String("a", "", "Provide algorithm md4 md5 bcrypt sha1 sha2-224 sha2-256 sha2-512 sha3-224 sha3-256 sha3-512")
	flag.Parse()

	value := *algo
	value = strings.ToUpper(value)

	switch value {
	case "MD4":
		algoFlag = "MD4"
	case "MD5":
		algoFlag = "MD5"
	case "BCRYPT":
		algoFlag = "BCRYPT"
	case "SHA1":
		algoFlag = "SHA1"
	case "SHA2-224":
		algoFlag = "SHA2-224"
	case "SHA2-256":
		algoFlag = "SHA2-256"
	case "SHA2-512":
		algoFlag = "SHA2-512"
	case "SHA3-224":
		algoFlag = "SHA3-224"
	case "SHA3-256":
		algoFlag = "SHA3-256"
	case "SHA3-512":
		algoFlag = "SHA3-512"

	default:
		fmt.Println("No compatible algorithm.")
		return

	}

	if flag.Arg(0) == "" {
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
		password := flag.Arg(0)
		hash, _ := hashPassword(password)
		fmt.Println(hash)
	}

	//password := flag.Arg(0)
	//hash, _ := hashPassword(password)
	//fmt.Println(hash)

}