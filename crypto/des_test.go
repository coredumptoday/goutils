package crypto

import (
	"fmt"
	"testing"

	"github.com/coredumptoday/goutils/bytes"
)

var desKey = []byte("abcdabcdabcdabcd")
var desIv = []byte("defgdefg")
var desData = []byte("DES是对称性加密里面常见一种，全称为Data Encryption Standard，即数据加密标准，是一种使用密钥加密的块算法。密钥长度是64位(bit)，超过位数密钥被忽略。所谓对称性加密，加密和解密密钥相同。对称性加密一般会按照固定长度，把待加密字符串分成块。不足一整块或者刚好最后有特殊填充字符。往往跨语言做DES加密解密，经常会出现问题。往往是填充方式不对、或者编码不一致、或者选择加密解密模式(ECB,CBC,CTR,OFB,CFB,NCFB,NOFB)没有对应上造成。常见的填充模式有： 'pkcs5','pkcs7','iso10126','ansix923','zero' 类型，包括DES-ECB,DES-CBC,DES-CTR,DES-OFB,DES-CFB")

func TestDesECB(t *testing.T) {
	de := NewDesEncipher(desKey)
	endata, err := de.ECB().SetPadding(PKCS5).Do(desData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bytes.Bytes(endata).HexEncode().ToString())

	dd := NewDesDecipher(desKey)
	origin, err := dd.ECB().SetPadding(PKCS5).Do(endata)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bytes.Bytes(origin).ToString())
}

func TestDesCBC(t *testing.T) {
	de := NewDesEncipher(desKey)
	endata, err := de.CBC(desIv).SetPadding(PKCS5).Do(desData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bytes.Bytes(endata).HexEncode().ToString())

	dd := NewDesDecipher(desKey)
	origin, err := dd.CBC(desIv).SetPadding(PKCS5).Do(endata)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bytes.Bytes(origin).ToString())
}

func TestDesCTR(t *testing.T) {
	de := NewDesEncipher(desKey)
	endata, err := de.CTR(desIv).Do(desData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bytes.Bytes(endata).HexEncode().ToString())

	dd := NewDesDecipher(desKey)
	origin, err := dd.CTR(desIv).Do(endata)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bytes.Bytes(origin).ToString())
}

func TestDesOFB(t *testing.T) {
	de := NewDesEncipher(desKey)
	endata, err := de.OFB(desIv).Do(desData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bytes.Bytes(endata).HexEncode().ToString())

	dd := NewDesDecipher(desKey)
	origin, err := dd.OFB(desIv).Do(endata)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bytes.Bytes(origin).ToString())
}

func TestDesCFB(t *testing.T) {
	de := NewDesEncipher(desKey)
	endata, err := de.CFB(desIv).Do(desData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bytes.Bytes(endata).HexEncode().ToString())

	dd := NewDesDecipher(desKey)
	origin, err := dd.CFB(desIv).Do(endata)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bytes.Bytes(origin).ToString())
}
