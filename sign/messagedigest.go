package sign

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
)

func NewMD5() *xhash {
	return &xhash{
		h: md5.New(),
	}
}

func NewSHA1() *xhash {
	return &xhash{
		h: sha1.New(),
	}
}

func NewSHA256() *xhash {
	return &xhash{
		h: sha256.New(),
	}
}

func NewSHA256_224() *xhash {
	return &xhash{
		h: sha256.New224(),
	}
}

func NewSHA512_224() *xhash {
	return &xhash{
		h: sha512.New512_224(),
	}
}

func NewSHA512_256() *xhash {
	return &xhash{
		h: sha512.New512_256(),
	}
}

func NewSHA384() *xhash {
	return &xhash{
		h: sha512.New384(),
	}
}

func NewSHA512() *xhash {
	return &xhash{
		h: sha512.New(),
	}
}
