package main

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"golang.org/x/time/rate"
	"time"
)

// loggingMidleware日志中间件
func loggingMidleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			logger.Log("msg", "开始调用")
			start := time.Now()
			defer logger.Log("msg", "结束", "cost", time.Since(start))
			return next(ctx, request)
		}
	}
}

var errRateLimit = errors.New("request rate limit")

// rateMiddleware限流中间件
func rateMiddleware(limit *rate.Limiter) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			//限流逻辑
			if limit.Allow() {
				return next(ctx, request)
			} else {
				return nil, errRateLimit
			}
		}
	}
}
