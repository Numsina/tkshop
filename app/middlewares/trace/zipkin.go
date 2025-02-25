package trace

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	zkOt "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	zkHttp "github.com/openzipkin/zipkin-go/reporter/http"
)

type ZipkinMiddleware struct{}

func NewZipKinMiddleware() *ZipkinMiddleware {
	return &ZipkinMiddleware{}
}

func (z *ZipkinMiddleware) Trace() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		reporter := zkHttp.NewReporter("http://192.168.84.10:9411/api/v2/spans")

		defer reporter.Close()
		endpoint, err := zipkin.NewEndpoint("tkshop", "localhost:9988")
		if err != nil {
			return
		}
		nativetracer, err := zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(endpoint))
		if err != nil {
			return
		}
		tracer := zkOt.Wrap(nativetracer)
		startSpan := tracer.StartSpan(ctx.FullPath())
		defer startSpan.Finish()
		opentracing.SetGlobalTracer(tracer)
		ctx.Set("tracer", tracer)
		ctx.Set("startSpan", startSpan)
		c := context.WithValue(ctx.Request.Context(), "tracer", tracer)
		c = context.WithValue(c, "startSpan", startSpan)
		ctx.Request = ctx.Request.WithContext(c)
		ctx.Next()
	}
}
