package xcrypto

import (
	"crypto/aes"
)

const (
	AES128Length = 16
	AES192Length = 24
	AES256Length = 32
)

func initAesKey(k []byte) []byte {
	l := len(k)

	if l == AES128Length || l == AES192Length || l == AES256Length {
		return k
	}

	var dst []byte

	if l < AES128Length {
		dst = make([]byte, AES128Length)
	} else if l > AES128Length && l < AES192Length {
		dst = make([]byte, AES192Length)
	} else {
		dst = make([]byte, AES256Length)
	}

	copy(dst, k)

	return dst
}

func NewAesEncipher(key []byte) *xCipher {
	key = initAesKey(key)
	c, e := aes.NewCipher(key)
	return &xCipher{
		block:     c,
		err:       e,
		isEncrypt: true,
		pad:       NOPADDING,
	}
}

func NewAesEncipherWithKeyCheck(key []byte) (*xCipher, error) {
	c, e := aes.NewCipher(key)
	return &xCipher{
		block:     c,
		err:       e,
		isEncrypt: true,
		pad:       NOPADDING,
	}, e
}

func NewAesDecipher(key []byte) *xCipher {
	key = initAesKey(key)
	c, e := aes.NewCipher(key)
	return &xCipher{
		block:     c,
		err:       e,
		isEncrypt: false,
		pad:       NOPADDING,
	}
}

func NewAesDecipherWithKeyCheck(key []byte) (*xCipher, error) {
	c, e := aes.NewCipher(key)
	return &xCipher{
		block:     c,
		err:       e,
		isEncrypt: false,
		pad:       NOPADDING,
	}, e
}
