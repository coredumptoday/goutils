package xcrypto

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"

	"github.com/coredumptoday/goutils/xtype"
)

const (
	AES128Length = 16
	AES192Length = 24
	AES256Length = 32
)

var ModelNilErr = errors.New("xcrypto/aes: cipher.BlockMode and cipher.Stream is nil")

func NewAesEncipher(key []byte) *aesCipher {
	key = initAesKey(key)
	c, e := aes.NewCipher(key)
	return &aesCipher{
		block:     c,
		err:       e,
		isEncrypt: true,
		pad:       NOPADDING,
	}
}

func NewAesEncipherWithKeyRestrict(key []byte) (*aesCipher, error) {
	c, e := aes.NewCipher(key)
	return &aesCipher{
		block:     c,
		err:       e,
		isEncrypt: true,
		pad:       NOPADDING,
	}, e
}

func NewAesDecipher(key []byte) *aesCipher {
	key = initAesKey(key)
	c, e := aes.NewCipher(key)
	return &aesCipher{
		block:     c,
		err:       e,
		isEncrypt: false,
		pad:       NOPADDING,
	}
}

func NewAesDecipherWithKeyRestrict(key []byte) (*aesCipher, error) {
	c, e := aes.NewCipher(key)
	return &aesCipher{
		block:     c,
		err:       e,
		isEncrypt: false,
		pad:       NOPADDING,
	}, e
}

func initAesKey(k []byte) []byte {
	l := len(k)

	if l == AES128Length || l == AES192Length || l == AES256Length {
		return k
	}

	var dst []byte

	if l < AES128Length {
		dst = make([]byte, AES128Length)
	} else if l > AES128Length && l < AES192Length {
		dst = make([]byte, AES192Length)
	} else {
		dst = make([]byte, AES256Length)
	}

	copy(dst, k)

	return dst
}

type aesCipher struct {
	block      cipher.Block
	model      cipher.BlockMode
	stream     cipher.Stream
	err        error
	isECB      bool
	isEncrypt  bool
	pad        padding
	restrictIV bool
}

func (ac *aesCipher) ECB() *aesCipher {
	ac.isECB = true
	return ac
}

func (ac *aesCipher) SetRestrictIV() *aesCipher {
	ac.restrictIV = true
	return ac
}

func (ac *aesCipher) initIV(iv []byte, size int) []byte {
	if len(iv) == size {
		return iv
	}
	dst := make([]byte, size)
	copy(dst, iv)
	return dst
}

func (ac *aesCipher) CBC(iv []byte) *aesCipher {
	if ac.err != nil {
		return ac
	}

	if !ac.restrictIV {
		iv = ac.initIV(iv, ac.block.BlockSize())
	}

	if ac.isEncrypt {
		ac.model = cipher.NewCBCEncrypter(ac.block, iv)
	} else {
		ac.model = cipher.NewCBCDecrypter(ac.block, iv)
	}

	return ac
}

func (ac *aesCipher) CFB(iv []byte) *aesCipher {
	if ac.err != nil {
		return ac
	}

	if !ac.restrictIV {
		iv = ac.initIV(iv, ac.block.BlockSize())
	}

	if ac.isEncrypt {
		ac.stream = cipher.NewCFBEncrypter(ac.block, iv)
	} else {
		ac.stream = cipher.NewCFBDecrypter(ac.block, iv)
	}
	return ac
}

func (ac *aesCipher) CTR(iv []byte) *aesCipher {
	if ac.err != nil {
		return ac
	}

	if !ac.restrictIV {
		iv = ac.initIV(iv, ac.block.BlockSize())
	}

	ac.stream = cipher.NewCTR(ac.block, iv)
	return ac
}

func (ac *aesCipher) OFB(iv []byte) *aesCipher {
	if ac.err != nil {
		return ac
	}

	if !ac.restrictIV {
		iv = ac.initIV(iv, ac.block.BlockSize())
	}

	ac.stream = cipher.NewOFB(ac.block, iv)

	return ac
}

func (ac *aesCipher) SetPadding(p padding) *aesCipher {
	if ac.err != nil {
		return ac
	}
	ac.pad = p
	return ac
}

func (ac *aesCipher) Do(data []byte) (xtype.XBS, error) {
	if ac.err != nil {
		return nil, ac.err
	}

	if ac.isEncrypt {
		data = ac.pad.padding(data, ac.block.BlockSize())
	}

	newData := make([]byte, len(data))

	if ac.isECB {
		dataLen := len(data)
		if ac.isEncrypt {
			for idx := 0; idx < dataLen; idx += ac.block.BlockSize() {
				ac.block.Encrypt(newData[idx:idx+ac.block.BlockSize()], data[idx:idx+ac.block.BlockSize()])
			}
		} else {
			for idx := 0; idx < dataLen; idx += ac.block.BlockSize() {
				ac.block.Decrypt(newData[idx:idx+ac.block.BlockSize()], data[idx:idx+ac.block.BlockSize()])
			}
		}
	} else {
		if ac.model != nil {
			ac.model.CryptBlocks(newData, data)
		} else if ac.stream != nil {
			ac.stream.XORKeyStream(newData, data)
		} else {
			ac.err = ModelNilErr
		}
	}

	if !ac.isEncrypt {
		data = ac.pad.unpadding(newData)
	}

	return newData, ac.err
}
