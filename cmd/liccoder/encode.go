package main

import (
	"fmt"

	"io/ioutil"

	"exobyte.org/pulse/files"
	"exobyte.org/pulse/license"
)

func encode(infile, outfile, privateKeyFile string) error {
	bytes, err := files.ReadBytes(infile)

	if err != nil {
		return err
	}

	privateKey, err := files.ReadBytes(privateKeyFile)

	if err != nil {
		return err
	}

	encoded, err := license.Encode(bytes, privateKey)

	if err != nil {
		return err
	}

	if outfile == "" {
		fmt.Println(encoded)

		return nil
	}

	return ioutil.WriteFile(output, []byte(encoded), 0644)
}
