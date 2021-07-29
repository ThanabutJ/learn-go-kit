package main

import (
	"context"

	"github.com/go-kit/kit/log"
)

func proxyingMidleware(ctx context.Context, instances string, logger log.Logger) ServiceMiddleware {

	return func(next StringService) StringService {
		return proxymw(ctx, next, retry)
	}
}
