package hasher

import (
	"math/big"
	"testing"
)

func benchmarkHasher_NonceCalc(grt int, b *testing.B) {
	p := &Payload{
		Msg:         "kek",
		ConfigValue: "lol",
		Nonce:       new(big.Int).SetInt64(0),
	}
	h := Hasher{
		Tol:        2,
		Goroutines: grt,
	}
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		h.NonceCalc(p)
	}
}

func BenchmarkHasher_NonceCalc1(b *testing.B) { benchmarkHasher_NonceCalc(1, b) }
func BenchmarkHasher_NonceCalc2(b *testing.B) { benchmarkHasher_NonceCalc(2, b) }
func BenchmarkHasher_NonceCalc3(b *testing.B) { benchmarkHasher_NonceCalc(4, b) }
func BenchmarkHasher_NonceCalc4(b *testing.B) { benchmarkHasher_NonceCalc(8, b) }
