package cryptil

import (
	"fmt"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
)

//MD5加密
func MD5(s string) (string) {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func SHA1(s string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(s)))
}

func SHA256(s string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(s)))
}