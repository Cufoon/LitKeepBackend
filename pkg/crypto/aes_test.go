package crypto

import (
	"encoding/base64"
	"fmt"
	"testing"
)

const aesKey = "X@V1)0wY)j-BoXiiii.P=]]yFv22@RuP"
const origin = "sshhhhhhhhh55555sshhhhhhhhh5555"

func TestAesCrypto(t *testing.T) {
	aesCrypto, err := NewAesCrypto(aesKey)
	if err != nil {
		return
	}
	encrypt, err := aesCrypto.Encrypt([]byte(origin))
	if err != nil {
		return
	}
	fmt.Println(base64.StdEncoding.EncodeToString(encrypt))
	decrypt := aesCrypto.Decrypt(encrypt)
	fmt.Println(base64.StdEncoding.EncodeToString(decrypt))
	fmt.Println(origin)
	fmt.Println(string(decrypt))
}
