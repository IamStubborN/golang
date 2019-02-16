package main

import (
	"crypto/sha256"
	"encoding/base64"
)

func encryptString(str string) string {
	h := sha256.Sum256([]byte(str))
	return base64.StdEncoding.EncodeToString(h[:])
}
