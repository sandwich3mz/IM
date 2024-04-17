package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Group holds the schema definition for the Group entity.
type Group struct {
	ent.Schema
}

// Fields of the Group.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.String("group_name").StructTag(`json:"group_name"`).Comment("群聊名称"),
		field.String("notification").StructTag(`json:"notification"`).Comment("群公告"),
		field.Int8("relationship").StructTag(`json:"relationship"`).Comment("用户关系"),
		field.Int64("owner_user_id").StructTag(`json:"owner_user_id,string"`).Comment("群主 ID"),
		field.String("face_url").StructTag(`json:"face_url"`).Comment("头像"),
		field.Int64("count").StructTag(`json:"count"`).Comment("群员数量"),
	}
}

// Edges of the Group.
func (Group) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("owner_user", User.Type).Ref("user_group").Field("owner_user_id").Unique().Required(),
		edge.To("group_member", GroupMember.Type),
	}
}

// Indexes of Group
func (Group) Indexes() []ent.Index {
	return []ent.Index{}
}

// Mixin of Group
func (Group) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (Group) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "t_group"},
	}
}
