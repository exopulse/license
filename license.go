// Package license contains support for writing and reading license files. License files are guarded with private/public
// key pair.
//
// Any structure could be used as in-memory license details holder.
//
// Create private/public keys like this:
//   openssl genrsa -out private.pem 2048
//   openssl rsa -in private.pem -outform PEM -pubout -out public.pem
package license

import (
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"

	"crypto/x509"

	"crypto/rsa"

	"crypto"
	"crypto/rand"

	"strings"

	"github.com/pkg/errors"
)

var (
	beginLicense   = "-----BEGIN LICENSE-----"
	endLicense     = "-----END LICENSE-----"
	beginSignature = "-----BEGIN SIGNATURE-----"
	endSignature   = "-----END SIGNATURE-----"
)

// Encode encodes license using private key.
func Encode(license interface{}, privKey []byte) (string, error) {
	bytes, err := json.Marshal(license)

	if err != nil {
		return "", err
	}

	block, _ := pem.Decode(privKey)

	if block == nil {
		return "", errors.New("key not found")
	}

	if block.Type != "RSA PRIVATE KEY" {
		return "", fmt.Errorf("unsupported key type %q", block.Type)
	}

	rsaPrivateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)

	if err != nil {
		return "", err
	}

	signature, err := rsa.SignPKCS1v15(rand.Reader, rsaPrivateKey, crypto.Hash(0), bytes)

	if err != nil {
		return "", err
	}

	encoder := base64.StdEncoding

	return pack(encoder.EncodeToString(bytes), encoder.EncodeToString(signature))
}

func pack(license, signature string) (string, error) {
	var sb strings.Builder

	if err := packBlock(license, beginLicense, endLicense, &sb); err != nil {
		return "", err
	}

	if err := packBlock(signature, beginSignature, endSignature, &sb); err != nil {
		return "", err
	}

	return sb.String(), nil
}

func packBlock(text, header, footer string, sb *strings.Builder) error {
	if _, err := sb.WriteString(header); err != nil {
		return err
	}

	if err := sb.WriteByte('\n'); err != nil {
		return err
	}

	if err := split(text, sb); err != nil {
		return err
	}

	if _, err := sb.WriteString(footer); err != nil {
		return err
	}

	if err := sb.WriteByte('\n'); err != nil {
		return err
	}

	return nil
}

func split(s string, sb *strings.Builder) error {
	i := 0
	size := 76
	ln := len(s)

	for i < ln {
		c := i + size

		if c >= ln {
			c = ln
		}

		if _, err := sb.WriteString(s[i:c]); err != nil {
			return err
		}

		if err := sb.WriteByte('\n'); err != nil {
			return err
		}

		i = c
	}

	return nil
}

// Decode decodes encoded string. Decoding is performed using public key.
func Decode(encoded string, publicKey []byte, license interface{}) error {
	block, _ := pem.Decode(publicKey)

	if block == nil {
		return errors.New("key not found")
	}

	if block.Type != "PUBLIC KEY" {
		return fmt.Errorf("unsupported key type %q", block.Type)
	}

	rsaPublicKey, err := x509.ParsePKIXPublicKey(block.Bytes)

	if err != nil {
		return err
	}

	bytes, err := extract(encoded, beginLicense, endLicense)

	if err != nil {
		return err
	}

	signature, err := extract(encoded, beginSignature, endSignature)

	if err != nil {
		return err
	}

	if err := rsa.VerifyPKCS1v15(rsaPublicKey.(*rsa.PublicKey), crypto.Hash(0), bytes, signature); err != nil {
		return err
	}

	if err := json.Unmarshal(bytes, license); err != nil {
		return err
	}

	return nil
}

func extract(text, header, footer string) ([]byte, error) {
	headerAt := strings.Index(text, header)
	footerAt := strings.Index(text, footer)

	if headerAt < 0 || footerAt < 0 && headerAt < footerAt {
		return nil, errors.New("invalid format")
	}

	s := text[headerAt+len(header) : footerAt]

	return base64.StdEncoding.DecodeString(s)
}
