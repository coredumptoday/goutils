package sign

import (
	"bytes"
	"crypto"
	"crypto/hmac"
	"encoding/hex"
	"hash"
)

func NewSignature(h crypto.Hash, key []byte) *hashWrap {
	if key == nil {
		return &hashWrap{
			h: h.New(),
		}
	} else { //hmac
		return &hashWrap{
			h:      hmac.New(h.New, key),
			isHmac: true,
		}
	}
}

type hashWrap struct {
	h      hash.Hash
	err    error
	isHmac bool
}

func (xh *hashWrap) Err() error {
	return xh.err
}

func (xh *hashWrap) WriteBytes(d []byte) int {
	if xh.err != nil {
		return 0
	}

	n, _ := xh.h.Write(d)
	return n
}

func (xh *hashWrap) WriteString(d string) int {
	return xh.WriteBytes([]byte(d))
}

func (xh *hashWrap) Sum(b []byte) ([]byte, error) {
	if xh.err != nil {
		return nil, xh.err
	}

	return xh.h.Sum(b), xh.err
}

func (xh *hashWrap) decodeHexString(str string) []byte {
	if xh.err != nil {
		return nil
	}

	b, err := hex.DecodeString(str)
	xh.err = err
	return b
}

func (xh *hashWrap) EqualHexString(str string) (bool, error) {
	targetSum := xh.decodeHexString(str)
	selfSum, _ := xh.Sum(nil)
	if xh.err != nil {
		return false, xh.err
	}

	if xh.isHmac {
		return hmac.Equal(selfSum, targetSum), xh.err
	}
	return bytes.Equal(selfSum, targetSum), xh.err
}
