package util

import (
	"crypto/md5"
	"fmt"
	"io"
)

// GenerateID returns md5(张三20160812)
func GenerateID(name, date string) string {
	h := md5.New()
	io.WriteString(h, name)
	io.WriteString(h, date)
	return fmt.Sprintf("%x", h.Sum(nil))
}
