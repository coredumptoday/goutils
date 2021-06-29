package internal

import (
	"crypto/cipher"
)

type cfb8 struct {
	b         cipher.Block
	next      []byte
	out       []byte
	blockSize int

	decrypt bool
}

func (x *cfb8) XORKeyStream(dst, src []byte) {
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}

	for i, _ := range src {
		x.b.Encrypt(x.out, x.next)
		copy(x.next[:x.blockSize-1], x.out[1:])
		if x.decrypt {
			x.next[x.blockSize-1] = src[i]
		}

		dst[i] = src[i] ^ x.out[0]

		if !x.decrypt {
			x.next[x.blockSize-1] = dst[i]
		}
	}
}

func NewCFB8Encrypter(block cipher.Block, iv []byte) cipher.Stream {
	return newCFB8(block, iv, false)
}

func NewCFB8Decrypter(block cipher.Block, iv []byte) cipher.Stream {
	return newCFB8(block, iv, true)
}

func newCFB8(block cipher.Block, iv []byte, decrypt bool) cipher.Stream {
	blockSize := block.BlockSize()
	if len(iv) != blockSize {
		// stack trace will indicate whether it was de or encryption
		panic("cipher.newCFB: IV length must equal block size")
	}
	x := &cfb8{
		b:         block,
		out:       make([]byte, blockSize),
		next:      make([]byte, blockSize),
		blockSize: blockSize,
		decrypt:   decrypt,
	}
	copy(x.next, iv)

	return x
}
