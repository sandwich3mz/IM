package middleware

import (
	"IM/response"
	"IM/response/code"
	"IM/utils"
	"github.com/gin-gonic/gin"
)

const (
	UserID = "user_id"
)

func TokenVer() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userID int64
		var err error
		if token := c.GetHeader("Authorization"); token != "" {
			userID, err, _ = utils.TokenToUserID(token)
			if err != nil {
				response.RespError(c, code.AuthFailed)
				c.Abort()
				return
			}

		} else {
			response.RespError(c, code.AuthFailed)
			c.Abort()
			return
		}
		c.Set(UserID, userID)
		c.Next()
		return
	}
}

func WsTokenVer() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userID int64
		var err error
		if token := c.GetHeader("Sec-Websocket-Protocol"); token != "" {
			userID, err, _ = utils.TokenToUserID(token)
			if err != nil {
				response.RespError(c, code.AuthFailed)
				c.Abort()
				return
			}

		} else {
			response.RespError(c, code.AuthFailed)
			c.Abort()
			return
		}
		c.Set(UserID, userID)
		c.Next()
		return
	}
}
