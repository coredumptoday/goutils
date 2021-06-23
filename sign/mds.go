package sign

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
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

func (b mds) SHA256_224() mds {
	md := sha256.Sum224(b)
	return md[:]
}

func (b mds) SHA256() mds {
	md := sha256.Sum256(b)
	return md[:]
}

func (b mds) SHA512_224() mds {
	md := sha512.Sum512_224(b)
	return md[:]
}

func (b mds) SHA512_256() mds {
	md := sha512.Sum512_256(b)
	return md[:]
}

func (b mds) SHA384() mds {
	md := sha512.Sum384(b)
	return md[:]
}

func (b mds) SHA512() mds {
	md := sha512.Sum512(b)
	return md[:]
}
