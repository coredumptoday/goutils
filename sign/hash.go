package sign

import (
	"crypto/hmac"
	"encoding/hex"
	"hash"
	"reflect"

	"github.com/coredumptoday/goutils/xtype"
)

type xhash struct {
	h      hash.Hash
	err    error
	isHmac bool
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

func (xh *xhash) Sum() (xtype.XBS, error) {
	if xh.err != nil {
		return nil, xh.err
	}

	return xh.h.Sum(nil), xh.err
}

func (xh *xhash) decodeHexString(str string) []byte {
	if xh.err != nil {
		return nil
	}

	b, err := hex.DecodeString(str)
	xh.err = err
	return b
}

func (xh *xhash) EqualHexString(str string) (bool, error) {
	targetSum := xh.decodeHexString(str)
	selfSum, _ := xh.Sum()
	if xh.err != nil {
		return false, xh.err
	}

	if xh.isHmac {
		return hmac.Equal(selfSum, targetSum), xh.err
	}
	return reflect.DeepEqual(selfSum, targetSum), xh.err
}
