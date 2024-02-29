package util

import (
	"crypto/rand"
	"math/big"
)

var big62 = big.NewInt(62)

func GenerateBytes(l int) string {
	result := make([]byte, l)
	var item byte
	var itemBig *big.Int
	for i := 0; i < l; i++ {
		itemBig, _ = rand.Int(rand.Reader, big62)
		item = uint8(itemBig.Uint64())
		if item < 26 {
			result[i] = item + 97
		} else if item < 52 {
			result[i] = item + 39
		} else {
			result[i] = item - 4
		}
	}
	return string(result)
}

func Generate8Bytes() string {
	result := make([]byte, 8)
	var item byte
	var itemBig *big.Int
	for i := 0; i < 8; i++ {
		itemBig, _ = rand.Int(rand.Reader, big62)
		item = uint8(itemBig.Uint64())
		if item < 26 {
			result[i] = item + 97
		} else if item < 52 {
			result[i] = item + 39
		} else {
			result[i] = item - 4
		}
	}
	return string(result)
}
