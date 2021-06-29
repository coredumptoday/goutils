package internal

import "crypto/cipher"

type ecb struct {
	b         cipher.Block
	blockSize int
	tmp       []byte
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
		tmp:       make([]byte, b.BlockSize()),
	}
}

type ecbEncipher ecb

func (x *ecbEncipher) BlockSize() int { return x.blockSize }

func (x *ecbEncipher) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}

	l := len(src)
	for idx := 0; idx < l; idx += x.BlockSize() {
		x.b.Encrypt(dst[idx:idx+x.BlockSize()], src[idx:idx+x.BlockSize()])
	}
}

func NewECBEncrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbEncipher)(newECB(b))
}

type ecbDecipher ecb

func (x *ecbDecipher) BlockSize() int { return x.blockSize }

func (x *ecbDecipher) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}

	l := len(src)
	for idx := 0; idx < l; idx += x.BlockSize() {
		x.b.Decrypt(dst[idx:idx+x.BlockSize()], src[idx:idx+x.BlockSize()])
	}
}

func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbDecipher)(newECB(b))
}
