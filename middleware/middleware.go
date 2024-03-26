package middleware

import (
	"IM/response"
	"IM/response/code"
	"IM/utils/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

const (
	UserID = "user_id"
)

func TokenVer() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userID int64
		var err error
		if token := c.GetHeader("Authorization"); token != "" {
			userID, err, _ = TokenToUserID(token)
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

func TokenToUserID(token string) (userID int64, err error, errMsg string) {
	if token == "" {
		return 0, ValidFail, "请求头中 auth 为空"
	}
	parts := strings.SplitN(token, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return 0, ValidFail, "请求头 auth 格式错误"
	}
	mc, err := jwt.ParseToken(parts[1])
	if err != nil {
		return 0, ValidFail, "token 失效"
	}
	return mc.UserID, nil, ""
}

type CodeMsg struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

var (
	ValidFail = &CodeMsg{Code: 40000, Msg: "权限错误"}
	Fail      = &CodeMsg{50000, "失败"}
)

func (e *CodeMsg) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Msg)
}
