package utif

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
)

//MD5Str gives MD5 hash of given string and salt
func MD5Str(s, salt string) string {
	h := md5.New()
	h.Write([]byte(s + salt))

	return fmt.Sprintf("%x", h.Sum(nil))
}

//MD5Struct gives MD5 hash of given struct and salt
func MD5Struct(o interface{}, salt string) string {
	h := md5.New()
	h.Write([]byte(fmt.Sprintf("%v%s", o, salt)))

	return fmt.Sprintf("%x", h.Sum(nil))
}

//Sha256Str gives SHA256 hash of given string and salt
func Sha256Str(s, salt string) string {
	h := sha256.New()
	h.Write([]byte(s + salt))

	return fmt.Sprintf("%x", h.Sum(nil))
}

//Sha256Struct gives SHA256 hash of given struct and salt
func Sha256Struct(o interface{}, salt string) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v%s", o, salt)))

	return fmt.Sprintf("%x", h.Sum(nil))
}
