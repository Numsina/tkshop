package prome

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"strconv"
	"time"
)

type Metrics struct {
	Namespace  string
	Subsystem  string
	Name       string
	Help       string
	InstanceID string
}

func NewMetrics(namespace string, instanceID string, name string, subsystem string, help string) *Metrics {
	return &Metrics{Namespace: namespace, InstanceID: instanceID, Name: name, Subsystem: subsystem, Help: help}
}

func (m *Metrics) Build() gin.HandlerFunc {

	labels := []string{"method", "pattern", "status_code"}
	// 统计请求响应时间
	summaryVec := prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Namespace: m.Namespace,
		Subsystem: m.Subsystem,
		Name:      m.Name + "_resp_time",
		Help:      m.Help,
		ConstLabels: map[string]string{
			"instance_id": m.InstanceID,
		},
		Objectives: map[float64]float64{
			0.5:   0.1,
			0.9:   0.01,
			0.99:  0.001,
			0.999: 0.0001,
		},
	}, labels)

	err := prometheus.Register(summaryVec)
	if err != nil {
		log.Println(err)
		return nil
	}
	// 统计当前活跃请求数
	gauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: m.Namespace,
		Subsystem: m.Subsystem,
		Name:      m.Name + "active_request",
	})

	prometheus.Register(gauge)

	return func(c *gin.Context) {
		start := time.Now()
		gauge.Inc()
		defer func() {
			gauge.Dec()
			pattern := c.FullPath()
			if pattern == "" {
				pattern = "unknown"
			}

			end := time.Since(start).Milliseconds()
			summaryVec.WithLabelValues(c.Request.Method, pattern, strconv.Itoa(c.Writer.Status())).Observe(float64(end))
		}()
		c.Next()

	}
}
