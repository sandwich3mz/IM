package utils

import (
	"IM/response/code"
	"github.com/gin-gonic/gin"
)

// GetUserID 通过请求中的 Authorization 获取 UserID
func GetUserID(c *gin.Context) (int64, error) {
	userID := c.GetInt64("user_id")
	if userID == 0 {
		return 0, code.AuthFailed
	}
	return userID, nil
}
