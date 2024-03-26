package v1

import (
	"IM/internel/controller"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(r *gin.RouterGroup) {
	user := r.Group("/user")
	{
		user.POST("/code", controller.SendCode)
		user.POST("/register", controller.Register)
		user.POST("/login", controller.Login)
	}
}
