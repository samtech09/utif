//Package utif stands for Utility Functions, providing different functions to facilitate common tasks
package utif

import (
	"math/rand"
	"time"
)

// Implementations

func init() {
	rand.Seed(time.Now().UnixNano())
}

const numBytes = "0123456789"
const alphaBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func newRandSource() rand.Source {
	return rand.NewSource(time.Now().UnixNano())
}

//RandStringAlpha returns random alphabet string of given length
func RandStringAlpha(lenght int) string {
	return genRandString(lenght, newRandSource(), alphaBytes)
}

//RandStringNumeric returns random numeric string of given length
func RandStringNumeric(lenght int) string {
	return genRandString(lenght, newRandSource(), numBytes)
}

//RandStringAlphaNumeric returns random alpha-numeric string of given length
func RandStringAlphaNumeric(lenght int) string {
	return genRandString(lenght, newRandSource(), numBytes+alphaBytes)
}

//RandString RandStringBytesMaskImprSrc
func genRandString(n int, src rand.Source, letterBytes string) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
