package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

type AesCrypto struct {
	block     cipher.Block
	blockSize int
}

func (ac *AesCrypto) Encrypt(origin []byte) (encrypted []byte, err error) {
	originData := ac.pkcs7Padding(origin)
	encrypted = make([]byte, ac.blockSize+len(originData))
	iv := encrypted[:ac.blockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}
	mode := cipher.NewCBCEncrypter(ac.block, iv)
	mode.CryptBlocks(encrypted[ac.blockSize:], originData)
	return
}

func (ac *AesCrypto) Decrypt(encrypted []byte) (origin []byte) {
	iv := encrypted[:ac.blockSize]
	encrypted = encrypted[ac.blockSize:]
	mode := cipher.NewCBCDecrypter(ac.block, iv)
	mode.CryptBlocks(encrypted, encrypted)
	origin = ac.pkcs7UnPadding(encrypted)
	return
}

func (ac *AesCrypto) pkcs7Padding(origin []byte) []byte {
	paddingNums := ac.blockSize - len(origin)%ac.blockSize
	paddingText := bytes.Repeat([]byte{byte(paddingNums)}, paddingNums)
	return append(origin, paddingText...)
}

func (ac *AesCrypto) pkcs7UnPadding(origin []byte) []byte {
	length := len(origin)
	unPadding := int(origin[length-1])
	return origin[:(length - unPadding)]
}

// NewAesCrypto
//
// key 应该为16、24或者32位，分别对应AES128、AES192和AES256
//
// 例如 key := []byte(")*N}e_D=jT5?!.d4v?__QNXG__nb~,ph")
func NewAesCrypto(key string) (ac *AesCrypto, err error) {
	k := []byte(key)
	block, err := aes.NewCipher(k)
	if err != nil {
		return
	}
	ac = &AesCrypto{
		block:     block,
		blockSize: block.BlockSize(),
	}
	return
}
