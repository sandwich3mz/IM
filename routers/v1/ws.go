package v1

import (
	"IM/internel/ws"
	"github.com/gin-gonic/gin"
)

func InitWS(r *gin.RouterGroup) {
	chat := r.Group("/chat")
	chat.GET("/ws", ws.WebSocketServer)
}
