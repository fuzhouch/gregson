package gregson

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

func TestGlobalPrometheusRegistration(t *testing.T) {
	s := NewSetting("testapp")
	g := gin.Default()
	HookPrometheus(g, s)

	reqCnt := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: s.Prometheus.Namespace,
			Subsystem: s.Prometheus.Subsystem,
			Name:      "requests_total",
			Help:      "N/A",
		},
		[]string{"code", "method", "handler", "host", "path"},
	)

	// Given prometheus library does not expose internal fields,
	// we have to verify our registration added to global registry
	// by intentionally inserting duplicated metrics. It should
	// trigger an internal panic, then gets caught.
	defer func() {
		if r := recover(); r != nil {
			t.Log("ExpectedCounterPanicCaught")
		} else {
			t.Fatal("ExpectedCounterPanicNotHappen")
		}
	}()
	prometheus.DefaultRegisterer.MustRegister(reqCnt)
}
