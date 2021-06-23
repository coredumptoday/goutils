package sign

import (
	"crypto"
	"net/url"
)

func NewSignBuilderWithMap(ht crypto.Hash, d map[string]string) *builder {
	b := newMdBuilder(ht)
	b.isMapData = true
	b.initDataFromMap(d)
	return b
}

func NewSignBuilderWithQuery(ht crypto.Hash, q url.Values) *builder {
	b := newMdBuilder(ht)
	b.isMapData = false
	b.initDataFromQuery(q)
	return b
}

func NewHmacBuilderWithMap(ht crypto.Hash, key []byte, d map[string]string) *builder {
	b := newHmacBuilder(ht, key)
	b.isMapData = true
	b.initDataFromMap(d)
	return b
}

func NewHmacBuilderWithQuery(ht crypto.Hash, key []byte, q url.Values) *builder {
	b := newHmacBuilder(ht, key)
	b.isMapData = false
	b.initDataFromQuery(q)
	return b
}

func newHmacBuilder(ht crypto.Hash, key []byte) *builder {
	b := newBaseBuilder()

	switch ht {
	case crypto.MD5:
		b.h = NewHmacMD5(key)
	case crypto.SHA1:
		b.h = NewHmacSHA1(key)
	case crypto.SHA224:
		b.h = NewHmacSHA256_224(key)
	case crypto.SHA256:
		b.h = NewHmacSHA256(key)
	case crypto.SHA512_224:
		b.h = NewHmacSHA512_224(key)
	case crypto.SHA512_256:
		b.h = NewHmacSHA512_256(key)
	case crypto.SHA384:
		b.h = NewHmacSHA384(key)
	case crypto.SHA512:
		b.h = NewHmacSHA512(key)
	}
	return b
}

func newMdBuilder(ht crypto.Hash) *builder {
	b := newBaseBuilder()

	switch ht {
	case crypto.MD5:
		b.h = NewMD5()
	case crypto.SHA1:
		b.h = NewSHA1()
	case crypto.SHA224:
		b.h = NewSHA256_224()
	case crypto.SHA256:
		b.h = NewSHA256()
	case crypto.SHA512_224:
		b.h = NewSHA512_224()
	case crypto.SHA512_256:
		b.h = NewSHA512_256()
	case crypto.SHA384:
		b.h = NewSHA384()
	case crypto.SHA512:
		b.h = NewSHA512()
	}
	return b
}
