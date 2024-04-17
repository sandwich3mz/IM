package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Friend holds the schema definition for the Friend entity.
type Friend struct {
	ent.Schema
}

// Fields of the Friend.
func (Friend) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("owner_user_id").StructTag(`json:"owner_user_id,string"`).Comment("当前登录用户 ID"),
		field.Int64("friend_user_id").StructTag(`json:"friend_user_id,string"`).Comment("好友 ID"),
		field.Int8("relationship").StructTag(`json:"relationship"`).Comment("用户关系"),
		field.String("remark").StructTag(`json:"remark"`).Comment("好友备注"),
		field.Int64("group_id").StructTag(`json:"group_id,string"`).Comment("好友分组 ID"),
		field.Time("last_talk_at").StructTag(`json:"last_talk_at"`).Comment("最后交谈时间"),
	}
}

// Edges of the Friend.
func (Friend) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("owner_user", User.Type).Ref("owner_user_friend").Field("owner_user_id").Unique().Required(),
		edge.From("friend_user", User.Type).Ref("friend_user_friend").Field("friend_user_id").Unique().Required(),
		edge.From("friend_group_friend", FriendGroup.Type).Ref("friends").Field("group_id").Unique().Required(),
	}
}

// Indexes of Friend
func (Friend) Indexes() []ent.Index {
	return []ent.Index{}
}

// Mixin of Friend
func (Friend) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (Friend) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "t_friend"},
	}
}
