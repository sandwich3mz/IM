package schema

import (
	"IM/internel/types/enums"
	"IM/utils"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("nickname").Default("请设置昵称").StructTag(`json:"nickname"`).Comment("用户昵称"),
		field.String("email").Default("").StructTag(`json:"email"`).Comment("邮箱"),
		field.String("password").Default("").StructTag(`json:"password"`).Comment("密码"),
		field.Enum("status").GoType(enums.UserStatusOnline).Default(string(enums.UserStatusOffline)).StructTag(`json:"status"`).Comment("用户状态"),
		field.Time("last_online_at").Default(utils.ZeroTime).StructTag(`json:"last_online_at"`).Comment("最后在线时间"),
		field.String("avatar").StructTag(`json:"avatar"`).Default("").Comment("头像"),
		field.Int8("sex").Default(0).StructTag(`json:"sex"`).Comment("性别"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("send_msg", Msg.Type),
		edge.To("receive_msg", Msg.Type),
		edge.To("owner_user_friend", Friend.Type),
		edge.To("friend_user_friend", Friend.Type),
		edge.To("user_group", Group.Type),
		edge.To("user_group_member", GroupMember.Type),
		edge.To("send_apply_user", FriendApply.Type),
		edge.To("apply_user", FriendApply.Type),
		edge.To("friend_group", FriendGroup.Type),
	}
}

// Indexes of User
func (User) Indexes() []ent.Index {
	return []ent.Index{}
}

// Mixin of User
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "t_user"},
	}
}
