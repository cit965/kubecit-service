package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"time"
)

// Wallet holds the schema definition for the Wallet entity.
type Wallet struct {
	ent.Schema
}

// Fields of the Wallet.
func (Wallet) Fields() []ent.Field {

	return []ent.Field{

		field.Int32("gold_leaf").SchemaType(map[string]string{
			dialect.MySQL: "int", // Override MySQL.
		}).Optional().Default(0).Comment("金叶子,分为单位"),

		field.Int32("silver_leaf").SchemaType(map[string]string{
			dialect.MySQL: "int", // Override MySQL.
		}).Optional().Default(0).Comment("银叶子,分为单位"),

		field.Int32("frozen_gold_leaf").SchemaType(map[string]string{
			dialect.MySQL: "int", // Override MySQL.
		}).Optional().Default(0).Comment("冻结金叶子,分为单位"),

		field.Int32("frozen_silver_leaf").SchemaType(map[string]string{
			dialect.MySQL: "int", // Override MySQL.
		}).Optional().Default(0).Comment("冻结银叶子,分为单位"),

		field.Int32("user_id").SchemaType(map[string]string{
			dialect.MySQL: "int", // Override MySQL.
		}).Optional().Comment("用户ID"),

		field.String("username").SchemaType(map[string]string{
			dialect.MySQL: "varchar(16)", // Override MySQL.
		}).Optional().Comment("用户名"),

		field.Time("created_at").SchemaType(map[string]string{
			dialect.MySQL: "datetime", // Override MySQL.
		}).Default(time.Now).Comment("创建时间"),

		field.Time("updated_at").SchemaType(map[string]string{
			dialect.MySQL: "datetime", // Override MySQL.
		}).Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}

}

// Edges of the Wallet.
func (Wallet) Edges() []ent.Edge {
	return nil
}
