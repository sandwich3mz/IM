package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// FriendGroup holds the schema definition for the FriendGroup entity.
type FriendGroup struct {
	ent.Schema
}

// Fields of the FriendGroup.
func (FriendGroup) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("owner_id").StructTag(`json:"owner_id,string"`).Comment("分组所有者 ID"),
		field.String("group_name").StructTag(`json:"group_name"`).Comment("好友分组名称"),
	}
}

// Edges of the FriendGroup.
func (FriendGroup) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.To("friends", Friend.Type),
		edge.From("friend_group", User.Type).Ref("friend_group").Field("owner_id").Unique().Required(),
	}
}

// Indexes of FriendGroup
func (FriendGroup) Indexes() []ent.Index {
	return []ent.Index{}
}

// Mixin of FriendGroup
func (FriendGroup) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (FriendGroup) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "t_friend_group"},
	}
}
