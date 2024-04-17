package types

type ListFriendGroupReq struct {
	GroupID int64 `json:"group_id,string"`
}

type AddFriendReq struct {
	Email   string `json:"email"`
	GroupID int64  `json:"group_id,string"`
}

type AcceptFriendReq struct {
	ID      int64 `json:"id,string"`
	GroupID int64 `json:"group_id,string"`
}

type RefuseFriendReq struct {
	ID int64 `json:"id,string"`
}

type AddFriendGroupReq struct {
	Name string `json:"name"`
}

type DeleteFriendReq struct {
	ID int64 `json:"id,string"`
}
