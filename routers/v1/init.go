package v1

import (
	"github.com/gin-gonic/gin"
)

var v1 *gin.RouterGroup

// SetupRouterV1 初始化路由
func SetupRouterV1(r *gin.Engine) {
	v1 = r.Group("/v1")

	InitUserRouter()

	return
}
