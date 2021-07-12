package sign

import (
	"crypto"
)

func NewMD5() *hashWrap {
	return NewSignature(crypto.MD5, nil)
}

func NewSHA1() *hashWrap {
	return NewSignature(crypto.SHA1, nil)
}

func NewSHA256() *hashWrap {
	return NewSignature(crypto.SHA256, nil)
}

func NewSHA256_224() *hashWrap {
	return NewSignature(crypto.SHA224, nil)
}

func NewSHA512_224() *hashWrap {
	return NewSignature(crypto.SHA512_224, nil)
}

func NewSHA512_256() *hashWrap {
	return NewSignature(crypto.SHA512_256, nil)
}

func NewSHA384() *hashWrap {
	return NewSignature(crypto.SHA384, nil)
}

func NewSHA512() *hashWrap {
	return NewSignature(crypto.SHA512, nil)
}
