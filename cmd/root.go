package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Algo bool

var rootCmd = &cobra.Command{
	Use:   "genhash",
	Short: "genhash is a generic command line tool for generating hashes of various cryptographic algorithms.",
	Long:  `A very flexible generic command line tool for generating hashes of various crypto algorithms.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Algo, "stdio", "s", false, "Take list from STDIO and generate hashes.")
	//	rootCmd.MarkFlagRequired("algo")
	rootCmd.AddCommand(ntlmCmd)
	rootCmd.AddCommand(lanmanCmd)
	rootCmd.AddCommand(bcryptCmd)
	rootCmd.AddCommand(md4Cmd)
	rootCmd.AddCommand(md5Cmd)
	rootCmd.AddCommand(ripemdCmd)
	rootCmd.AddCommand(sha1Cmd)
	rootCmd.AddCommand(sha2224Cmd)
	rootCmd.AddCommand(sha2256Cmd)
	rootCmd.AddCommand(sha2512Cmd)
	rootCmd.AddCommand(sha3224Cmd)
	rootCmd.AddCommand(sha3256Cmd)
	rootCmd.AddCommand(sha3512Cmd)
	rootCmd.AddCommand(uuid1Cmd)
	rootCmd.AddCommand(uuid4Cmd)
	rootCmd.AddCommand(randomCmd)
	rootCmd.AddCommand(whirlCmd)
	rootCmd.AddCommand(argon2Cmd)
	rootCmd.AddCommand(tigerCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
