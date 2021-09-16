package md5

import (
	"crypto/md5"
	"encoding/hex"
)

func HashMd5(hashString string) (hash string) {
	hashMd5 := md5.New()
	hashMd5.Write([]byte(hashString))

	return hex.EncodeToString(hashMd5.Sum(nil))
}
