package xcrypto

import (
	"fmt"
	"testing"
)

var key = []byte("abcdabcdabcdabcd")
var iv = []byte("defgdefgdefgdefg")
var data = []byte("AES,高级加密标准（英语：Advanced Encryption Standard，缩写：AES），在密码学中又称Rijndael加密法，是美国联邦政府采用的一种区块加密标准。这个标准用来替代原先的DES，已经被多方分析且广为全世界所使用。严格地说，AES和Rijndael加密法并不完全一样（虽然在实际应用中二者可以互换），因为Rijndael加密法可以支持更大范围的区块和密钥长度：AES的区块长度固定为128 比特，密钥长度则可以是128，192或256比特；而Rijndael使用的密钥和区块长度可以是32位的整数倍，以128位为下限，256比特为上限。包括AES-ECB,AES-CBC,AES-CTR,AES-OFB,AES-CFB")

func TestAesECB(t *testing.T) {
	ae := NewAesEncipher(key)
	endata, err := ae.ECB().SetPadding(PKCS5).Do(data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(endata.ToHex().ToString())
	fmt.Println(endata.ToStdBase64().ToString())
	fmt.Println(endata.ToUrlBase64().ToString())

	ad := NewAesDecipher(key)
	origin, err := ad.ECB().SetPadding(PKCS5).Do(endata)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(origin.ToString())
}

func TestAesCBC(t *testing.T) {
	ae := NewAesEncipher(key)
	endata, err := ae.CBC(iv).SetPadding(PKCS5).Do(data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(endata.ToHex().ToString())
	fmt.Println(endata.ToStdBase64().ToString())
	fmt.Println(endata.ToUrlBase64().ToString())

	ad := NewAesDecipher(key)
	origin, err := ad.CBC(iv).SetPadding(PKCS5).Do(endata)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(origin.ToString())
}

func TestAesCTR(t *testing.T) {
	ae := NewAesEncipher(key)
	endata, err := ae.CTR(iv).SetPadding(PKCS5).Do(data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(endata.ToHex().ToString())
	fmt.Println(endata.ToStdBase64().ToString())
	fmt.Println(endata.ToUrlBase64().ToString())

	ad := NewAesDecipher(key)
	origin, err := ad.CTR(iv).Do(endata)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(origin.ToString())
}

func TestAesOFB(t *testing.T) {
	ae := NewAesEncipher(key)
	endata, err := ae.OFB(iv).SetPadding(PKCS5).Do(data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(endata.ToHex().ToString())
	fmt.Println(endata.ToStdBase64().ToString())
	fmt.Println(endata.ToUrlBase64().ToString())

	ad := NewAesDecipher(key)
	origin, err := ad.OFB(iv).Do(endata)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(origin.ToString())
}

func TestAesCFB(t *testing.T) {
	ae := NewAesEncipher(key)
	endata, err := ae.CFB(iv).SetPadding(PKCS5).Do(data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(endata.ToHex().ToString())
	fmt.Println(endata.ToStdBase64().ToString())
	fmt.Println(endata.ToUrlBase64().ToString())

	ad := NewAesDecipher(key)
	origin, err := ad.CFB(iv).Do(endata)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(origin.ToString())
}
