package hash

import (
	"crypto/md5"
	"fmt"
)

func GetSmallHash(s string) string {
	hash := md5.Sum([]byte(s))

	return fmt.Sprintf("%x", hash)
}
