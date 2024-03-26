package ws

import (
	"IM/internel/db"
	"IM/internel/db/ent"
	"IM/internel/db/ent/groupmember"
	"IM/internel/types/enums"
	"IM/internel/ws/message_handle"
	"bytes"
	"context"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	// 发送等待时间
	writeWait = 10 * time.Second

	// 允许收到下一个 pong 时间
	pongWait = 60 * time.Second

	// 发送 ping 的时间周期
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Client struct {
	hub *Hub
	// websocket 连接
	conn *websocket.Conn
	// 缓冲区
	send   chan []byte
	userID int64
}

func (c *Client) Read() {
	defer func() {
		c.hub.unregister <- c
		if err := c.conn.Close(); err != nil {
			logrus.Errorf("ws receive message client(%d) close failed, err: %v", c.userID, err)
		}
	}()

	c.conn.SetReadLimit(maxMessageSize)
	if err := c.conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		logrus.Errorf("ws client setReadDeadLine failed, err: %v", err)
	}
	c.conn.SetPongHandler(func(string) error {
		err := c.conn.SetReadDeadline(time.Now().Add(pongWait))
		if err != nil {
			logrus.Errorf("ws client setReadDeadLine failed, err: %v", err)
			return err
		}
		return nil
	})
	// 读消息
	for {
		c.conn.PongHandler()
		messageType, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logrus.Errorf("error: %v", err)
			}
			break
		}
		logrus.Debugf("message is: %s", message)
		var _msg *ent.Msg
		switch messageType {
		case websocket.PingMessage:
			if err = c.conn.WriteMessage(websocket.PingMessage, []byte("pong")); err != nil {
				continue
			}
		case websocket.CloseMessage:
			logrus.Debugf("client(%d) is closing", c.userID)
			return
		default:
			if err = json.Unmarshal(message, &_msg); err != nil {
				c.send <- message_handle.FailedRespInvalidParams()
			}
		}
		bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		MessageHandle(&message, _msg.SessionType, _msg.ReceiveID, c)
	}
}

func (c *Client) Write() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		if err := c.conn.Close(); err != nil {
			logrus.Errorf("ws receive message client(%d) close failed, err: %v", c.userID, err)
		}
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// 关闭通道
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func MessageHandle(message *[]byte, sessionType enums.SessionType, targetID int64, c *Client) {
	switch sessionType {
	// 单聊
	case enums.SessionTypeSingle:
		targetClientInterface, ok := c.hub.clients.Load(targetID)
		if ok {
			// 对方客户端在线则发送信息
			if targetClient, ok := targetClientInterface.(*Client); ok {
				targetClient.send <- *message
			}
		}
	// 群聊
	case enums.SessionTypeGroup:
		members, err := db.DB.GroupMember.Query().
			Where(groupmember.GroupIDEQ(targetID)).
			Select(groupmember.FieldUserID).
			All(context.Background())
		if err != nil {
			logrus.Errorf("failed to query group(%d) member, err: %v", targetID, err)
			return
		}
		for _, member := range members {
			targetClientInterface, ok := c.hub.clients.Load(member.UserID)
			if ok {
				// 对方客户端在线则发送信息
				if targetClient, ok := targetClientInterface.(*Client); ok {
					targetClient.send <- *message
				}
			}
		}
	}
}
