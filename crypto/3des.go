package crypto

import (
	"crypto/des"
)

func initTripleDesKey(k []byte) []byte {
	l := len(k)

	if l == des.BlockSize*3 {
		return k
	}

	var dst []byte
	dst = make([]byte, des.BlockSize*3)

	copy(dst, k)

	return dst
}

func NewTripleDesEncipher(key []byte) *xCipher {
	key = initTripleDesKey(key)
	c, e := des.NewTripleDESCipher(key)
	return &xCipher{
		block:     c,
		err:       e,
		isEncrypt: true,
		pad:       NOPADDING,
	}
}

func NewTripleDesEncipherWithKeyCheck(key []byte) (*xCipher, error) {
	c, e := des.NewTripleDESCipher(key)
	return &xCipher{
		block:     c,
		err:       e,
		isEncrypt: true,
		pad:       NOPADDING,
	}, e
}

func NewTripleDesDecipher(key []byte) *xCipher {
	key = initTripleDesKey(key)
	c, e := des.NewTripleDESCipher(key)
	return &xCipher{
		block:     c,
		err:       e,
		isEncrypt: false,
		pad:       NOPADDING,
	}
}

func NewTripleDesDecipherWithKeyCheck(key []byte) (*xCipher, error) {
	c, e := des.NewTripleDESCipher(key)
	return &xCipher{
		block:     c,
		err:       e,
		isEncrypt: false,
		pad:       NOPADDING,
	}, e
}
