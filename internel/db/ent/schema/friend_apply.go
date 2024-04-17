package schema

import (
	"IM/internel/types/enums"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// FriendApply holds the schema definition for the FriendApply entity.
type FriendApply struct {
	ent.Schema
}

// Fields of the FriendApply.
func (FriendApply) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("from_user_id").StructTag(`json:"from_user_id,string"`).Comment("申请发起人 ID"),
		field.String("from_nickname").StructTag(`json:"from_nickname"`).Comment("申请发起人昵称"),
		field.String("from_face_url").StructTag(`json:"from_face_url"`).Comment("申请发起人头像"),
		field.Int64("to_user_id").StructTag(`json:"to_user_id,string"`).Comment("申请接收人 ID"),
		field.String("to_face_url").StructTag(`json:"to_face_url"`).Comment("申请接收人 昵称"),
		field.String("to_nickname").StructTag(`json:"to_nickname"`).Comment("申请接收人 头像"),
		field.String("result").GoType(enums.ApplyTypeRefuse).Default(string(enums.ApplyTypePending)).StructTag(`json:"result"`).Comment("处理结果"),
		field.String("req_msg").StructTag(`json:"req_msg"`).Default("").Comment("请求信息"),
		field.Int64("group_id").StructTag(`json:"group_id,string"`).Comment("群组 ID"),
	}
}

// Edges of the FriendApply.
func (FriendApply) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("from_user", User.Type).Ref("send_apply_user").Field("from_user_id").Unique().Required(),
		edge.From("to_user", User.Type).Ref("apply_user").Field("to_user_id").Unique().Required(),
	}
}

// Indexes of FriendApply
func (FriendApply) Indexes() []ent.Index {
	return []ent.Index{}
}

// Mixin of FriendApply
func (FriendApply) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (FriendApply) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "t_friend_apply"},
	}
}
