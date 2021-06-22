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

	fmt.Println(string(h.Sum().ToHex()))           //计算sha1，转换hex编码
	fmt.Println(string(h.Sum().ToHex().ToUpper())) //计算sha1，转换hex编码，大写hex编码

	fmt.Println(h.Sum().ToHex().ToString())           //计算sha1，转换hex编码，转换string类型
	fmt.Println(h.Sum().ToHex().ToUpper().ToString()) //计算sha1，转换hex编码，大写hex编码，转换string类型

	fmt.Println(string(h.Sum().ToHex().MD5().ToHex())) //计算sha1，转换hex编码，对hex编码求md5值，转换MD5值的hex
	fmt.Println(string(h.Sum().ToHex().MD5().ToHex().ToUpper()))

	fmt.Println("=============")
	fmt.Println(h.Sum().ToHex().ToString())
	fmt.Println(h.Sum().ToHex().MD5().ToHex().ToString())
	fmt.Println(h.Sum().ToHex().SHA1().ToHex().ToString())
	fmt.Println(h.Sum().ToHex().SHA256().ToHex().ToString())
	fmt.Println(h.Sum().ToHex().SHA256_224().ToHex().ToString())
	fmt.Println(h.Sum().ToHex().SHA512_224().ToHex().ToString())
	fmt.Println(h.Sum().ToHex().SHA512_256().ToHex().ToString())
	fmt.Println(h.Sum().ToHex().SHA384().ToHex().ToString())
	fmt.Println(h.Sum().ToHex().SHA512().ToHex().ToString())
}
