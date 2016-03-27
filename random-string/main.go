package main

import (
	"crypto/rand"
	"encoding/base32"
	"strings"
)

func main() {
	b := make([]byte, 32)
	rand.Read(b)

	// println(base32.URLEncoding.EncodeToString(b))
	println(strings.TrimRight(base32.StdEncoding.EncodeToString(b), "="))
}
