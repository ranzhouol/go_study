package service

import (
	"context"
	"github.com/go-kit/log"
)

// 1. 定义日志中间件，包含日志logger和定义的服务Service
type loggingMiddleware struct {
	logger log.Logger
	next   Service
}

// 在原有的业务逻辑上，添加中间件的业务逻辑：打印日志
func (l loggingMiddleware) Sum(ctx context.Context, a, b int) (int, error) {
	// 打印日志（go-kit提供的log包，使用键值对）
	defer func() {
		l.logger.Log("method", "Sum", "a", a, "b", b)
	}()

	// 不改变原来的业务逻辑
	return l.next.Sum(ctx, a, b)
}

func (l loggingMiddleware) Concat(ctx context.Context, a, b string) (string, error) {
	// 打印日志（go-kit提供的log包，使用键值对）
	defer func() {
		l.logger.Log("method", "Concat", "a", a, "b", b)
	}()

	// 不改变原来的业务逻辑
	return l.next.Concat(ctx, a, b)
}

// 2. 日志中间件的初始化
type Middleware func(service Service) Service

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(service Service) Service {
		return loggingMiddleware{
			logger: logger,
			next:   service,
		}
	}
}
