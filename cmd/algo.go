package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/vschwaberow/genhash/pkg"
)

func randomGenerate(cmd *cobra.Command, args []string) {
	if Algo == true {
		fmt.Println("Sorry. The stdin flag has not action in combination with random bytes generation.")
	} else {
		if len(args) > 0 {
			bsize, _ := strconv.Atoi(args[0])
			rndbytes, rndsize := pkg.GenerateRandomBytes(bsize)
			fmt.Println("Here are random values of", rndsize, "bytes length")
			fmt.Println(rndbytes)
			fmt.Printf("%x\n", rndbytes)
		}
	}
}

func output(password string) {
	hash, _ := pkg.HashPassword(password)
	fmt.Printf("%s\n", hash)
}

func hashGenerate(cmd *cobra.Command, args []string) {
	if Algo == false {
		if len(args) > 0 {
			pkg.AlgoFlag = cmd.Name()
			output(args[0])
		} else {
			fmt.Println("You need to provide the password after the algo command.")
		}
	} else {
		sc := bufio.NewScanner(os.Stdin)

		for sc.Scan() {
			pkg.AlgoFlag = cmd.Name()
			password := sc.Text()
			hash, _ := pkg.HashPassword(password)
			fmt.Println(hash)
		}
		if err := sc.Err(); err != nil {
			fmt.Println("Failed to read input. ", err)
		}

	}
}

var ntlmCmd = &cobra.Command{
	Use:   "ntlm",
	Short: "Generate NTLM hash",
	Run:   hashGenerate,
}
var lanmanCmd = &cobra.Command{
	Use:   "lanman",
	Short: "Generate LanManager hash",
	Run:   hashGenerate,
}
var bcryptCmd = &cobra.Command{
	Use:   "bcrypt",
	Short: "Generate Bcrypt hash",
	Run:   hashGenerate,
}
var md4Cmd = &cobra.Command{
	Use:   "md4",
	Short: "Generate MD4 hash",
	Run:   hashGenerate,
}
var md5Cmd = &cobra.Command{
	Use:   "md5",
	Short: "Generate MD5 hash",
	Run:   hashGenerate,
}
var ripemdCmd = &cobra.Command{
	Use:   "ripemd160",
	Short: "Generate RIPEMD160 hash",
	Run:   hashGenerate,
}
var sha1Cmd = &cobra.Command{
	Use:   "sha1",
	Short: "Generate SHA-1 hash",
	Run:   hashGenerate,
}
var sha2224Cmd = &cobra.Command{
	Use:   "sha2-224",
	Short: "Generate SHA2-224 hash",
	Run:   hashGenerate,
}
var sha2256Cmd = &cobra.Command{
	Use:   "sha2-256",
	Short: "Generate SHA2-256 hash",
	Run:   hashGenerate,
}
var sha2512Cmd = &cobra.Command{
	Use:   "sha2-512",
	Short: "Generate SHA2-512 hash",
	Run:   hashGenerate,
}
var sha3224Cmd = &cobra.Command{
	Use:   "sha3-224",
	Short: "Generate SHA3-224 hash",
	Run:   hashGenerate,
}
var sha3256Cmd = &cobra.Command{
	Use:   "sha3-256",
	Short: "Generate SHA3-256 hash",
	Run:   hashGenerate,
}
var sha3512Cmd = &cobra.Command{
	Use:   "sha3-512",
	Short: "Generate SHA3-512 hash",
	Run:   hashGenerate,
}
var uuid1Cmd = &cobra.Command{
	Use:   "uuid1",
	Short: "Generate UUID-1 random string",
	Run:   hashGenerate,
}
var uuid4Cmd = &cobra.Command{
	Use:   "uuid4",
	Short: "Generate UUID-4 random string",
	Run:   hashGenerate,
}
var randomCmd = &cobra.Command{
	Use:   "rand",
	Short: "Generate random bytes of size provided",
	Run:   randomGenerate,
}
