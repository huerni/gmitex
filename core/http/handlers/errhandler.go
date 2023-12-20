package handlers

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/huerni/gmitex/core/errno"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
	"io"
	"net/http"
)

type ErrResponse struct {
	Code    int64
	Message string
}

// 封转成HTTP Err格式
func ErrorHandler(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, writer http.ResponseWriter, request *http.Request, err error) {
	const fallback = `{"error": "failed to marshal error message"}`
	s, ok := status.FromError(err)
	if !ok {
		s, _ = status.FromError(status.Error(errno.ServiceErr.ErrCode, errno.ServiceErr.ErrMsg))
	}

	var resp = &ErrResponse{
		Code:    int64(s.Code()),
		Message: s.Message(),
	}

	contentType := marshaler.ContentType(s)
	writer.Header().Set("Content-Type", contentType)

	buf, merr := marshaler.Marshal(resp)
	if merr != nil {
		grpclog.Infof("Failed to marshal error message %q: %v", s, merr)
		writer.WriteHeader(http.StatusInternalServerError)
		if _, err := io.WriteString(writer, fallback); err != nil {
			grpclog.Infof("Failed to write response: %v", err)
		}
		return
	}

	st := 200
	writer.WriteHeader(st)
	if _, err := writer.Write(buf); err != nil {
		grpclog.Infof("Failed to write response: %v", err)
	}
}
