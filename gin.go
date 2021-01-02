package gregson

import (
	"github.com/gin-gonic/gin"
)

// NewApp function creates an gregson web app.
// An Gregson web app is designed to build a REST API server that can
// work with a command line client or web client. The server is
// configured with a set of integrations, to make sure the server meets
// engineering requirements like monitoring, logging, etc.
//
// Besides, Gregson engine configured some default endpoints.
//
// - POST /login provides a form with username and password.
// - GET  /metrics provides Prometheus integration
func NewGin(s *Setting) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	HookGlobalZeroLog(r)
	HookPrometheus(r, s)
	return r
}
