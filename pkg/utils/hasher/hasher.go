package hasher

import (
	"crypto/sha256"
	"encoding/json"
	"golang.org/x/net/context"
	"math/big"
)

type Payload struct {
	Msg         string   `json:"msg"`
	ConfigValue string   `json:"cfg"`
	Nonce       *big.Int `json:"nonce"`
}

type Hasher struct {
	Tol        int
	Goroutines int
}

func (p *Payload) Next(step *big.Int) {
	p.Nonce.Add(p.Nonce, step)
}

func (p *Payload) GetJSON() []byte {
	bs, _ := json.Marshal(p)
	return bs
}

func (h Hasher) NonceCalc(p *Payload) *big.Int {
	grt := h.Goroutines
	if h.Goroutines < 1 {
		grt = 1
	}
	step := new(big.Int).SetInt64(int64(grt))
	tol := h.Tol
	if h.Tol > 40 {
		tol = 40
	}
	if h.Tol < 1 {
		tol = 1
	}
	resChan := make(chan *big.Int, grt)
	ctx, cancel := context.WithCancel(context.Background())
	for i:=0;i<grt;i++{
		go ncRoutine(ctx, p, step, new(big.Int).SetInt64(int64(i)), tol, resChan)
	}
	res := <-resChan
	cancel()
	return res
}

func ncRoutine(ctx context.Context, p *Payload, step, initVal *big.Int, tol int, res chan *big.Int) {
	cp := &Payload{
		Msg:         p.Msg,
		ConfigValue: p.ConfigValue,
		Nonce:       new(big.Int).Set(initVal),
	}
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
		cp.Next(step)
		hash := sha256.New()
		hash.Write(cp.GetJSON())
		str := string(hash.Sum(nil))
		//fmt.Printf("init: %v nonce: %v hash: %v\n", initVal, cp.Nonce, str)
		cur := 0
		for i := 0; i < tol; i++ {
			if str[i] == '0' {
				cur++
			}
		}
		if cur == tol {
			res <- cp.Nonce
			return
		}
	}
}
