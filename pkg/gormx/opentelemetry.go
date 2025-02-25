package gormx

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"time"
)

//todo

func newJaegerTraceProvider(ctx context.Context) (*trace.TracerProvider, error) {
	// 使用http连接jaeger暴露的端口
	exporter, err := otlptracehttp.New(ctx, otlptracehttp.WithEndpoint("192.168.84.10:6831"),
		otlptracehttp.WithInsecure())
	if err != nil {
		return nil, err
	}
	res, err := resource.New(ctx, resource.WithAttributes(semconv.ServiceName("tkshop")))
	if err != nil {
		return nil, err
	}
	traceProvider := trace.NewTracerProvider(trace.WithResource(res), trace.WithSampler(trace.AlwaysSample()), trace.WithBatcher(exporter, trace.WithBatchTimeout(time.Second)))

	return traceProvider, nil
}

func InitTracer(ctx context.Context) (*trace.TracerProvider, error) {
	tp, err := newJaegerTraceProvider(ctx)
	if err != nil {
		return nil, err
	}

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tp, nil
}
