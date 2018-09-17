package license

import (
	"reflect"
	"testing"
	"time"
)

var privKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAyxKPRKyyIKgoi/XNdqXHLndMRkWkESZlF6xJ8P/AY+vXe359
BkQFGIjZ9vLOMsOBgJz/yw+xedLq+nqBT6mf0PIFW05Z4q1mdBEUPjqai/nzDD9s
1Rw89i9aZp9t+yI/KAYBH+ukp/m4dTS00wYcrTloumEeOe0wbrW9iR/wLB2kzRvh
aXwsxyQ6O+Y4oo7BRP/KaRyX5qZXoBhhmg4Mp0Q+CTn+Umf7o4a9vmZo+KXVAi4g
JRL+Rk3IV9tJck1hwGy2fXa7wXgorG2KAUaybj62DOcxQPUN38hW8SU/kcrnwDLL
/ahTn0IJ50NQynexlUdNY7lWImctisIs7o4rmwIDAQABAoIBACq+vl+DRn9vRhGn
7NbTsGiMfgawtAdUIKVqA6px+ypNQ29gQarm+HegntjHZTX0RzcuVP2GUSJGjmWp
Eb0WBGjz+MKubGc1fsJhsfQjW9KpC0sBssPtDLv5XbsKRhk9rTOch/ITJPairvwk
2wcsWy7vTacluDoTPkWfhcTubxypxMC7ZZ63cNXGsXQyiZtmDuuPvGqL3zTy22ok
qySL0NA+CuHzBb5/7qebDd4/UZaPRaCzBYYmAHNpv+T5YCbaDvAsldybtDvVtMrD
6MzRd4VBKv8tXIrs7qdd73uoV6uZ25VaooZKiZtOOetBFbThAoE8b2vm8bixJs9h
6MWTwjECgYEA5Muss2xtsE3mX5+FkMT1AObxW6ohjVANTkzL05N0Fkvt+atYt4nq
+HXyYiP82mQ9H/rk98m6TiBDliYIu8ONOF3WkTc/MD80FlRDE5+GtuifzsromjS6
FQJp1qrOeknZW9gHQVDfj9PghDrPpcDiOYo6BkGSHjoH1GckKPXgyAkCgYEA4zfl
ytV9sNa0p4rzyGXkiDN5MasQKhCmIZk40qpyJmHj82tV/W825kYsbWToYbEG9ipn
KTBSDUyd9pydlpFCSla8tRKGI4C+zIqEjGChrAIqYGA3yg6PThMxgIYIBnlx7GuX
Apgz75yLiDUEGfl7M/hKmRhsHFHMt2tr1VPXF4MCgYEAwgx8zTCJV566ZR/HCKID
qmA1FKmXC2N55Y780NWe/8ombDsHY6N1xoVCrjoXmqKptRX2XLKbcOesawyXG6kf
fU8ZTNCupxzz8HxG573FZByobZtq7F8L9AW3/nmf5df8ogj3Hk6ZTkNFiODYfJy7
QlIaJHQ4mapeOafA43ymyQECgYEAzJl35SZKsOMU90zq6ce+elqHRg84Vfh58AUt
opzyftdk0LsVSklL0RzdQoA7tEQY9a8HB5LIP8t+7iLm0D7qx0/FSvNsMCntJPeS
FWmqCt5EVFYvxlVH+1DTIV0Peuz+hZn399Ef6yI2jqbWk6EPrKNmYOcZ6oJJ6dvj
1MtTAosCgYBu830Z9Wxwvv6m3EGrD7VjshoVe6Er4beOxTeAvf26Xj2XR77aX+2G
XZfwLWTr30wZdauZnYwwFxyNAC2R9RTKnODrsNJRW/NpE0YzzfLWfd+Nx14VxspJ
GDrFAeoVdaXX2go4Lj7psfxbxqFe27i7ldTQibIGV3qOJdC+Tc169g==
-----END RSA PRIVATE KEY-----
`

var publicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAyxKPRKyyIKgoi/XNdqXH
LndMRkWkESZlF6xJ8P/AY+vXe359BkQFGIjZ9vLOMsOBgJz/yw+xedLq+nqBT6mf
0PIFW05Z4q1mdBEUPjqai/nzDD9s1Rw89i9aZp9t+yI/KAYBH+ukp/m4dTS00wYc
rTloumEeOe0wbrW9iR/wLB2kzRvhaXwsxyQ6O+Y4oo7BRP/KaRyX5qZXoBhhmg4M
p0Q+CTn+Umf7o4a9vmZo+KXVAi4gJRL+Rk3IV9tJck1hwGy2fXa7wXgorG2KAUay
bj62DOcxQPUN38hW8SU/kcrnwDLL/ahTn0IJ50NQynexlUdNY7lWImctisIs7o4r
mwIDAQAB
-----END PUBLIC KEY-----
`

var encodedLicense = `-----BEGIN LICENSE-----
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
`

type myLicense struct {
	ID       string    `json:"id"`
	IssuedTo string    `json:"issuedTo"`
	ValidTo  time.Time `json:"validTo"`
}

func TestEncode(t *testing.T) {
	var myLic = sampleLicense()

	encoded, err := Encode(myLic, []byte(privKey))

	if err != nil {
		t.Fatal(err)
	}

	if encoded != encodedLicense {
		t.Error("encoded licenses differ")
	}
}

func TestDecode(t *testing.T) {
	myLic := myLicense{}

	if err := Decode(encodedLicense, []byte(publicKey), &myLic); err != nil {
		t.Fatal(err)
	}

	sampleLic := sampleLicense()

	if !reflect.DeepEqual(myLic, sampleLic) {
		t.Error("decoded licenses differ")
	}
}

func sampleLicense() myLicense {
	var validTo, err = time.Parse(time.RFC3339, "2018-09-17T09:27:00+02:00")

	if err != nil {
		panic(err)
	}

	return myLicense{ID: "1234", IssuedTo: "some company", ValidTo: validTo}
}
