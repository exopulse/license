package main

import (
	"fmt"
	"io/ioutil"

	"exobyte.org/pulse/files"
	"exobyte.org/pulse/license"
)

func decode(infile, outfile, publicKeyFile string) error {
	bytes, err := files.ReadBytes(infile)

	if err != nil {
		return err
	}

	publicKey, err := files.ReadBytes(publicKeyFile)

	if err != nil {
		return err
	}

	var decoded []byte

	if err := license.Decode(string(bytes), publicKey, &decoded); err != nil {
		return err
	}

	if outfile == "" {
		fmt.Println(string(decoded))

		return nil
	}

	return ioutil.WriteFile(output, decoded, 0644)
}
