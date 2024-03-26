package ws

import (
	"IM/response"
	"IM/response/code"
	"IM/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"sync/atomic"
)

var hub = NewHub()

func init() {
	go hub.run()
}

func WebSocketServer(c *gin.Context) {
	userID, err := utils.GetUserID(c)
	if err != nil {
		response.RespError(c, code.UnLogin)
		return
	}
	if err := serveWs(hub, c.Writer, c.Request, userID); err != nil {
		response.RespErrorWithMsg(c, code.InvalidParams, "ws 连接失败")
		return
	}

	response.RespSuccess(c, nil)
}

// 处理来自对方的请求
func serveWs(hub *Hub, w http.ResponseWriter, r *http.Request, userID int64) error {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.Errorf("failed to upgrader conn, err: %v", err)
		return err
	}

	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256), userID: userID}
	// client 编号加 1
	atomic.AddInt64(&client.hub.clientNum, 1)
	client.hub.register <- client

	go client.Write()
	go client.Read()

	return nil
}
