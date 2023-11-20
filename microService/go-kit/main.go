package main

import (
	"fmt"
	"github.com/go-kit/log"
	"net/http"
	"os"
	"ranzhouol/go_study/microService/go-kit/pkg/endpoint"
	"ranzhouol/go_study/microService/go-kit/pkg/service"
	"ranzhouol/go_study/microService/go-kit/pkg/transport"
)

func main() {
	// 1. 创建服务
	logger := log.NewLogfmtLogger(os.Stderr)
	svc := service.New(logger)
	// 2. 创建端口
	endponits := endpoint.New(svc, logger)
	// 3. 调用传输层
	handler := http.Handler(transport.NewHttpHandler(endponits))
	http.Handle("/", handler)

	// 4.启动服务
	fmt.Println(http.ListenAndServe(":9090", nil))
}
