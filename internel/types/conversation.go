package types

import "IM/internel/db/ent"

type GetMessageReq struct {
	ID int64 `uri:"id,string"`
	PaginateReq
}

type MessageResp struct {
	Type int8     `json:"type"`
	Msg  *ent.Msg `json:"msg"`
}
