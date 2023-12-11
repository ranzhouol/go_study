package main

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel/exporters/prometheus"
	instrument "go.opentelemetry.io/otel/metric"
	metric2 "go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/sdk/metric"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	ctx := context.Background()
	// 数据导出器，也是数据采集器
	exp, err := prometheus.New()
	if err != nil {
		log.Fatal(err)
	}

	// 新建 meter provider
	provider := metric.NewMeterProvider(metric.WithReader(exp))
	// 新建 meter
	meter := provider.Meter("prometheus")

	go serveMetrics()

	//attrs := []attribute.KeyValue{
	//	attribute.Key("A").String("B"),
	//	attribute.Key("C").String("D"),
	//}

	// 创建counter，只增不减
	counter, err := meter.Float64Counter("foo", instrument.WithDescription("counter 描述"))
	if err != nil {
		log.Fatal(err)
	}
	// 每访问localhost:2223/metrics，counter加5
	counter.Add(ctx, 5)

	// 创建gauge
	gauge, err := meter.Float64ObservableGauge("bar", instrument.WithDescription("gruge 描述"))
	if err != nil {
		log.Fatal(err)
	}
	_, err = meter.RegisterCallback(func(ctx context.Context, observer metric2.Observer) error {
		n := -10. + rand.Float64()*(90.)
		observer.ObserveFloat64(gauge, n)
		return nil
	}, gauge)

	// 创建histogram直方图
	histogram, err := meter.Float64Histogram("baz", instrument.WithDescription("histogram 描述"))
	if err != nil {
		log.Fatal(err)
	}
	histogram.Record(ctx, 23)
	histogram.Record(ctx, 7)
	histogram.Record(ctx, 50)
	histogram.Record(ctx, 101)
	histogram.Record(ctx, 105)

	ctx, _ = signal.NotifyContext(ctx, os.Interrupt)
	<-ctx.Done()
}

// 启动一个服务
func serveMetrics() {
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":2223", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
