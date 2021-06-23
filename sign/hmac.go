package sign

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
)

func NewHmacMD5(key []byte) *xhash {
	return &xhash{
		h:      hmac.New(md5.New, key),
		isHmac: true,
	}
}

func NewHmacSHA1(key []byte) *xhash {
	return &xhash{
		h:      hmac.New(sha1.New, key),
		isHmac: true,
	}
}

func NewHmacSHA256(key []byte) *xhash {
	return &xhash{
		h:      hmac.New(sha256.New, key),
		isHmac: true,
	}
}

func NewHmacSHA256_224(key []byte) *xhash {
	return &xhash{
		h:      hmac.New(sha256.New224, key),
		isHmac: true,
	}
}

func NewHmacSHA512_224(key []byte) *xhash {
	return &xhash{
		h:      hmac.New(sha512.New512_224, key),
		isHmac: true,
	}
}

func NewHmacSHA512_256(key []byte) *xhash {
	return &xhash{
		h:      hmac.New(sha512.New512_256, key),
		isHmac: true,
	}
}

func NewHmacSHA384(key []byte) *xhash {
	return &xhash{
		h:      hmac.New(sha512.New384, key),
		isHmac: true,
	}
}

func NewHmacSHA512(key []byte) *xhash {
	return &xhash{
		h:      hmac.New(sha512.New, key),
		isHmac: true,
	}
}
