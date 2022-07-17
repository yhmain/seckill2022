package main

import (
	"context"
	"errors"
	"time"

	"github.com/go-kit/kit/endpoint"
	"golang.org/x/time/rate"
)

// "golang.org/x/time/rate" 该限流器是基于令牌桶限流算法实现的

var ErrLimitExceed = errors.New("Rate limit exceed!")

// NewTokenBucketLimitterWithBuildIn 使用x/time/rate创建限流 中间件
func NewTokenBucketLimitterWithBuildIn(bkt *rate.Limiter) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			// 使用Limiter的Allow方法，如果限流不放行，则直接返回限流异常
			if !bkt.Allow() {
				return nil, ErrLimitExceed
			}
			return next(ctx, request)
		}
	}
}

func DynamicLimitter(interval int, burst int) endpoint.Middleware {
	// 添加限流器，每interval秒补充一次，设置容量为burst
	bucket := rate.NewLimiter(rate.Every(time.Second*time.Duration(interval)), burst)
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if !bucket.Allow() {
				return nil, ErrLimitExceed
			}
			return next(ctx, request)
		}
	}
}
