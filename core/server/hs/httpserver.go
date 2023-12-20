package hs

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"net/http"
)

type Middleware func(http.Handler) http.Handler

type HtServer struct {
	server          *http.Server
	registerHandler interface{}
	muxOps          []runtime.ServeMuxOption
	dialOps         []grpc.DialOption
	middlewares     map[string][]Middleware
}

func NewHtServer(registerHandler interface{}) *HtServer {
	mws := map[string][]Middleware{"/": make([]Middleware, 0)}
	return &HtServer{registerHandler: registerHandler, middlewares: mws}
}

func (h *HtServer) AddMuxOp(opts ...runtime.ServeMuxOption) {
	for _, opt := range opts {
		h.muxOps = append(h.muxOps, opt)
	}
}

func (h *HtServer) AddDialOps(opts ...grpc.DialOption) {
	for _, opt := range opts {
		h.dialOps = append(h.dialOps, opt)
	}
}

func (h *HtServer) Start(ctx context.Context, RpcListenOn string, HttpListenOn string) error {
	mux := runtime.NewServeMux(h.muxOps...)
	switch handler := h.registerHandler.(type) {
	case func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error:
		err := handler(ctx, mux, RpcListenOn, h.dialOps)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported function signature for registerHandler")
	}

	httpmux := http.NewServeMux()
	h.applyMws(httpmux, mux)
	h.server = &http.Server{
		Addr:    HttpListenOn,
		Handler: httpmux,
	}

	return h.server.ListenAndServe()
}

func (h *HtServer) AddMws(pattern string, middlewares ...Middleware) {
	_, ok := h.middlewares[pattern]
	if !ok {
		h.middlewares[pattern] = make([]Middleware, 0)
	}
	h.middlewares[pattern] = append(h.middlewares[pattern], middlewares...)
}

func (h *HtServer) applyMws(httpmux *http.ServeMux, mux *runtime.ServeMux) {
	for pattern, mws := range h.middlewares {
		var handler http.Handler = mux
		for i := len(mws) - 1; i >= 0; i-- {
			handler = mws[i](handler)
		}
		httpmux.Handle(pattern, handler)
	}
}

func (h *HtServer) Stop(ctx context.Context) error {
	fmt.Println("stop gateway server...")
	return h.server.Shutdown(ctx)
}
