package gormx

import (
	"fmt"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"gorm.io/gorm"
)

type JaegerTracer struct {
}

func InitJaeger() {
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
}

func NewJaegerTracer() *JaegerTracer {
	return &JaegerTracer{}
}

func (j JaegerTracer) Name() string {
	return "JaegerTracer"
}

func (j JaegerTracer) Initialize(db *gorm.DB) error {
	return j.registerAll(db)
}

func (j JaegerTracer) Before() func(*gorm.DB) {
	return func(db *gorm.DB) {
		fmt.Println(db.Statement.SQL.String())

		// 检查 db.Statement 是否为 nil，避免空指针解引用
		if db.Statement.SQL.String() == "SHOW STATUS" {
			return
		}

		var span opentracing.Span
		var parentSpan opentracing.Span
		var tracer opentracing.Tracer
		var ok bool
		//tp := db.Statement.Context.Value("tracer")
		//tracer, ok = tp.(opentracing.Tracer)
		//if ok {
		//	fmt.Println(tracer)
		//}

		s := db.Statement.Context.Value("startSpan")
		parentSpan, ok = s.(opentracing.Span)
		if ok {
			fmt.Println(s)
		}

		if db.Statement == nil {
			return
		}

		// 获取全局追踪器
		tracer = opentracing.GlobalTracer()
		if tracer == nil {
			return
		}

		if parentSpan != nil {
			// 如果有父跨度，创建一个子跨度
			span = tracer.StartSpan("exec_sql_exec", opentracing.ChildOf(parentSpan.Context()))
		} else {
			return
		}

		// 将新创建的跨度存储在 GORM 的数据库会话中
		db.Set("gorm_span", span)
	}
}

func (j JaegerTracer) After() func(*gorm.DB) {
	return func(db *gorm.DB) {
		val, ok := db.Get("gorm_span")
		if !ok {
			return
		}
		span := val.(opentracing.Span)
		span.LogEvent(db.Statement.SQL.String())
		span.Finish()
	}
}

func (j JaegerTracer) registerAll(db *gorm.DB) error {
	err := db.Callback().Create().Before("*").Register("create_before", j.Before())
	if err != nil {
		return err
	}
	err = db.Callback().Create().After("*").Register("create_After", j.After())
	if err != nil {
		return err
	}
	err = db.Callback().Update().Before("*").Register("update_before", j.Before())
	if err != nil {
		return err
	}
	err = db.Callback().Update().After("*").Register("update_After", j.After())
	if err != nil {
		return err
	}
	err = db.Callback().Delete().Before("*").Register("delete_before", j.Before())
	if err != nil {
		return err
	}
	err = db.Callback().Delete().After("*").Register("delete_After", j.After())
	if err != nil {
		return err
	}
	err = db.Callback().Query().Before("*").Register("query_before", j.Before())
	if err != nil {
		return err
	}
	err = db.Callback().Query().After("*").Register("query_After", j.After())
	if err != nil {
		return err
	}
	err = db.Callback().Raw().Before("*").Register("Raw_before", j.Before())
	if err != nil {
		return err
	}
	err = db.Callback().Raw().After("*").Register("Raw_After", j.After())
	if err != nil {
		return err
	}
	err = db.Callback().Row().Before("*").Register("Row_before", j.Before())
	if err != nil {
		return err
	}
	err = db.Callback().Row().After("*").Register("Row_After", j.After())
	if err != nil {
		return err
	}
	return nil
}
