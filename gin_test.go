package gregson

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
)

func TestAppCreation(t *testing.T) {
	s := NewSetting("testapp")
	// Avoid using global registry so unit test will not break
	// due to duplicated registration.
	reg := prometheus.NewRegistry()
	s.Prometheus.SetRegistry(reg)

	r := NewGin(s)

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "")
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusOK, "ResponseCodeNot200")
}

func TestMetricsPathInstalled(t *testing.T) {
	s := NewSetting("testapp")
	// Avoid using global registry so unit test will not break
	// due to duplicated registration.
	reg := prometheus.NewRegistry()
	s.Prometheus.SetRegistry(reg)

	r := NewGin(s)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/metrics", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusOK, "ResponseCodeNot200")
}
