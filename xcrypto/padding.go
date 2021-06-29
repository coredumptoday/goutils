package xcrypto

import "bytes"

type padding interface {
	padding([]byte, int) []byte
	unpadding([]byte) []byte
}

type nopadding struct{}

func (p nopadding) padding(data []byte, blockSize int) []byte {
	return data
}
func (p nopadding) unpadding(data []byte) []byte {
	return data
}

type pkcs5 struct{}

func (p pkcs5) padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}
func (p pkcs5) unpadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}

type pkcs7 struct{}

func (p pkcs7) padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}
func (p pkcs7) unpadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}

var NOPADDING = nopadding{}
var PKCS5 = pkcs5{}
var PKCS7 = pkcs7{}
