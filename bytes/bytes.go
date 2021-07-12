package bytes

import (
	"bytes"
	"crypto"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
)

type Bytes []byte

func (x Bytes) MD5() Bytes {
	md := md5.Sum(x)
	return md[:]
}

func (x Bytes) SHA1() Bytes {
	md := sha1.Sum(x)
	return md[:]
}

func (x Bytes) SHA256_224() Bytes {
	md := sha256.Sum224(x)
	return md[:]
}

func (x Bytes) SHA256() Bytes {
	md := sha256.Sum256(x)
	return md[:]
}

func (x Bytes) SHA512_224() Bytes {
	md := sha512.Sum512_224(x)
	return md[:]
}

func (x Bytes) SHA512_256() Bytes {
	md := sha512.Sum512_256(x)
	return md[:]
}

func (x Bytes) SHA384() Bytes {
	md := sha512.Sum384(x)
	return md[:]
}

func (x Bytes) SHA512() Bytes {
	md := sha512.Sum512(x)
	return md[:]
}

func (x Bytes) HMAC(hn crypto.Hash, k []byte) Bytes {
	hm := hmac.New(hn.New, k)
	hm.Write(x)
	return hm.Sum(nil)
}

func (x Bytes) DigestSum(hn crypto.Hash) Bytes {
	md := hn.New()
	md.Write(x)
	return md.Sum(nil)
}

func (x Bytes) ToString() string {
	return string(x)
}

func (x Bytes) ToByteSlice() []byte {
	return x
}

func (x Bytes) ToUpper() Bytes {
	return bytes.ToUpper(x)
}

func (x Bytes) ToLower() Bytes {
	return bytes.ToLower(x)
}

func (x Bytes) HexEncode() Bytes {
	dst := make([]byte, hex.EncodedLen(len(x)))
	n := hex.Encode(dst, x)
	return dst[:n]
}

func (x Bytes) HexDecode() (Bytes, error) {
	dst := make([]byte, hex.DecodedLen(len(x)))
	n, err := hex.Decode(dst, x)
	return dst[:n], err
}

func (x Bytes) StdBase64Encode() Bytes {
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(x)))
	base64.StdEncoding.Encode(dst, x)
	return dst
}

func (x Bytes) StdBase64Decode() (Bytes, error) {
	dst := make([]byte, base64.StdEncoding.DecodedLen(len(x)))
	n, err := base64.StdEncoding.Decode(dst, x)
	return dst[:n], err
}

func (x Bytes) URLBase64Encode() Bytes {
	dst := make([]byte, base64.RawURLEncoding.EncodedLen(len(x)))
	base64.RawURLEncoding.Encode(dst, x)
	return dst
}

func (x Bytes) URLBase64Decode() (Bytes, error) {
	dst := make([]byte, base64.RawURLEncoding.DecodedLen(len(x)))
	n, err := base64.RawURLEncoding.Decode(dst, x)
	return dst[:n], err
}
