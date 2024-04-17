package ws

import (
	"IM/internel/db"
	"IM/internel/db/ent"
	"IM/internel/db/ent/groupmember"
	"IM/internel/types"
	"IM/internel/types/enums"
	"IM/internel/ws/message_handle"
	"bytes"
	"context"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const (

	// 验证等待时间
	tokenVerifyWait = 10 * time.Second

	// 发送等待时间
	writeWait = 10 * time.Second

	// 允许收到下一个 pong 时间
	pongWait = 60 * time.Second

	// 发送 ping 的时间周期
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 1024
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
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
		if c.conn != nil {
			if err := c.conn.Close(); err != nil {
				logrus.Errorf("ws receive message client(%d) close failed, err: %v", c.userID, err)
			}
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
			logrus.Error(err)
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logrus.Errorf("error: %v", err)
			}
			break
		}
		logrus.Debugf("message is: %s", message)
		var msgData types.MsgData
		switch messageType {
		case websocket.PingMessage:
			if err = c.conn.WriteMessage(websocket.PingMessage, []byte("pong")); err != nil {
				logrus.Debug(1)
				continue
			}
		case websocket.CloseMessage:
			logrus.Debugf("client(%d) is closing", c.userID)
			return
		default:
			if err = json.Unmarshal(message, &msgData); err != nil {
				logrus.Errorf("ws client(%d) message unmarshal failed, err: %v", c.userID, err)
				c.send <- message_handle.FailedRespInvalidParams()
			}
		}

		if msgData.Message != nil {
			bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
			MessageHandle(msgData.Message, c)
		}
	}
}

func (c *Client) Write() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		if c.conn != nil {
			if err := c.conn.Close(); err != nil {
				logrus.Errorf("ws receive message client(%d) close failed, err: %v", c.userID, err)
			}
		}
	}()
	for {
		select {
		case message, ok := <-c.send:
			err := c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err != nil {
				logrus.Errorf("ws client setWriteDeadLine failed, err: %v", err)
				return
			}
			if !ok {
				// 关闭通道
				err := c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				if err != nil {
					logrus.Errorf("ws client write message failed, err: %v", err)
					return
				}
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			_, err = w.Write(message)
			if err != nil {
				logrus.Errorf("ws client write message failed, err: %v", err)
				return
			}

			n := len(c.send)
			for i := 0; i < n; i++ {
				_, err := w.Write(newline)
				if err != nil {
					logrus.Errorf("ws client write message failed, err: %v", err)
					return
				}
				_, err = w.Write(<-c.send)
				if err != nil {
					logrus.Errorf("ws client write message failed, err: %v", err)
					return
				}
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			err := c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err != nil {
				logrus.Errorf("ws client setWriteDeadLine failed, err: %v", err)
				return
			}
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func MessageHandle(_msg *ent.Msg, c *Client) {
	// 返回确认序列号
	sendClientInterface, ok := c.hub.clients.Load(_msg.SendID)
	if ok {
		var ackMsg types.AckMsg
		ackMsg.Ack = _msg.Ack
		ackByte, err := json.Marshal(ackMsg)
		if err != nil {
			logrus.Errorf("ws client(%d) ack message marshal failed, err: %v", c.userID, err)
			return
		}
		// 向发送方客户端返回 ack 确认信息
		if targetClient, ok := sendClientInterface.(*Client); ok {
			targetClient.send <- ackByte
		}
	}

	var respMsg types.MessageResp
	respMsg.Type = 1
	respMsg.Msg = _msg
	resp, err := json.Marshal(&respMsg)
	if err != nil {
		logrus.Errorf("ws client(%d) message marshal failed, err: %v", c.userID, err)
	}

	switch _msg.SessionType {
	// 单聊
	case enums.SessionTypeSingle:
		targetClientInterface, ok := c.hub.clients.Load(_msg.ReceiveID)
		if ok {
			// 对方客户端在线则发送信息
			if targetClient, ok := targetClientInterface.(*Client); ok {
				targetClient.send <- resp
			}
		}
	// 群聊
	case enums.SessionTypeGroup:
		members, err := db.DB.GroupMember.Query().
			Where(groupmember.GroupIDEQ(_msg.ReceiveID)).
			Select(groupmember.FieldUserID).
			All(context.Background())
		if err != nil {
			logrus.Errorf("failed to query group(%d) member, err: %v", _msg.ReceiveID, err)
			return
		}
		for _, member := range members {
			targetClientInterface, ok := c.hub.clients.Load(member.UserID)
			if ok {
				// 对方客户端在线则发送信息
				if targetClient, ok := targetClientInterface.(*Client); ok {
					targetClient.send <- resp
				}
			}
		}
	}

	_, err = db.DB.Msg.Create().
		SetSendID(_msg.SendID).
		SetSessionType(_msg.SessionType).
		SetReceiveID(_msg.ReceiveID).
		SetSendAt(_msg.SendAt).
		SetContentType(_msg.ContentType).
		SetAck(_msg.Ack).
		SetTextElem(_msg.TextElem).
		SetURL(_msg.URL).
		SetStatus(enums.MessageStatusSent).
		Save(context.Background())
	if err != nil {
		logrus.Errorf("failed to save message, err: %v", err)
		return
	}
}
