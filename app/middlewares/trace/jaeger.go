package trace

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

func Trace() gin.HandlerFunc {
	return func(c *gin.Context) {
		cfg := jaegercfg.Configuration{
			Sampler: &jaegercfg.SamplerConfig{
				Type:  jaeger.SamplerTypeConst,
				Param: 1,
			},
			Reporter: &jaegercfg.ReporterConfig{
				LogSpans:           true,
				LocalAgentHostPort: fmt.Sprintf("%s:%d", "192.168.84.10", 6831),
			},
			ServiceName: "tkshop",
		}

		// 可以在里接入自己实现的logger
		tracer, closer, err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
		if err != nil {
			return
		}
		defer closer.Close()
		opentracing.SetGlobalTracer(tracer)
		startSpan := tracer.StartSpan(c.FullPath())
		c.Set("tracer", tracer)
		c.Set("startSpan", startSpan)
		defer startSpan.Finish()
		ctx := context.WithValue(c.Request.Context(), "tracer", tracer)
		ctx = context.WithValue(ctx, "startSpan", startSpan)
		c.Request = c.Request.WithContext(ctx)
		c.Next()

	}
}
