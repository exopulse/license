package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	input  string
	output string
	key    string
)

var rootCmd = &cobra.Command{
	Use:     "liccoder",
	Short:   "License coder and decoder.",
	Version: "1.0.0",
}

var encodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "Encode license file",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if err := encode(input, output, key); err != nil {
			fmt.Println(err)

			os.Exit(1)
		}
	},
}

var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Decode license file",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if err := decode(input, output, key); err != nil {
			fmt.Println(err)

			os.Exit(1)
		}
	},
}

func init() {
	encodeCmd.Flags().StringVarP(&input, "in", "i", "", "input file")
	encodeCmd.Flags().StringVarP(&output, "out", "o", "", "output file")
	encodeCmd.Flags().StringVarP(&key, "key-file", "k", "", "private key file")

	decodeCmd.Flags().StringVarP(&input, "in", "i", "", "input file")
	decodeCmd.Flags().StringVarP(&output, "out", "o", "", "output file")
	decodeCmd.Flags().StringVarP(&key, "key-file", "k", "", "public key file")

	if err := encodeCmd.MarkFlagRequired("key-file"); err != nil {
		panic(err)
	}

	if err := decodeCmd.MarkFlagRequired("key-file"); err != nil {
		panic(err)
	}

	rootCmd.AddCommand(encodeCmd, decodeCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
