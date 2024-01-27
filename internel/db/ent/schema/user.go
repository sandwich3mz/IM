package schema

import (
	"IM/internel/types/enums"
	"IM/utils"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("nick_name").Default("请设置昵称").StructTag(`json:"nick_name"`).Comment("用户昵称"),
		field.String("email").Default("").StructTag(`json:"email"`).Comment("邮箱"),
		field.String("password").Default("").StructTag(`json:"password"`).Comment("密码"),
		field.Enum("status").GoType(enums.UserStatusOnline).Default(string(enums.UserStatusOffline)).StructTag(`json:"status"`).Comment("用户状态"),
		field.Time("last_online_at").Default(utils.ZeroTime).StructTag(`json:"last_online_at"`).Comment("最后在线时间"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
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
