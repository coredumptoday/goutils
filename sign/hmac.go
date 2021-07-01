package sign

import (
	"crypto"
)

func NewHmacMD5(key []byte) *signature {
	return NewSignature(crypto.MD5, key)
}

func NewHmacSHA1(key []byte) *signature {
	return NewSignature(crypto.SHA1, key)
}

func NewHmacSHA256(key []byte) *signature {
	return NewSignature(crypto.SHA256, key)
}

func NewHmacSHA256_224(key []byte) *signature {
	return NewSignature(crypto.SHA224, key)
}

func NewHmacSHA512_224(key []byte) *signature {
	return NewSignature(crypto.SHA512_224, key)
}

func NewHmacSHA512_256(key []byte) *signature {
	return NewSignature(crypto.SHA512_256, key)
}

func NewHmacSHA384(key []byte) *signature {
	return NewSignature(crypto.SHA384, key)
}

func NewHmacSHA512(key []byte) *signature {
	return NewSignature(crypto.SHA512, key)
}
