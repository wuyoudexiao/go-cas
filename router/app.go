package router

import (
	"gin-cas/handler"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.GET("/ping", handler.Ping)
	cas := r.Group("/api/cas")
	cas.GET("/logout", handler.Logout)
	cas.GET("/call-back", handler.CallBack)
	cas.GET("/login", handler.Login)
}
