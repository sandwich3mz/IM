package v1

import (
	"IM/internel/controller"
	"IM/middleware"
	"github.com/gin-gonic/gin"
)

func InitFriendRouter(r *gin.RouterGroup) {
	friend := r.Group("/friend")
	friend.Use(middleware.TokenVer())
	{
		friend.GET("/list", controller.ListFriendGroup)
		friend.DELETE("", controller.DeleteFriend)
	}
	group := friend.Group("/group")
	{
		group.POST("/add", controller.AddFriendGroup)
	}
	apply := friend.Group("/apply")
	{
		apply.GET("/list", controller.ListFriendApply)
		apply.POST("/add", controller.AddFriend)
		apply.PUT("/accept", controller.AcceptApply)
		apply.PUT("/refuse", controller.RefuseApply)
	}
}
