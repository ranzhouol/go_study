package service

import (
	"context"
	"github.com/go-kit/log"
)

// 定义 service 服务接口
// 两个方法：int类型和string类型的相加
type Service interface {
	Sum(ctx context.Context, a, b int) (int, error)
	Concat(ctx context.Context, a, b string) (string, error)
}

// 基础服务的接口实现（具体的业务逻辑）
type basicService struct{}

func (bs basicService) Sum(ctx context.Context, a, b int) (int, error) {
	return a + b, nil
}

func (bs basicService) Concat(ctx context.Context, a, b string) (string, error) {
	return a + b, nil
}

// 基础服务的初始化
func NewBasicService() Service {
	return basicService{}
}

func New(logger log.Logger) Service {
	// 1. 基础服务初始化
	svc := NewBasicService()

	// 2. 包含中间件的服务初始化
	// 第二个括号的svc用于 LoggingMiddleware返回的 type Middleware func(service Service) Service 传参
	svc = LoggingMiddleware(logger)(svc)

	return svc
}
