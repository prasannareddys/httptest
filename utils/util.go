package utils

import (
	"crypto/md5"
	"io"
)

func GetHash(s string) string {
	nh := md5.New()
	_, _ = io.WriteString(nh, s)
	return string(nh.Sum(nil))
}
