package utif

import (
	"log"
	"testing"
	"time"
)

type person struct {
	Name string
	Age  int
	Dob  time.Time
}

func Test_MD5Struct(t *testing.T) {
	dob, _ := time.Parse(time.RFC3339, "1996-02-04")
	p := person{"Santosh", 21, dob}
	log.Printf("%#v : ", p)
	h1 := MD5Struct(p, "a1c4b8g7")
	log.Printf("hash-1: %s\n", h1)

	p2 := person{"Santosh", 21, dob}
	log.Printf("%#v : ", p2)
	h2 := MD5Struct(p2, "a1c4b8g7")
	log.Printf("hash-2: %s\n", h2)

	if h1 != h2 {
		t.Error("Both hash not same")
	}
}
