package schema

import (
	"entgo.io/ent"
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
		field.String("id").Unique(),
		field.Int32("gender"),
		field.Time("birthday"),
		field.String("sysCode"),
		field.String("city"),
		field.String("country"),
		field.String("highestEducation"),
		field.Bool("isEnable"),
		field.String("password").Sensitive(),
		field.String("province"),
		field.String("updateBy"),
		field.String("createBy"),
		field.String("platformAuthUser"),
		field.String("email"),
		field.Float32("totalHour"),
		field.String("qq"),
		field.String("address"),
		field.String("nickName"),
		field.String("emergencyContact"),
		field.String("emergencyContactNumber"),
		field.String("mobile"),
		field.String("wechat"),
		field.Int32("regSource"),
		field.String("telephone"),
		field.Time("updateTime"),
		field.String("avatar"),
		field.String("realName"),
		field.Time("createTime"),
		field.String("personalSignature"),
		field.String("certificateNumber"),
		field.Int32("age"),
		field.String("residenceAddress"),
		field.String("username"),
		field.String("certificateType"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("vipMember", Member.Type),
	}
}
