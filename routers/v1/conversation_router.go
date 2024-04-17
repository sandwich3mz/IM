package v1

import (
	"IM/internel/controller"
	"IM/middleware"
	"github.com/gin-gonic/gin"
)

func InitConversationRouter(r *gin.RouterGroup) {
	conversation := r.Group("/conversation")
	conversation.Use(middleware.TokenVer())
	{
		conversation.GET("/list", controller.ListConversation)
		conversation.GET("/msg/:id", controller.GetMessage)
		conversation.POST("/upload", controller.UploadPic)
	}
}
