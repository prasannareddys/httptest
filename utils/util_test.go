package utils

import (
	"crypto/md5"
	"io"
	"testing"
)

func TestGetHash(t *testing.T) {
	sh := md5.New()
	io.WriteString(sh, "this is test")
	expected := sh.Sum(nil)
	actual := GetHash("this is test")
	if actual != string(expected) {
		t.Fail()
	}
}
