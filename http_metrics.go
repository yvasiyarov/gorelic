package gorelic

import (
	metrics "github.com/yvasiyarov/go-metrics"
	"github.com/yvasiyarov/newrelic_platform_go"
	"time"
)

type THttpHandlerFunc func(http.ResponseWriter, *http.Request)
type THttpHandler struct {
	originalHandler     http.Handler
	originalHandlerFunc THttpHandlerFunc
	isFunc              bool
}

func NewWebHandlerFunc(h THttpHandlerFunc) *THttpHandler {
	return &THttpHandler{
		isFunc:              true,
		originalHandlerFunc: h,
	}
}
func NewHttpHandler(h http.Handler) *THttpHandler {
	return &THttpHandler{
		isFunc:          false,
		originalHandler: h,
	}
}

func (handler *THttpHandler) BeforeStart() {
}
func (handler *THttpHandler) AfterEnd() {
}

func (handler *THttpHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	handler.BeforeStart()
	defer handler.AfterEnd()

	if handler.isFunc {
		handler.originalHandlerFunc(w, req)
	} else {
		handler.originalHandler.ServeHTTP(w, req)
	}
}

func WrapHttpHandlerFunc(h THttpHandlerFunc) THttpHandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		proxy := NewHttpHandlerFunc(h)
		proxy.ServeHTTP(w, req)
	}
}

func WrapHttpHandler(h http.Handler) http.Handler {
	return NewHttpHandler(h)
}
