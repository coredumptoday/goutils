package xtype

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
)

type ebs struct {
	d   []byte
	err error
}

func (x *ebs) Err() error {
	return x.err
}

func (x *ebs) HexEncode() *ebs {
	if x.err != nil {
		return x
	}

	dst := make([]byte, hex.EncodedLen(len(x.d)))
	n := hex.Encode(dst, x.d)
	x.d = dst[:n]
	return x
}

func (x *ebs) HexDecode() *ebs {
	if x.err != nil {
		return x
	}

	dst := make([]byte, hex.DecodedLen(len(x.d)))
	n, err := hex.Decode(dst, x.d)
	x.d = dst[:n]
	x.err = err
	return x
}

func (x *ebs) StdBase64Encode() *ebs {
	if x.err != nil {
		return x
	}

	dst := make([]byte, base64.StdEncoding.EncodedLen(len(x.d)))
	base64.StdEncoding.Encode(dst, x.d)
	x.d = dst
	return x
}

func (x *ebs) StdBase64Decode() *ebs {
	if x.err != nil {
		return x
	}

	dst := make([]byte, base64.StdEncoding.DecodedLen(len(x.d)))
	n, err := base64.StdEncoding.Decode(dst, x.d)
	x.d = dst[:n]
	x.err = err
	return x
}

func (x *ebs) URLBase64Encode() *ebs {
	if x.err != nil {
		return x
	}

	dst := make([]byte, base64.RawURLEncoding.EncodedLen(len(x.d)))
	base64.RawURLEncoding.Encode(dst, x.d)
	x.d = dst
	return x
}

func (x *ebs) URLBase64Decode() *ebs {
	if x.err != nil {
		return x
	}

	dst := make([]byte, base64.RawURLEncoding.DecodedLen(len(x.d)))
	n, err := base64.RawURLEncoding.Decode(dst, x.d)
	x.d = dst[:n]
	x.err = err
	return x
}

func (x *ebs) ToString() (string, error) {
	return string(x.d), x.err
}

func (x *ebs) ToByteSlice() ([]byte, error) {
	return x.d, x.err
}

func (x *ebs) ToUpper() *ebs {
	if x.err != nil {
		return x
	}

	x.d = bytes.ToUpper(x.d)
	return x
}

func (x *ebs) ToLower() *ebs {
	if x.err != nil {
		return x
	}

	x.d = bytes.ToLower(x.d)
	return x
}
