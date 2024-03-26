package types

import (
	"IM/internel/types/enums"
)

type SendCodeReq struct {
	Email  string           `json:"email"`
	Action enums.ActionType `json:"action"`
}

type RegisterReq struct {
	Email      string `json:"email"`
	Code       string `json:"code"`
	NickName   string `json:"nick_name"`
	Password   string `json:"password"`
	ConfirmPwd string `json:"confirm_pwd"`
}

type LoginReq struct {
	Email    string         `json:"email"`
	Code     string         `json:"code"`
	Password string         `json:"password"`
	Way      enums.LoginWay `json:"way" binding:"required"`
}
