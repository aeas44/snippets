package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func toHash(password string) string {
	converted := sha256.Sum256([]byte(password))
	return hex.EncodeToString(converted[:])
}

func main() {
	str := "hello"
	fmt.Printf("SHA-256 : %x\n", sha256.Sum256([]byte(str)))
	fmt.Println(toHash(str))
}
