package dto

import (
	"IM/internel/types/enums"
	"time"
)

type UserDto struct {
	ID           int64            `json:"id,string"`
	NickName     string           `json:"nick_name"`
	Email        string           `json:"email"`
	Status       enums.UserStatus `json:"status"`
	LastOnlineAt time.Time        `json:"lastOnlineAt"`
	Avatar       string           `json:"avatar"`
	Sex          int8             `json:"sex"`
}
