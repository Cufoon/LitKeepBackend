package jwt

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"errors"

	"cufoon.litkeep.service/pkg/crypto"
)

// 经过测试 AES在加密时会进行块的填充，为了最大限度的减少 Token字节数，在实际操作时，Token实际有意义
// 字节数应该在 1-31字节之间，这样可以保证最终需要加密的位数为 32个字节，加上初始化向量就是 48个字节，base64之后为 64个字节

// 因此根据数据库设计，数据设计如下
// ------------------------------------------------------------------
// |   2    |   8    |   1    |    8     |    8    |     共 27      |
// -----------------------------------------------------------------
// | 签署者  | 用户id  | 会话id  | 签署时间  | 过期时间  | 其他4字节未使用  |
// ----------------------------------------------------------------

type TokenProperty struct {
	UserId     string
	SessionId  uint8
	SignedTime int64
	ExpireTime int64
}

const prefixLength = 3

var base64Encoding = base64.RawURLEncoding

var aesCrypto *crypto.AesCrypto
var ed25519Signer *crypto.Ed25519Signer
var issuer []byte

func Init(iss string, aes string, edPrivate string, edPublic string) (err error) {
	issuer = make([]byte, 2)
	copy(issuer, iss)
	aesCrypto, err = crypto.NewAesCrypto(aes)
	if err != nil {
		return
	}
	ed25519Signer, err = crypto.NewEd25519Signer(edPrivate, edPublic)
	return
}

func sign(encoded []byte) (string, error) {
	encrypted, err := aesCrypto.Encrypt(encoded)
	if err != nil {
		return "", err
	}
	signature := ed25519Signer.Sign(encrypted)
	eLen := base64Encoding.EncodedLen(len(encrypted))
	sLen := base64Encoding.EncodedLen(len(signature))
	divideDotPosition := eLen + prefixLength
	rLen := divideDotPosition + sLen + 1
	result := make([]byte, rLen)
	copy(result[:prefixLength], "LT.")
	base64Encoding.Encode(result[prefixLength:divideDotPosition], encrypted)
	base64Encoding.Encode(result[divideDotPosition+1:], signature)
	result[divideDotPosition] = '.'
	return string(result), nil
}

func verify(token string) (decrypt []byte, err error) {
	t := []byte(token)[prefixLength:]
	tLen := len(t)
	divide := bytes.IndexByte(t, '.')
	if divide <= 0 || divide == tLen-1 {
		err = errors.New("token-broken-outer")
		return
	}
	mLen := base64Encoding.DecodedLen(divide)
	sLen := base64Encoding.DecodedLen(tLen - divide - 1)
	message := make([]byte, mLen)
	signature := make([]byte, sLen)
	_, err = base64Encoding.Decode(message, t[:divide])
	if err != nil {
		return
	}
	_, err = base64Encoding.Decode(signature, t[divide+1:])
	if err != nil {
		return
	}
	if !ed25519Signer.Verify(message, signature) {
		err = errors.New("token-unverified")
		return
	}
	decrypt = aesCrypto.Decrypt(message)
	return
}

func Token(p *TokenProperty) (token string, err error) {
	signedTimeChar := make([]byte, 8)
	binary.BigEndian.PutUint64(signedTimeChar, uint64(p.SignedTime))
	expireTimeChar := make([]byte, 8)
	binary.BigEndian.PutUint64(expireTimeChar, uint64(p.ExpireTime))
	originData := make([]byte, 27)
	copy(originData[:2], issuer)
	copy(originData[2:10], p.UserId)
	originData[10] = p.SessionId
	copy(originData[11:19], signedTimeChar)
	copy(originData[19:], expireTimeChar)
	token, err = sign(originData)
	return
}

func Parse(token string) (p *TokenProperty, err error) {
	decrypt, err := verify(token)
	if err != nil {
		return
	}
	if decrypt[0] != issuer[0] && decrypt[1] != issuer[1] {
		err = errors.New("token-broken")
		return
	}
	userId := decrypt[2:10]
	sessionId := decrypt[10]
	signedTime := int64(binary.BigEndian.Uint64(decrypt[11:19]))
	expireTime := int64(binary.BigEndian.Uint64(decrypt[19:]))
	p = &TokenProperty{
		UserId:     string(userId),
		SessionId:  sessionId,
		SignedTime: signedTime,
		ExpireTime: expireTime,
	}
	return
}
