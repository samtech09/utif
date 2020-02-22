package utif

import (
	"fmt"
	"testing"
)

func Test_RandStringAlpha(t *testing.T) {
	str := RandStringAlpha(20)
	fmt.Printf("Alpha: %s\n", str)
}

func Test_RandStringNumeric(t *testing.T) {
	str := RandStringNumeric(20)
	fmt.Printf("Num: %s\n", str)
}

func Test_RandStringAlphaNumeric(t *testing.T) {
	str := RandStringAlphaNumeric(20)
	fmt.Printf("AlphaNum: %s\n", str)
}
