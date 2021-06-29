package xcrypto

import (
	"fmt"
	"testing"
)

var tdesKey = []byte("abcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcd")
var tdesIv = []byte("defgdefg")
var tdesData = []byte("3DES（又叫Triple DES）是三重数据加密算法（TDEA，Triple Data Encryption Algorithm）块密码的通称。它相当于是对每个数据块应用三次DES加密算法。密钥长度是128位，192位(bit)，如果密码位数少于等于64位，加密结果与DES相同。原版DES容易被破解，新的3DES出现，增加了加密安全性,避免被暴力破解。它同样是对称性加密，同样涉及到加密编码方式，及填充方式。包括3DES-ECB,3DES-CBC,3DES-CTR,3DES-OFB,3DES-CFB")

func Test3DesECB(t *testing.T) {
	de := NewTripleDesEncipher(tdesKey)
	endata, err := de.ECB().SetPadding(PKCS5).Do(tdesData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(endata.HexEncode().ToString())
	fmt.Println(endata.StdBase64Encode().ToString())
	fmt.Println(endata.URLBase64Encode().ToString())

	dd := NewTripleDesDecipher(tdesKey)
	origin, err := dd.ECB().SetPadding(PKCS5).Do(endata)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(origin.ToString())
}

func Test3DesCBC(t *testing.T) {
	de := NewTripleDesEncipher(tdesKey)
	endata, err := de.CBC(tdesIv).SetPadding(PKCS5).Do(tdesData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(endata.HexEncode().ToString())
	fmt.Println(endata.StdBase64Encode().ToString())
	fmt.Println(endata.URLBase64Encode().ToString())

	dd := NewTripleDesDecipher(tdesKey)
	origin, err := dd.CBC(tdesIv).SetPadding(PKCS5).Do(endata)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(origin.ToString())
}

func Test3DesCTR(t *testing.T) {
	de := NewTripleDesEncipher(tdesKey)
	endata, err := de.CTR(tdesIv).Do(tdesData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(endata.HexEncode().ToString())
	fmt.Println(endata.StdBase64Encode().ToString())
	fmt.Println(endata.URLBase64Encode().ToString())

	dd := NewTripleDesDecipher(tdesKey)
	origin, err := dd.CTR(tdesIv).Do(endata)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(origin.ToString())
}

func Test3DesOFB(t *testing.T) {
	de := NewTripleDesEncipher(tdesKey)
	endata, err := de.OFB(tdesIv).Do(tdesData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(endata.HexEncode().ToString())
	fmt.Println(endata.StdBase64Encode().ToString())
	fmt.Println(endata.URLBase64Encode().ToString())

	dd := NewTripleDesDecipher(tdesKey)
	origin, err := dd.OFB(tdesIv).Do(endata)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(origin.ToString())
}

func Test3DesCFB(t *testing.T) {
	de := NewTripleDesEncipher(tdesKey)
	endata, err := de.CFB(tdesIv).Do(tdesData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(endata.HexEncode().ToString())
	fmt.Println(endata.StdBase64Encode().ToString())
	fmt.Println(endata.URLBase64Encode().ToString())

	dd := NewTripleDesDecipher(tdesKey)
	origin, err := dd.CFB(tdesIv).Do(endata)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(origin.ToString())
}
