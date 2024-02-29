package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func GenerateEd25519Key() {
	public, private, _ := ed25519.GenerateKey(rand.Reader)
	publicKey := base64.StdEncoding.EncodeToString(public)
	privateKey := base64.StdEncoding.EncodeToString(private)
	fmt.Printf("publicKey:%s\nprivateKey:%s\n", publicKey, privateKey)
}

type Ed25519Signer struct {
	privateKey []byte
	publicKey  []byte
}

func (es *Ed25519Signer) Sign(message []byte) []byte {
	return ed25519.Sign(es.privateKey, message)
}

func (es *Ed25519Signer) Verify(message []byte, signature []byte) bool {
	return ed25519.Verify(es.publicKey, message, signature)
}

func NewEd25519Signer(private string, public string) (eds *Ed25519Signer, err error) {
	privateKey, err := base64.StdEncoding.DecodeString(private)
	if err != nil {
		return
	}
	publicKey, err := base64.StdEncoding.DecodeString(public)
	if err != nil {
		return
	}
	eds = &Ed25519Signer{
		privateKey: privateKey,
		publicKey:  publicKey,
	}
	return
}
