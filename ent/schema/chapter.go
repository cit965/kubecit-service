package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Chapter holds the schema definition for the Chapter entity.
type Chapter struct {
	ent.Schema
}

// Fields of the Chapter.
func (Chapter) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("章节名称"),
		field.Time("released_time").Default(time.Now).UpdateDefault(time.Now).Comment("发布时间"),
		field.String("description").Comment("章节描述"),
		field.Int("sort").Comment("序号"),
		field.Int("course_id").Optional().Comment("课程id"),
	}
}

// Edges of the Chapter.
func (Chapter) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("lessons", Lesson.Type),
		edge.From("course", Course.Type).Ref("chapters").Unique().Field("course_id"),
	}
}
