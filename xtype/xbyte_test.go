package xtype

import (
	"crypto/sha1"
	"fmt"
	"testing"
)

//一次性输入进行hash计算
func TestXBS(t *testing.T) {
	ob := XBS("abcdefghijklmn")
	fmt.Println(ob.MD5().HexEncode().ToString())                    //计算md5
	fmt.Println(ob.SHA1().HexEncode().MD5().HexEncode().ToString()) //计算md5(sha1(str))

	fmt.Println(ob.HMAC(sha1.New, []byte("123")).HexEncode().ToUpper().ToString())
}

func TestEXBS(t *testing.T) {
	ob := XBS("20a56a6e9cb88e6d68365c763fc9fefd9d21b6f9")

	fmt.Println(ob.HexDecode().URLBase64Encode().ToString())
	fmt.Println(ob.HexDecode().URLBase64Encode().URLBase64Decode().HexEncode().ToString())
}
