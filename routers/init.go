package routers

import (
	v1 "IM/routers/v1"
	"github.com/gin-gonic/gin"
)

// SetupRouter 初始化路由
func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1.SetupRouterV1(r)

	return r
}
