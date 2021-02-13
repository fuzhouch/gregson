package gregson

import (
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type JWTSetting struct {
	Realm string
	Key   string
}

func HookJWTLoginAPI(r *gin.Engine) {
	g := r.Group("/v1")
	jwtMiddleware := &jwt.GinJWTMiddleware{}
	g.Use(jwtMiddleware.MiddlewareFunc())
	g.GET("/refreshToken", jwtMiddleware.RefreshHandler)
	r.POST("/login", jwtMiddleware.LoginHandler)
}
