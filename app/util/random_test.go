package util

import (
	"fmt"
	"testing"
)

func TestGenerate8Bytes(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(Generate8Bytes())
	}
}
