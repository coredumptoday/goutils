package crypto

import (
	"crypto/cipher"
	"errors"
	"fmt"

	"github.com/coredumptoday/goutils/crypto/internal/model"
	"github.com/coredumptoday/goutils/crypto/internal/stream"
)

var ErrPadIsNil = errors.New("crypto/cipher: padding processer is not set")
var ErrPadSetNo = errors.New("crypto/cipher: block model pad must not nopadding")
var ErrModelErr = errors.New("crypto/cipher: crypto's model set fail")

type xCipher struct {
	block     cipher.Block
	model     cipher.BlockMode
	stream    cipher.Stream
	err       error
	isEncrypt bool
	isModel   bool
	pad       padding
}

func (c *xCipher) checkIV(iv []byte) {
	if len(iv) != c.block.BlockSize() {
		c.err = fmt.Errorf("crypto/cipher: IV length must equal block size, blockSize length: %v bytes", c.block.BlockSize())
	}
}

func (c *xCipher) ECB() *xCipher {
	if c.err != nil {
		return c
	}

	if c.isEncrypt {
		c.model = model.NewECBEncrypter(c.block)
	} else {
		c.model = model.NewECBDecrypter(c.block)
	}
	c.isModel = true

	return c
}

func (c *xCipher) CBC(iv []byte) *xCipher {
	c.checkIV(iv)
	if c.err != nil {
		return c
	}

	if c.isEncrypt {
		c.model = cipher.NewCBCEncrypter(c.block, iv)
	} else {
		c.model = cipher.NewCBCDecrypter(c.block, iv)
	}
	c.isModel = true

	return c
}

func (c *xCipher) CFB(iv []byte) *xCipher {
	c.checkIV(iv)
	if c.err != nil {
		return c
	}

	if c.isEncrypt {
		c.stream = cipher.NewCFBEncrypter(c.block, iv)
	} else {
		c.stream = cipher.NewCFBDecrypter(c.block, iv)
	}
	c.isModel = false

	return c
}

func (c *xCipher) CTR(iv []byte) *xCipher {
	c.checkIV(iv)
	if c.err != nil {
		return c
	}

	c.stream = cipher.NewCTR(c.block, iv)
	c.isModel = false
	return c
}

func (c *xCipher) OFB(iv []byte) *xCipher {
	c.checkIV(iv)
	if c.err != nil {
		return c
	}

	c.stream = cipher.NewOFB(c.block, iv)
	c.isModel = false

	return c
}

func (c *xCipher) OFB8(iv []byte) *xCipher {
	c.checkIV(iv)
	if c.err != nil {
		return c
	}

	c.stream = stream.NewOFB8(c.block, iv)
	c.isModel = false

	return c
}

func (c *xCipher) CFB8(iv []byte) *xCipher {
	c.checkIV(iv)
	if c.err != nil {
		return c
	}

	if c.isEncrypt {
		c.stream = stream.NewCFB8Encrypter(c.block, iv)
	} else {
		c.stream = stream.NewCFBDecrypter(c.block, iv)
	}
	c.isModel = false

	return c
}

func (c *xCipher) SetPadding(p padding) *xCipher {
	if c.err != nil {
		return c
	}
	c.pad = p
	return c
}

func (c *xCipher) Do(data []byte) ([]byte, error) {
	if c.err != nil {
		return nil, c.err
	}

	if c.isModel {
		if c.pad == nil {
			return nil, ErrPadIsNil
		}
		if _, ok := c.pad.(nopadding); ok {
			return nil, ErrPadSetNo
		}
	}

	if c.isEncrypt {
		data = c.pad.padding(data, c.block.BlockSize())
	}

	dst := make([]byte, len(data))

	if c.isModel && c.model != nil {
		c.model.CryptBlocks(dst, data)
	} else if !c.isModel && c.stream != nil {
		c.stream.XORKeyStream(dst, data)
	} else {
		c.err = ErrModelErr
	}

	if !c.isEncrypt {
		dst = c.pad.unpadding(dst)
	}

	return dst, c.err
}
