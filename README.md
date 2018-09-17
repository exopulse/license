# exopulse license package
Golang support for writing and reading license files.

[![CircleCI](https://circleci.com/gh/exopulse/license.svg?style=svg)](https://circleci.com/gh/exopulse/license)
[![Build Status](https://travis-ci.org/exopulse/license.svg?branch=master)](https://travis-ci.org/exopulse/license)
[![GitHub license](https://img.shields.io/github/license/exopulse/license.svg)](https://github.com/exopulse/license/blob/master/LICENSE)

# Overview

This module contains support for writing and reading license files. License files are guarded with private/public key pair.

Any structure could be used as in-memory license details holder.

Encoded license looks like this:

    -----BEGIN LICENSE-----
    eyJpZCI6IjEyMzQiLCJpc3N1ZWRUbyI6InNvbWUgY29tcGFueSIsInZhbGlkVG8iOiIyMDE4LTA5
    LTE3VDA5OjI3OjAwKzAyOjAwIn0=
    -----END LICENSE-----
    -----BEGIN SIGNATURE-----
    FqkckB4ugQcWM+BFNbNZxFLohw2a/gk8kWe1Vq2p4keE6d/TvMDUjYIc6PJqSsHH0PdOn/ya7N3y
    8FP7s4GtR0epGzxwfc7U8VjrzGSqBXWm2PhpkMVDUsjA1lS++EKZMzdBF7JpVSYHMTTpZgxQ2JPX
    Pqt5cAF3YRb8JHQdhQdTux8M0rN8kJ5YiDiYtdPc5KENVCRqNtKeD1ubsSf0Ya9yCoXJBHDvNM5Y
    UsNnK34dmtjoGHfQtfz81BZmeNzxliu3jblxO3lNqPj/XKtekEhmfiFn0/MSwc2SwKk4bmMDEQaV
    TMHjefTSrrkD0byxQlknxYWkdHjT8zXXpEDT/w==
    -----END SIGNATURE-----

Create private/public keys like this:

    $ openssl genrsa -out private.pem 2048
    $ openssl rsa -in private.pem -outform PEM -pubout -out public.pem

# Features

## Encode/decode methods

Use this methods in your application to encode or decode license files, using private and public keys.

## Liccoder command line tool

Embedded within this project is a command line tool for encoding/decoding licenses.
See examples for more info.

Use go install to install this tool to $GOBIN directory.

    $ go install ./cmd/liccoder

Use _encode_ command to encode input file.

    $ liccoder encode --in license.txt --key-file private.pem --out encoded.txt

Use _decode_ command to decode input file.

    $ liccoder decode --in encoded.txt --key-file public.pem --out decoded.txt

# Using license package

## Installing package

Use go get to install the latest version of the library.

    $ go get github.com/exopulse/license
 
Include license in your application.
```go
import "github.com/exopulse/license"
```

## Functions
```go
// Encode encodes license using private key.
func Encode(license interface{}, privKey []byte) (string, error)

// Decode decodes encoded string. Decoding is performed using public key.
func Decode(encoded string, publicKey []byte, license interface{}) error
```

### Examples

```go
type myLicense struct {
	ID       string    `json:"id"`
	IssuedTo string    `json:"issuedTo"`
	ValidTo  time.Time `json:"validTo"`
}

func encode() {
	var validTo, _ = time.Parse(time.RFC3339, "2018-09-17T09:27:00+02:00")
    var myLic = myLicense{ID: "1234", IssuedTo: "some company", ValidTo: validTo}

    encoded, err := Encode(myLic, []byte(privKey))
}

func decode() {
	myLic := myLicense{}

	if err := Decode(encodedLicense, []byte(publicKey), &myLic); err != nil {
		panic(err)
	}
}
```

# About the project

## Contributors

* [exopulse](https://github.com/exopulse)

## License

License package is released under the MIT license. See
[LICENSE](https://github.com/exopulse/license/blob/master/LICENSE)
