package sign

import (
	"crypto"
)

func NewMD5() *signature {
	return newSignature(crypto.MD5, nil)
}

func NewSHA1() *signature {
	return newSignature(crypto.SHA1, nil)
}

func NewSHA256() *signature {
	return newSignature(crypto.SHA256, nil)
}

func NewSHA256_224() *signature {
	return newSignature(crypto.SHA224, nil)
}

func NewSHA512_224() *signature {
	return newSignature(crypto.SHA512_224, nil)
}

func NewSHA512_256() *signature {
	return newSignature(crypto.SHA512_256, nil)
}

func NewSHA384() *signature {
	return newSignature(crypto.SHA384, nil)
}

func NewSHA512() *signature {
	return newSignature(crypto.SHA512, nil)
}
