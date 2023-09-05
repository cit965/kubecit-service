package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Course holds the schema definition for the Course entity.
type Course struct {
	ent.Schema
}

// Fields of the Course.
func (Course) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("isRecommend").Default(false),
		field.Bool("isIntegral").Default(false),
		field.Int32("saleType"),
		field.Float32("discountPrice"),
		field.Int32("teachingType"),
		field.Int32("courseLevel"),
		field.Time("updateBy"),
		field.String("lecturerName"),
		field.Int32("purchaseCnt"),
		field.Float32("totalHour"),
		field.String("id"),
		field.String("bizCourseDetail"),
		field.String("courseCover"),
		field.String("bizCourseChapters"),
		field.Float32("salePrice"),
		field.String("bizCourseTeacher"),
		field.String("bizCourseAttachments"),
		field.Time("updateTime"),
		field.String("tags"),
		field.String("courseName"),
		field.String("createBy"),
		field.Int32("purchaseCounter"),
		field.Time("createTime"),
		field.Int32("clicks"),
		field.String("status"),
	}
}

// Edges of the Course.
func (Course) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("categories", Category.Type),
	}
}
