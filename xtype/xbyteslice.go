package xtype

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
)

type XBS []byte

func (x XBS) MD5() XBS {
	md := md5.Sum(x)
	return md[:]
}

func (x XBS) SHA1() XBS {
	md := sha1.Sum(x)
	return md[:]
}

func (x XBS) SHA256_224() XBS {
	md := sha256.Sum224(x)
	return md[:]
}

func (x XBS) SHA256() XBS {
	md := sha256.Sum256(x)
	return md[:]
}

func (x XBS) SHA512_224() XBS {
	md := sha512.Sum512_224(x)
	return md[:]
}

func (x XBS) SHA512_256() XBS {
	md := sha512.Sum512_256(x)
	return md[:]
}

func (x XBS) SHA384() XBS {
	md := sha512.Sum384(x)
	return md[:]
}

func (x XBS) SHA512() XBS {
	md := sha512.Sum512(x)
	return md[:]
}

func (x XBS) ToString() string {
	return string(x)
}

func (x XBS) ToByteSlice() []byte {
	return x
}

func (x XBS) ToUpper() XBS {
	return bytes.ToUpper(x)
}

func (x XBS) ToLower() XBS {
	return bytes.ToLower(x)
}

func (x XBS) HexEncode() XBS {
	dst := make([]byte, hex.EncodedLen(len(x)))
	n := hex.Encode(dst, x)
	return dst[:n]
}

func (x XBS) HexDecode() *ebs {
	dst := make([]byte, hex.DecodedLen(len(x)))
	n, err := hex.Decode(dst, x)
	return &ebs{
		d:   dst[:n],
		err: err,
	}
}

func (x XBS) StdBase64Encode() XBS {
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(x)))
	base64.StdEncoding.Encode(dst, x)
	return dst
}

func (x XBS) StdBase64Decode() *ebs {
	dst := make([]byte, base64.StdEncoding.DecodedLen(len(x)))
	n, err := base64.StdEncoding.Decode(dst, x)
	return &ebs{
		d:   dst[:n],
		err: err,
	}
}

func (x XBS) URLBase64Encode() XBS {
	dst := make([]byte, base64.RawURLEncoding.EncodedLen(len(x)))
	base64.RawURLEncoding.Encode(dst, x)
	return dst
}

func (x XBS) URLBase64Decode() *ebs {
	dst := make([]byte, base64.RawURLEncoding.DecodedLen(len(x)))
	n, err := base64.RawURLEncoding.Decode(dst, x)
	return &ebs{
		d:   dst[:n],
		err: err,
	}
}
