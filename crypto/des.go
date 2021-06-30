package crypto

import (
	"crypto/des"
)

func initDesKey(k []byte) []byte {
	l := len(k)

	if l == des.BlockSize {
		return k
	}

	var dst []byte
	dst = make([]byte, des.BlockSize)

	copy(dst, k)

	return dst
}

func NewDesEncipher(key []byte) *xCipher {
	key = initDesKey(key)
	c, e := des.NewCipher(key)
	return &xCipher{
		block:     c,
		err:       e,
		isEncrypt: true,
		pad:       NOPADDING,
	}
}

func NewDesEncipherWithKeyCheck(key []byte) (*xCipher, error) {
	c, e := des.NewCipher(key)
	return &xCipher{
		block:     c,
		err:       e,
		isEncrypt: true,
		pad:       NOPADDING,
	}, e
}

func NewDesDecipher(key []byte) *xCipher {
	key = initDesKey(key)
	c, e := des.NewCipher(key)
	return &xCipher{
		block:     c,
		err:       e,
		isEncrypt: false,
		pad:       NOPADDING,
	}
}

func NewDesDecipherWithKeyCheck(key []byte) (*xCipher, error) {
	c, e := des.NewCipher(key)
	return &xCipher{
		block:     c,
		err:       e,
		isEncrypt: false,
		pad:       NOPADDING,
	}, e
}
