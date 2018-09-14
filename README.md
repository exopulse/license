# exopulse license package
Golang support for writing and reading license files.

[![CircleCI](https://circleci.com/gh/exopulse/license.svg?style=svg)](https://circleci.com/gh/exopulse/license)
[![Build Status](https://travis-ci.org/exopulse/license.svg?branch=master)](https://travis-ci.org/exopulse/license)
[![GitHub license](https://img.shields.io/github/license/exopulse/license.svg)](https://github.com/exopulse/license/blob/master/LICENSE)

# Overview

This contains support for writing and reading license files. License files are guarded with private/public key pair.

Any structure could be used as in-memory license details holder.

Create private/public keys like this:

    $ openssl genrsa -out private.pem 2048
    $ openssl rsa -in private.pem -outform PEM -pubout -out public.pem

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

# About the project

## Contributors

* [exopulse](https://github.com/exopulse)

## License

License package is released under the MIT license. See
[LICENSE](https://github.com/exopulse/license/blob/master/LICENSE)
