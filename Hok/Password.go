package Hok

import (
	"crypto/md5"
	"fmt"
)

func Md5PassWord(Raw string) string {
	data := []byte(Raw)
	Has := md5.Sum(data)
	return fmt.Sprintf("%x", Has)
}
