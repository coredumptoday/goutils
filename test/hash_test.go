package test

import (
	"fmt"
	"testing"

	xhash "github.com/coredumptoday/goutils/hash"
)

func TestMD5(t *testing.T) {
	h := xhash.NewSHA1()
	h.WriteString("abc")
	h.WriteString("def")
	h.WriteString("hig")
	h.WriteString("klm")

	if h.Err() != nil {
		fmt.Println(h.Err())
	}

	fmt.Println(string(h.Sum().ToHex()))
	fmt.Println(string(h.Sum().ToHex().ToUpper()))

	fmt.Println(string(h.Sum().ToHex().MD5().ToHex()))
	fmt.Println(string(h.Sum().ToHex().MD5().ToHex().ToUpper()))
}
