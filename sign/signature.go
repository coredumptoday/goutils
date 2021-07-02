package sign

import (
	"bytes"
	"crypto"
	"crypto/hmac"
	"encoding/hex"
	"hash"
)

func NewSignature(h crypto.Hash, key []byte) *signature {
	if key == nil {
		return &signature{
			h: h.New(),
		}
	} else { //hmac
		return &signature{
			h:      hmac.New(h.New, key),
			isHmac: true,
		}
	}
}

type signature struct {
	h      hash.Hash
	err    error
	isHmac bool
}

func (xh *signature) Err() error {
	return xh.err
}

func (xh *signature) WriteBytes(d []byte) int {
	if xh.err != nil {
		return 0
	}

	n, e := xh.h.Write(d)
	xh.err = e
	return n
}

func (xh *signature) WriteString(d string) int {
	if xh.err != nil {
		return 0
	}

	n, e := xh.h.Write([]byte(d))
	xh.err = e
	return n
}

func (xh *signature) Sum() ([]byte, error) {
	if xh.err != nil {
		return nil, xh.err
	}

	return xh.h.Sum(nil), xh.err
}

func (xh *signature) decodeHexString(str string) []byte {
	if xh.err != nil {
		return nil
	}

	b, err := hex.DecodeString(str)
	xh.err = err
	return b
}

func (xh *signature) EqualHexString(str string) (bool, error) {
	targetSum := xh.decodeHexString(str)
	selfSum, _ := xh.Sum()
	if xh.err != nil {
		return false, xh.err
	}

	if xh.isHmac {
		return hmac.Equal(selfSum, targetSum), xh.err
	}
	return bytes.Equal(selfSum, targetSum), xh.err
}
