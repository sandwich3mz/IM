package types

import "IM/internel/db/ent"

type TokenData struct {
	Token string `json:"token"`
}

type MsgData struct {
	Message *ent.Msg `json:"message"`
}

type AckMsg struct {
	Ack string `json:"ack"`
}
