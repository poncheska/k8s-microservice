package handlers

import (
	"net/http"
)

type FastHandler struct{
	ConfigValue string
}

// just redirects get parameter "msg" to response body with config value
func (h *FastHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query().Get("msg")
	w.Write([]byte(h.ConfigValue+msg))
}
