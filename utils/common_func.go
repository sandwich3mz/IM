package utils

import (
	"IM/response/code"
	"IM/utils/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

// GetUserID 通过请求中的 Authorization 获取 UserID
func GetUserID(c *gin.Context) (int64, error) {
	userID := c.GetInt64("user_id")
	if userID == 0 {
		return 0, code.AuthFailed
	}
	return userID, nil
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
)

func (e *CodeMsg) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Msg)
}
