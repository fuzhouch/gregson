package gregson

import (
	"github.com/Depado/ginprom"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

// PromSetting specifies how to create a new registry.
type PromSetting struct {
	Namespace      string
	Subsystem      string
	Path           string
	Ignore         []string
	CreateRegistry bool
	registry       *prometheus.Registry
}

func setPromDefault(s *PromSetting, namespace string) {
	s.Namespace = namespace
	s.Subsystem = "gin"
	s.Path = "/metrics"
	s.Ignore = make([]string, 1, 1)
	s.Ignore[0] = s.Path
	s.registry = nil // By default it uses global registry
}

// SetRegistry allow using a customized registry instead of
// default registry used by Promethues when creating Prometheus
// integration. This is useful in scenarios that developers must
// register registry multiple times, for example, unit test.
func (s *PromSetting) SetRegistry(r *prometheus.Registry) {
	s.registry = r
}

// SetPrometheusToGin allow integration of Prometheus to /metrics path.
func HookPrometheus(r *gin.Engine, s *Setting) {
	var reg func(*ginprom.Prometheus)
	if s.Prometheus.registry != nil {
		reg = ginprom.Registry(s.Prometheus.registry)
	} else {
		// Do nothing. Then ginprom will use default
		// global registry.
		reg = func(p *ginprom.Prometheus) {
			return
		}
	}

	g := ginprom.New(
		ginprom.Engine(r),
		reg,
		ginprom.Namespace(s.Prometheus.Namespace),
		ginprom.Subsystem(s.Prometheus.Subsystem),
		ginprom.Path(s.Prometheus.Path),
		ginprom.Ignore(s.Prometheus.Ignore...))
	r.Use(g.Instrument())
}
