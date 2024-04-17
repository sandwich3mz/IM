package schema

import (
	"IM/internel/types/enums"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Msg holds the schema definition for the Msg entity.
type Msg struct {
	ent.Schema
}

// Fields of the User.
func (Msg) Fields() []ent.Field {
	return []ent.Field{
		field.Time("send_at").Default(time.Now()).StructTag(`json:"send_at"`).Comment("发送时间"),
		field.Enum("session_type").GoType(enums.SessionTypeUnknown).Default(string(enums.SessionTypeUnknown)).StructTag(`json:"session_type"`).Comment("会话类型"),
		field.Int64("send_id").StructTag(`json:"send_id,string"`).Comment("发送者 ID"),
		field.Int64("receive_id").StructTag(`json:"receive_id,string"`).Comment("发送者 ID"),
		field.Enum("content_type").GoType(enums.MessageTypeUnknown).Default(string(enums.MessageTypeUnknown)).StructTag(`json:"content_type"`).Comment("消息类型"),
		field.String("ack").Default("").StructTag(`json:"ack"`).Comment("确认序列号"),
		field.Int8("status").GoType(enums.MessageStatusUnknown).Default(int8(enums.MessageStatusUnknown)).StructTag(`json:"status"`).Comment("消息状态"),
		field.Text("text_elem").StructTag(`json:"text_elem"`).Default("").Comment("文本信息"),
		field.String("url").StructTag(`json:"url"`).Default("").Comment("资源地址"),
	}
}

// Edges of the Msg.
func (Msg) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("send_user", User.Type).Ref("send_msg").Field("send_id").Unique().Required(),
		edge.From("receive_user", User.Type).Ref("receive_msg").Field("receive_id").Unique().Required(),
	}
}

// Indexes of Msg
func (Msg) Indexes() []ent.Index {
	return []ent.Index{}
}

// Mixin of Msg
func (Msg) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (Msg) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "t_message"},
	}
}
