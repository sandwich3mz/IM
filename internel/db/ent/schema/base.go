package schema

import (
	"IM/utils"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type BaseMixin struct {
	mixin.Schema
}

func (BaseMixin) Fields() []ent.Field {
	falsePtr := false
	return []ent.Field{
		field.Int64("id").
			Unique().
			Immutable().
			Annotations(entsql.Annotation{Incremental: &falsePtr}).
			DefaultFunc(func() int64 {
				return utils.GenSnowflakeInt64()
			}).
			StructTag(`json:"id,string"`).
			Comment("19 位雪花 ID"),
		field.Time("created_at").Immutable().Default(time.Now).StructTag(`json:"created_at"`).Comment("创建时刻，带时区"),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).StructTag(`json:"updated_at"`).Comment("更新时刻，带时区"),
		field.Time("deleted_at").Default(utils.ZeroTime).StructTag(`json:"deleted_at"`).Comment("软删除时刻，带时区"),
	}
}

// Annotations of the BaseMixin.
func (BaseMixin) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// Adding this annotation to the schema enables
		// comments for the table and all its fields.
		entsql.WithComments(true),
	}
}
