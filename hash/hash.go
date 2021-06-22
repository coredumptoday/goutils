package xhash

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
)

type mds []byte

func (b mds) ToString() string {
	return string(b)
}

func (b mds) ToHex() mds {
	h := make([]byte, hex.EncodedLen(len(b)))
	n := hex.Encode(h, b)
	return h[:n]
}

func (b mds) ToUpper() mds {
	return bytes.ToUpper(b)
}

func (b mds) MD5() mds {
	md := md5.Sum(b)
	return md[:]
}

func (b mds) SHA1() mds {
	md := sha1.Sum(b)
	return md[:]
}

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
