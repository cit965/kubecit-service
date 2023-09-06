package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Slider holds the schema definition for the Slider entity.
type Slider struct {
	ent.Schema
}

// Fields of the Slider.
func (Slider) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty().Comment("轮播图标题"),
		field.String("content").NotEmpty().Comment("轮播图内容简介"),
		field.String("image_link").NotEmpty().Comment("轮播图链接"),
		field.Time("create_at").Default(time.Now()).Comment("创建时间"),
		field.Time("update_at").Default(time.Now()).UpdateDefault(time.Now()).Comment("修改时间"),
		field.Bool("is_valid").Default(true).Comment("是否有效"),
		field.Int("priority").Comment("优先级"),
	}
}

// Edges of the Slider.
func (Slider) Edges() []ent.Edge {
	return nil
}
