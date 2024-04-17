package types

const RegisterCodeLength = 6

const (
	RedisEmailCodeExpireTime = 5                                        // 验证码有效期
	RedisEmailCode           = "im:user:code:action:%s:email:%s:string" // 1. 发送验证码目的  2. 邮箱
)

const (
	DefaultGroupName = "我的好友"
)

const (
	RelationshipTypeFriend = 1 // 好友
	RelationshipTypeBlock  = 2 // 已拉黑
	RelationshipTypeDelete = 0 // 已删除
)
