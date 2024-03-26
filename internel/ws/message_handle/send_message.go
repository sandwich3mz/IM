package message_handle

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
)

var EmptyData = struct{}{}

type RespMsg struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func FailedRespInvalidParams() []byte {
	resp := RespMsg{
		Code:    CodeRespFailed,
		Message: "请求参数错误",
		Data:    EmptyData,
	}
	bytes, err := json.Marshal(resp)
	if err != nil {
		logrus.Errorf("json marshal resp data failed, err: %v", err)
	}
	return bytes
}
