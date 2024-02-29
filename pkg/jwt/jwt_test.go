package jwt

import (
	"fmt"
	"os"
	"runtime/pprof"
	"testing"

	"cufoon.litkeep.service/pkg/crypto"
)

const aesKey = "Ng1^1h:6_xVoU^Vx1X1%4Z8A_.AoT0Qn"
const edPrivateKey = "WZkixDXuIVSiHU0Fn8ire/bUTNOxFW11pWMe1/caDd6l1YgKSC6y3itj4H20zzccsIh7/dSJPzhJFdROuTQqVA=="
const edPublicKey = "pdWICkgust4rY+B9tM83HLCIe/3UiT84SRXUTrk0KlQ="

func BenchmarkTokenGenerate(b *testing.B) {
	f, err := os.Create("./profile.out")
	if err != nil {
		os.Exit(1)
		return
	}
	err = Init("qnxg.auth.a", aesKey, edPrivateKey, edPublicKey)
	if err != nil {
		os.Exit(1)
		return
	}
	err = pprof.StartCPUProfile(f)
	if err != nil {
		os.Exit(1)
		return
	}
	defer pprof.StopCPUProfile()
	fmt.Println("xxxxx", b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err != nil {
			return
		}
		b.StartTimer()
		_, err = Token(&TokenProperty{
			UserId:     "1212aabb",
			SessionId:  17,
			SignedTime: 666,
			ExpireTime: 999,
		})
		b.StopTimer()
	}
	hf, err := os.Create("./heap.out")
	if err != nil {
		os.Exit(1)
		return
	}
	err = pprof.WriteHeapProfile(hf)
	if err != nil {
		return
	}
}

func TestToken(t *testing.T) {
	err := Init("qnxg.auth.a", aesKey, edPrivateKey, edPublicKey)
	if err != nil {
		return
	}
	token, err := Token(&TokenProperty{
		UserId:     "1212aabb",
		SessionId:  17,
		SignedTime: 666,
		ExpireTime: 999,
	})
	if err != nil {
		return
	}
	fmt.Println(token)
}

func TestParse(t *testing.T) {
	err := Init("qnxg.auth.a", aesKey, edPrivateKey, edPublicKey)
	if err != nil {
		panic(err)
	}
	token, err := Token(&TokenProperty{
		UserId:     "1212aabb",
		SessionId:  17,
		SignedTime: 666,
		ExpireTime: 999,
	})
	if err != nil {
		panic(err)
	}
	data, err := Parse(token)
	if err != nil {
		panic(err)
	}
	fmt.Println("UserId\t\t", string(data.UserId))
	fmt.Println("SessionId\t", data.SessionId)
	fmt.Println("SignedTime\t", data.SignedTime)
	fmt.Println("ExpireTime\t", data.ExpireTime)
}

func TestGenerateKeyPass(t *testing.T) {
	crypto.GenerateEd25519Key()
}
