package wutl

import (
	"fmt"
	"net/http"
)

type HandlerFn func(next http.Handler) http.Handler

type Handler struct {
	main       http.Handler
	middleware []http.Handler
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.main == nil {
		panic("main handler not set")
	}
	h.main.ServeHTTP(w, r)
}

func (h *Handler) AddMiddleware(mws ...HandlerFn) error {
	if h.main == nil {
		return fmt.Errorf("main handler has not been set")
	}

	if len(mws) == 0 {
		return nil
	}

	var fn HandlerFn
	for i := range mws {
		idx := (len(mws) - 1) - i
		fn = mws[idx]
		h.main = fn(h.main)
	}

	return nil
}

func NewHandler(main http.Handler) *Handler {
	return &Handler{main, nil}
}
