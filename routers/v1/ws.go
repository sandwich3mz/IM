package v1

import (
	"IM/internel/ws"
	"IM/middleware"
	"github.com/gin-gonic/gin"
)

func InitWS(r *gin.RouterGroup) {
	chat := r.Group("/chat")
	chat.Use(middleware.TokenVer())
	{
		chat.GET("/ws", ws.WebSocketServer)
	}
}
