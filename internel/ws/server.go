package ws

import (
	"IM/internel/types"
	"IM/response"
	"IM/response/code"
	"IM/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"sync/atomic"
	"time"
)

var hub = NewHub()

func init() {
	go hub.run()
}

func WebSocketServer(c *gin.Context) {
	if err := serveWs(hub, c.Writer, c.Request); err != nil {
		logrus.Errorf("failed to serve ws, err: %v", err)
		response.RespErrorWithMsg(c, code.InvalidParams, "ws 连接失败")
		return
	}

	response.RespSuccess(c, nil)
}

// 处理来自对方的请求
func serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) error {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.Errorf("failed to upgrader conn, err: %v", err)
		return err
	}

	// 创建一个定时器，在 tokenVerifyWait 时间后关闭连接
	timer := time.AfterFunc(tokenVerifyWait, func() {
		logrus.Errorf("token verification timeout, closing connection")
		err := conn.Close()
		if err != nil {
			logrus.Errorf("failed to close conn, err: %v", err)
			return
		}
	})
	defer timer.Stop()

	// 连接上后先读取 token 并校验
	_, tokenMsg, err := conn.ReadMessage()
	if err != nil {
		logrus.Errorf("failed to read token message, err: %v", err)
		err = conn.Close()
		if err != nil {
			logrus.Errorf("failed to close conn, err: %v", err)
			return err
		}
		return err
	}
	// 解析JSON格式的token
	var tokenData types.TokenData
	if err := json.Unmarshal(tokenMsg, &tokenData); err != nil {
		logrus.Errorf("failed to parse token message, err: %v", err)
		err = conn.Close()
		if err != nil {
			logrus.Errorf("failed to close conn, err: %v", err)
			return err
		}
		return err
	}

	token := tokenData.Token
	timer.Stop()

	userID, err, _ := utils.TokenToUserID(token)
	if err != nil {
		logrus.Errorf("failed to token to user id, err: %v", err)
		err = conn.Close()
		if err != nil {
			logrus.Errorf("failed to close conn, err: %v", err)
			return err
		}
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
