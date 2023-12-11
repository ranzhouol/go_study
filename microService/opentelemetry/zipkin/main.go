package main

import (
	"context"
	"flag"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"log"
)

const (
	service    = "trace-demo-zipkin" //服务名
	enviroment = "production"        //环境
	id         = 1                   //进程Id
)

func traceProvider(url string) (*tracesdk.TracerProvider, error) {
	// 创建 OTel exporter
	exp, err := zipkin.New(url)
	if err != nil {
		log.Fatal(err)
	}

	tp := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exp),
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			// 服务名
			semconv.ServiceName(service),
			// 属性attribute
			attribute.String("enviroment", enviroment),
			attribute.Int64("id", id),
		),
		),
	)

	return tp, nil
}

func main() {
	// zipkin接受数据的地址
	url := flag.String("zipkin", "http://192.168.239.100:9411/api/v2/spans", "")

	// 创建TracerProvider
	tp, err := traceProvider(*url)
	if err != nil {
		log.Fatal(err)
	}

	// 创建上下文
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	defer func(ctx context.Context) {
		// 关闭TracerProvider
		err := tp.Shutdown(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(ctx)
	otel.SetTracerProvider(tp)

	// 创建tracer
	tr := tp.Tracer("component-main")
	// 开始记录, 指定span名，创建span
	ctx, span := tr.Start(ctx, "foo")
	defer span.End()
	bar(ctx)
}

func bar(ctx context.Context) {
	tr := otel.Tracer("component-bar")
	_, span := tr.Start(ctx, "bar")
	defer span.End()

	//业务逻辑
	span.SetAttributes(attribute.Key("testset").String("value"))
}
