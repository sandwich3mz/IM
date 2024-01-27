package types

import "IM/internel/types/enums"

type SendCodeReq struct {
	Email  string           `json:"email"`
	Action enums.ActionType `json:"action"`
}
