package bytes

import (
	"crypto"
	"fmt"
	"testing"
)

//一次性输入进行hash计算
func TestBytes(t *testing.T) {
	ob := Bytes("abcdefghijklmn")
	fmt.Println(ob.MD5().HexEncode().ToString())                    //计算md5
	fmt.Println(ob.SHA1().HexEncode().MD5().HexEncode().ToString()) //计算md5(sha1(str))

	fmt.Println(ob.HMAC(crypto.SHA1, []byte("123")).HexEncode().ToUpper().ToString())
}
