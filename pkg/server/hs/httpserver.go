package hs

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"net/http"
)

type HtServer struct {
	server          *http.Server
	registerHandler interface{}
	muxOps          []runtime.ServeMuxOption
	dialOps         []grpc.DialOption
}

func NewHtServer(registerHandler interface{}) *HtServer {
	return &HtServer{registerHandler: registerHandler}
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
	//err := user.RegisterUserHandlerFromEndpoint(ctx, mux, RpcListenOn, h.dialOps)
	//if err != nil {
	//	return err
	//}

	httpmux := http.NewServeMux()
	httpmux.Handle("/", mux)
	h.server = &http.Server{
		Addr:    HttpListenOn,
		Handler: httpmux,
	}

	return h.server.ListenAndServe()
}

func (h *HtServer) Stop(ctx context.Context) error {
	fmt.Println("stop gateway server...")
	return h.server.Shutdown(ctx)
}
