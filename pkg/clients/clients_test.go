package clients

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	MNT_HOST = "ip172-18-0-22-c2g00ulmrepg009qa0ug-8080.direct.labs.play-with-docker.com"
	MS_HOST  = "35.244.188.167"
)

func TestFastRequest(t *testing.T) {
	tests := []struct {
		name string
		msg  string
		host string
		res  string
		ok   bool
	}{
		{
			"ok-test",
			"kek",
			"35.244.188.167",
			"kek",
			true,
		},
	}
	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			res, err := FastRequest(v.host, v.msg)
			if v.ok {
				assert.Nil(t, err)
				assert.Equal(t, v.res, res)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

func benchmarkFastRequest(addSlow int, host string, b *testing.B) {
	for i := 0; i < addSlow; i++ {
		go func() {
			SlowRequest(host, "msg")
			fmt.Println("slow req finished")
		}()
	}
	for n := 0; n < b.N; n++ {
		FastRequest(host, "msg")
	}
}

func BenchmarkFastRequest_MS_1(b *testing.B) { benchmarkFastRequest(1, MS_HOST, b) }
func BenchmarkFastRequest_MS_2(b *testing.B) { benchmarkFastRequest(1, MS_HOST, b) }
func BenchmarkFastRequest_MS_3(b *testing.B) { benchmarkFastRequest(1, MS_HOST, b) }
func BenchmarkFastRequest_MS_4(b *testing.B) { benchmarkFastRequest(1, MS_HOST, b) }
func BenchmarkFastRequest_MS_5(b *testing.B) { benchmarkFastRequest(1, MS_HOST, b) }
func BenchmarkFastRequest_MS_6(b *testing.B) { benchmarkFastRequest(1, MS_HOST, b) }
func BenchmarkFastRequest_MS_7(b *testing.B) { benchmarkFastRequest(1, MS_HOST, b) }
func BenchmarkFastRequest_MS_8(b *testing.B) { benchmarkFastRequest(1, MS_HOST, b) }

//func BenchmarkFastRequest_Monolith_1(b *testing.B) { benchmarkFastRequest(1, MNT_HOST, b) }
//func BenchmarkFastRequest_Monolith_2(b *testing.B) { benchmarkFastRequest(1, MNT_HOST, b) }
//func BenchmarkFastRequest_Monolith_3(b *testing.B) { benchmarkFastRequest(1, MNT_HOST, b) }
//func BenchmarkFastRequest_Monolith_4(b *testing.B) { benchmarkFastRequest(1, MNT_HOST, b) }
//func BenchmarkFastRequest_Monolith_5(b *testing.B) { benchmarkFastRequest(1, MNT_HOST, b) }
//func BenchmarkFastRequest_Monolith_6(b *testing.B) { benchmarkFastRequest(1, MNT_HOST, b) }
//func BenchmarkFastRequest_Monolith_7(b *testing.B) { benchmarkFastRequest(1, MNT_HOST, b) }
//func BenchmarkFastRequest_Monolith_8(b *testing.B) { benchmarkFastRequest(1, MNT_HOST, b) }
