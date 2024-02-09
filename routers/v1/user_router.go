package v1

import "IM/internel/controller"

func InitUserRouter() {
	user := v1.Group("/user")
	{
		user.POST("/code", controller.SendCode)
		user.POST("/register", controller.Register)
		user.POST("/login", controller.Login)
	}
}
