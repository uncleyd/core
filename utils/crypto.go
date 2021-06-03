package utils

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
)

func MD5(value string) string {
	data := []byte(value)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}

func EncodeBase64(value string) string {
	data := []byte(value)
	return base64.StdEncoding.EncodeToString(data)
}
