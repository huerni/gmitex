package mws

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/huerni/gmitex/pkg/jwtToken"
	"google.golang.org/grpc/grpclog"
	"net/http"
)

const (
	jwtAudience    = "aud"
	jwtExpire      = "exp"
	jwtId          = "jti"
	jwtIssueAt     = "iat"
	jwtIssuer      = "iss"
	jwtNotBefore   = "nbf"
	jwtSubject     = "sub"
	noDetailReason = "no detail reason"
)

var (
	errInvalidToken = errors.New("invalid auth token")
	errNoClaims     = errors.New("no auth params")
)

func Authorize(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tok, err := jwtToken.ParseTokenRequest(r)
		if err != nil {
			// 返回未认证错误
			unauthorized(w, r, err)
			return
		}
		if !tok.Valid {
			unauthorized(w, r, errInvalidToken)
			return
		}
		claims, ok := tok.Claims.(jwt.MapClaims)
		if !ok {
			unauthorized(w, r, errNoClaims)
			return
		}

		ctx := r.Context()
		for k, v := range claims {
			switch k {
			case jwtAudience, jwtExpire, jwtId, jwtIssueAt, jwtIssuer, jwtNotBefore, jwtSubject:
				// ignore the standard claims
			default:
				ctx = context.WithValue(ctx, k, v)
			}
		}
		handler.ServeHTTP(w, r.WithContext(ctx))
	})
}

func unauthorized(writer http.ResponseWriter, request *http.Request, err error) {
	writer.WriteHeader(http.StatusUnauthorized)
	errorMessage := fmt.Sprintf("auth failed: %s", err.Error())

	if _, err := writer.Write([]byte(errorMessage)); err != nil {
		grpclog.Infof("Failed to write response: %v", err)
	}
}
