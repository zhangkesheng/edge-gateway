package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func HmacSha256Sign(key, data string) string {
	secret := []byte(key)
	msg := []byte(data)

	h := hmac.New(sha256.New, secret)
	h.Write(msg)

	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
