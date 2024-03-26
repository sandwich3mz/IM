package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// GroupMember holds the schema definition for the GroupMember entity.
type GroupMember struct {
	ent.Schema
}

// Fields of the GroupMember.
func (GroupMember) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("group_id").StructTag(`json:"group_id"`).Comment("群 ID"),
		field.Int64("user_id").StructTag(`json:"user_id"`).Comment("用户 ID"),
		field.Int8("nickname").StructTag(`json:"nickname"`).Comment("群成员昵称"),
		field.String("face_url").StructTag(`json:"face_url"`).Comment("头像"),
		field.Int64("count").StructTag(`json:"count"`).Comment("群员数量"),
		field.String("remark").StructTag(`json:"remark"`).Comment("备注"),
	}
}

// Edges of the GroupMember.
func (GroupMember) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("group_user", User.Type).Ref("user_group_member").Field("user_id").Unique().Required(),
		edge.From("member_group", Group.Type).Ref("group_member").Field("group_id").Unique().Required(),
	}
}

// Indexes of GroupMember
func (GroupMember) Indexes() []ent.Index {
	return []ent.Index{}
}

// Mixin of GroupMember
func (GroupMember) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (GroupMember) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "t_group_member"},
	}
}
