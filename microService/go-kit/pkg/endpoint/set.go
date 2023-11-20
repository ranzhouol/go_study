package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"ranzhouol/go_study/microService/go-kit/pkg/service"
)

// 定义方法对应的 Request 和 Response
type SumRequest struct {
	A, B int
}

type ConcatRequest struct {
	A, B string
}

type SumResponse struct {
	V   int
	Err error
}

type ConcatResponse struct {
	V   string
	Err error
}

// 定义方法对应的端口Endpoint
type Set struct {
	SumEndpoint    endpoint.Endpoint
	ConcatEndpoint endpoint.Endpoint
}

// 在端口层实现对应的方法
func (s Set) Sum(ctx context.Context, a, b int) (int, error) {
	resp, err := s.SumEndpoint(ctx, SumRequest{A: a, B: b})
	if err != nil {
		return 0, err
	}

	// 类型断言
	response := resp.(SumResponse)
	return response.V, response.Err
}

func (s Set) Concat(ctx context.Context, a, b string) (string, error) {
	resp, err := s.ConcatEndpoint(ctx, ConcatRequest{A: a, B: b})
	if err != nil {
		return "", err
	}

	// 类型断言
	response := resp.(ConcatResponse)
	return response.V, response.Err
}

// 构造endpoint
func MakeSumEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(SumRequest)
		v, err := s.Sum(ctx, req.A, req.B)
		return SumResponse{v, err}, err
	}
}

func MakeConcatEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(ConcatRequest)
		v, err := s.Concat(ctx, req.A, req.B)
		return ConcatResponse{v, err}, err
	}
}

// 端口初始化
func New(svc service.Service, logger log.Logger) Set {
	var sumEndpoint endpoint.Endpoint
	{
		// 基础endponit
		sumEndpoint = MakeSumEndpoint(svc)

		// 包含日志中间件的endpoint
		sumEndpoint = LoggingMiddleware(logger)(sumEndpoint)
	}

	var concatEndpoint endpoint.Endpoint
	{
		// 基础endponit
		concatEndpoint = MakeConcatEndpoint(svc)

		// 包含日志中间件的endpoint
		concatEndpoint = LoggingMiddleware(logger)(concatEndpoint)
	}

	return Set{
		sumEndpoint,
		concatEndpoint,
	}
}
