package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// ApplyRecord holds the schema definition for the ApplyRecord entity.
type ApplyRecord struct {
	ent.Schema
}

// Fields of the ApplyRecord.
func (ApplyRecord) Fields() []ent.Field {
	return []ent.Field{
		field.Text("detail").Optional().Comment("讲师详情"),
		field.Text("curriculum_vitae").Optional().Comment("履历描述"),
		field.Text("works").Optional().Comment("以往作品"),
		field.String("skills").Optional().Comment("技能点"),
		field.String("name").Comment("名字"),
		field.Int("level").Comment("级别"),
		field.String("avatar").Optional().Comment("头像"),
		field.Time("create_at").Default(time.Now).Comment("创建时间"),
		field.Time("update_at").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
		field.Int8("is_passed").Default(2).Comment("审核结果"),
		field.Text("messages").Optional().Comment("审核人意见"),
		field.Int("auditor_id").Optional().Comment("审核人意见"),
		field.Int("user_id").Optional(),
	}
}

// Edges of the ApplyRecord.
func (ApplyRecord) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("apply_record").Field("user_id").Unique(),
	}
}
