package sign

import (
	"crypto"
	"net/url"
)

func NewSignBuilderWithMap(ht crypto.Hash, d map[string]string) *builder {
	b := NewMdBuilder(ht)
	b.isMapData = true
	b.initDataFromMap(d)
	return b
}

func NewSignBuilderWithQuery(ht crypto.Hash, q url.Values) *builder {
	b := NewMdBuilder(ht)
	b.isMapData = false
	b.initDataFromQuery(q)
	return b
}

func NewHmacBuilderWithMap(ht crypto.Hash, key []byte, d map[string]string) *builder {
	b := NewHmacBuilder(ht, key)
	b.isMapData = true
	b.initDataFromMap(d)
	return b
}

func NewHmacBuilderWithQuery(ht crypto.Hash, key []byte, q url.Values) *builder {
	b := NewHmacBuilder(ht, key)
	b.isMapData = false
	b.initDataFromQuery(q)
	return b
}

func NewHmacBuilder(ht crypto.Hash, key []byte) *builder {
	return newBaseBuilder(NewSignature(ht, key))
}

func NewMdBuilder(ht crypto.Hash) *builder {
	return newBaseBuilder(NewSignature(ht, nil))
}
