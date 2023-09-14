package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Lesson holds the schema definition for the Lesson entity.
type Lesson struct {
	ent.Schema
}

// Fields of the Lesson.
func (Lesson) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("课程名称"),
		field.Time("released_time").Default(time.Now).UpdateDefault(time.Now).Comment("发布时间"),
		field.Int("sort").Comment("序号"),
		field.Int("type").Comment("课程类型"),
		field.String("storage_path").Comment("课程存储路径"),
		field.String("source").Comment("课程来源"),
		field.String("courseware").Comment("课件地址"),
		field.Int("is_free_preview").Default(2).Comment("是否免费试看"),
		field.Int("chapter_id").Optional().Comment("章节Id"),
	}
}

// Edges of the Lesson.
func (Lesson) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("chapter", Chapter.Type).Ref("lessons").Unique().Field("chapter_id"),
	}
}
