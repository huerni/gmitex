package gw

import "context"

type GatewayClient interface {
	AddRoute(ctx context.Context, info any)
}
