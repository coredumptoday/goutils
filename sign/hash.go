package sign

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
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

type xhash struct {
	h   hash.Hash
	err error
}

func (xh *xhash) Err() error {
	return xh.err
}

func (xh *xhash) WriteBytes(d []byte) int {
	if xh.err != nil {
		return 0
	}

	n, e := xh.h.Write(d)
	xh.err = e
	return n
}

func (xh *xhash) WriteString(d string) int {
	if xh.err != nil {
		return 0
	}

	n, e := xh.h.Write([]byte(d))
	xh.err = e
	return n
}

func (xh *xhash) Sum() mds {
	if xh.err != nil {
		return nil
	}

	return xh.h.Sum(nil)
}
