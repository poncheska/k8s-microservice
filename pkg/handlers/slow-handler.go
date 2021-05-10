package handlers

import (
	"github.com/poncheska/k8s-microservice/pkg/utils/hasher"
	"math/big"
	"net/http"
)

type SlowHandler struct {
	Hr          hasher.Hasher
	ConfigValue string
}

// calc nonce for hasher.Payload struct hash with few leading zeros
func (h *SlowHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query().Get("msg")
	res := h.Hr.NonceCalc(&hasher.Payload{
		Msg:         msg,
		ConfigValue: h.ConfigValue,
		Nonce:       new(big.Int).SetInt64(0),
	})
	w.Write([]byte(res.String()))
}
