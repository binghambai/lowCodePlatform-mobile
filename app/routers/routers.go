package routers

import (
	service "binghambai.com/lowCodePlatform-mobile/app/service"
	"github.com/gin-gonic/gin"
)

func RegisterAllRouters(router *gin.Engine) {
	v1 := router.Group("/api/json")
	{
		v1.POST("/login", service.TestService)
		v1.POST("/test", service.TestFuncService)
	}
}
