package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
)

func NewMD5() *gHash {
	return &gHash{
		h: md5.New(),
	}
}

func NewSHA1() *gHash {
	return &gHash{
		h: sha1.New(),
	}
}

func NewSHA256() *gHash {
	return &gHash{
		h: sha256.New(),
	}
}

func NewSHA256_224() *gHash {
	return &gHash{
		h: sha256.New224(),
	}
}

func NewSHA512_224() *gHash {
	return &gHash{
		h: sha512.New512_224(),
	}
}

func NewSHA512_256() *gHash {
	return &gHash{
		h: sha512.New512_256(),
	}
}

func NewSHA384() *gHash {
	return &gHash{
		h: sha512.New384(),
	}
}

func NewSHA512() *gHash {
	return &gHash{
		h: sha512.New(),
	}
}

type gHash struct {
	h hash.Hash
}

func (gh *gHash) WriteBytes(d []byte) (n int, err error) {
	return gh.h.Write(d)
}

func (gh *gHash) WriteString(d string) (n int, err error) {
	return gh.h.Write([]byte(d))
}

func (gh *gHash) Sum() []byte {
	return gh.h.Sum(nil)
}

func (gh *gHash) SumToHex() []byte {
	encode := make([]byte, hex.EncodedLen(gh.h.Size()))
	n := hex.Encode(encode, gh.h.Sum(nil))
	return encode[:n]
}
