package v1

import (
	"IM/internel/controller"
	"IM/middleware"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(r *gin.RouterGroup) {
	r.POST("/code", controller.SendCode)
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	r.POST("/modify", controller.ModifyPwd)
	user := r.Group("/user")
	user.Use(middleware.TokenVer())
	{
		user.GET("", controller.GetOneInfo)
		user.PUT("", controller.UpdateOneInfo)
	}
}
