package sign

import (
	"net/url"
	"sort"
)

const extendKeyDefaultCap = 5

func newBaseBuilder(h *signature) *builder {
	return &builder{
		kvSep:    []byte("="),
		paramSep: []byte("&"),
		h:        h,
	}
}

type builder struct {
	h            *signature
	ks           []string
	mData        map[string]string
	qData        url.Values
	exData       map[string]string
	kvSep        []byte
	paramSep     []byte
	isMapData    bool
	isAsc        bool
	signWithoutK bool
	signWithoutV bool
}

func (b *builder) initDataFromMap(d map[string]string) {
	b.mData = d
	b.ks = make([]string, 0, len(d)+extendKeyDefaultCap)
	for k, _ := range d {
		b.ks = append(b.ks, k)
	}
}

func (b *builder) initDataFromQuery(q url.Values) {
	b.qData = q
	b.ks = make([]string, 0, len(q)+extendKeyDefaultCap)
	for k, _ := range q {
		b.ks = append(b.ks, k)
	}
}

func (b *builder) Set(k, v string) *builder {
	if b.exData == nil {
		b.exData = make(map[string]string, extendKeyDefaultCap)
	}
	b.exData[k] = v
	b.ks = append(b.ks, k)
	return b
}

func (b *builder) get(k string) string {
	if b.isMapData {
		if s, ok := b.mData[k]; ok {
			return s
		}
	} else {
		if s, ok := b.qData[k]; ok {
			return s[0]
		}
	}
	return b.exData[k]
}

func (b *builder) WritePrefixString(p ...string) *builder {
	for _, v := range p {
		b.h.WriteString(v)
	}
	return b
}

func (b *builder) WritePrefixByte(p ...[]byte) *builder {
	for _, v := range p {
		b.h.WriteBytes(v)
	}
	return b
}

func (b *builder) sort() {
	sort.Strings(b.ks)
}

func (b *builder) ASCSort() *builder {
	b.isAsc = true
	b.sort()
	return b
}

func (b *builder) DESCSort() *builder {
	b.isAsc = false
	b.sort()
	return b
}

func (b *builder) SignWithoutKey() *builder {
	b.signWithoutK = true
	return b
}

func (b *builder) SignWithoutVal() *builder {
	b.signWithoutV = true
	return b
}

func (b *builder) SetKVSepStr(sep string) *builder {
	b.kvSep = []byte(sep)
	return b
}

func (b *builder) SetKVSepByte(sep []byte) *builder {
	b.kvSep = sep
	return b
}

func (b *builder) SetParamSepStr(sep string) *builder {
	b.paramSep = []byte(sep)
	return b
}

func (b *builder) SetParamSepByte(sep []byte) *builder {
	b.paramSep = sep
	return b
}

func (b *builder) writeAsc() {
	l := len(b.ks)
	for i, k := range b.ks {
		if !b.signWithoutK {
			b.h.WriteString(k)
		}
		if b.kvSep != nil {
			b.h.WriteBytes(b.kvSep)
		}
		if !b.signWithoutV {
			b.h.WriteString(b.get(k))
		}
		if b.paramSep != nil && i < l-1 {
			b.h.WriteBytes(b.paramSep)
		}
	}
}

func (b *builder) writeDesc() {
	l := len(b.ks)
	for i := l - 1; i >= 0; i-- {
		if !b.signWithoutK {
			b.h.WriteString(b.ks[i])
		}
		if b.kvSep != nil {
			b.h.WriteBytes(b.kvSep)
		}
		if !b.signWithoutV {
			b.h.WriteString(b.get(b.ks[i]))
		}
		if b.paramSep != nil && i > 0 {
			b.h.WriteBytes(b.paramSep)
		}
	}
}

func (b *builder) write() {
	if b.isAsc {
		b.writeAsc()
	} else {
		b.writeDesc()
	}
}

func (b *builder) Sign() ([]byte, error) {
	b.write()
	return b.h.Sum()
}

func (b *builder) EqualHexString(str string) (bool, error) {
	b.write()
	return b.h.EqualHexString(str)
}

func (b *builder) SignWithPostfixStr(strs ...string) ([]byte, error) {
	b.write()

	for _, s := range strs {
		b.h.WriteString(s)
	}

	return b.h.Sum()
}

func (b *builder) SignWithPostfixByte(ba ...[]byte) ([]byte, error) {
	b.write()

	for _, s := range ba {
		b.h.WriteBytes(s)
	}

	return b.h.Sum()
}
