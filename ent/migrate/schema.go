// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AccountsColumns holds the columns for the "accounts" table.
	AccountsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_id", Type: field.TypeUint64},
		{Name: "openid", Type: field.TypeString, Size: 32},
		{Name: "password", Type: field.TypeString, Size: 32},
		{Name: "method", Type: field.TypeString, Size: 32},
	}
	// AccountsTable holds the schema information for the "accounts" table.
	AccountsTable = &schema.Table{
		Name:       "accounts",
		Columns:    AccountsColumns,
		PrimaryKey: []*schema.Column{AccountsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "account_openid_method",
				Unique:  true,
				Columns: []*schema.Column{AccountsColumns[2], AccountsColumns[4]},
			},
		},
	}
	// CategoriesColumns holds the columns for the "categories" table.
	CategoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "level", Type: field.TypeInt},
		{Name: "parent_id", Type: field.TypeInt, Nullable: true},
	}
	// CategoriesTable holds the schema information for the "categories" table.
	CategoriesTable = &schema.Table{
		Name:       "categories",
		Columns:    CategoriesColumns,
		PrimaryKey: []*schema.Column{CategoriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "categories_categories_children",
				Columns:    []*schema.Column{CategoriesColumns[3]},
				RefColumns: []*schema.Column{CategoriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ChaptersColumns holds the columns for the "chapters" table.
	ChaptersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "released_time", Type: field.TypeTime},
		{Name: "description", Type: field.TypeString},
		{Name: "sort", Type: field.TypeInt},
		{Name: "course_id", Type: field.TypeInt, Nullable: true},
	}
	// ChaptersTable holds the schema information for the "chapters" table.
	ChaptersTable = &schema.Table{
		Name:       "chapters",
		Columns:    ChaptersColumns,
		PrimaryKey: []*schema.Column{ChaptersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "chapters_courses_chapters",
				Columns:    []*schema.Column{ChaptersColumns[5]},
				RefColumns: []*schema.Column{CoursesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CoursesColumns holds the columns for the "courses" table.
	CoursesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "level", Type: field.TypeInt32},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "detail", Type: field.TypeString},
		{Name: "cover", Type: field.TypeString},
		{Name: "price", Type: field.TypeInt32},
		{Name: "tags", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeInt32},
		{Name: "score", Type: field.TypeInt32, Default: 0},
		{Name: "duration", Type: field.TypeInt32, Default: 0},
		{Name: "people", Type: field.TypeInt32, Default: 0},
		{Name: "category_id", Type: field.TypeInt, Nullable: true},
	}
	// CoursesTable holds the schema information for the "courses" table.
	CoursesTable = &schema.Table{
		Name:       "courses",
		Columns:    CoursesColumns,
		PrimaryKey: []*schema.Column{CoursesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "courses_categories_courses",
				Columns:    []*schema.Column{CoursesColumns[13]},
				RefColumns: []*schema.Column{CategoriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// LessonsColumns holds the columns for the "lessons" table.
	LessonsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "released_time", Type: field.TypeTime},
		{Name: "sort", Type: field.TypeInt},
		{Name: "type", Type: field.TypeInt},
		{Name: "storage_path", Type: field.TypeString},
		{Name: "source", Type: field.TypeString},
		{Name: "courseware", Type: field.TypeString},
		{Name: "is_free_preview", Type: field.TypeInt, Default: 2},
		{Name: "chapter_id", Type: field.TypeInt, Nullable: true},
	}
	// LessonsTable holds the schema information for the "lessons" table.
	LessonsTable = &schema.Table{
		Name:       "lessons",
		Columns:    LessonsColumns,
		PrimaryKey: []*schema.Column{LessonsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "lessons_chapters_lessons",
				Columns:    []*schema.Column{LessonsColumns[9]},
				RefColumns: []*schema.Column{ChaptersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// OrderInfosColumns holds the columns for the "order_infos" table.
	OrderInfosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "order_id", Type: field.TypeInt32, SchemaType: map[string]string{"mysql": "int"}},
		{Name: "product_id", Type: field.TypeInt32, SchemaType: map[string]string{"mysql": "int"}},
		{Name: "product_name", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(64)"}},
		{Name: "product_price", Type: field.TypeInt32, SchemaType: map[string]string{"mysql": "int"}},
		{Name: "product_describe", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"mysql": "text"}},
		{Name: "create_time", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "update_time", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
	}
	// OrderInfosTable holds the schema information for the "order_infos" table.
	OrderInfosTable = &schema.Table{
		Name:       "order_infos",
		Columns:    OrderInfosColumns,
		PrimaryKey: []*schema.Column{OrderInfosColumns[0]},
	}
	// OrdersColumns holds the columns for the "orders" table.
	OrdersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_id", Type: field.TypeInt32, SchemaType: map[string]string{"mysql": "int"}},
		{Name: "order_sn", Type: field.TypeString, SchemaType: map[string]string{"mysql": "varchar(64)"}},
		{Name: "pay_type", Type: field.TypeInt32, Default: 1, SchemaType: map[string]string{"mysql": "int"}},
		{Name: "pay_status", Type: field.TypeInt32, Nullable: true, Default: 1, SchemaType: map[string]string{"mysql": "int"}},
		{Name: "trade_price", Type: field.TypeInt32, SchemaType: map[string]string{"mysql": "int"}},
		{Name: "trade_no", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"mysql": "varchar(64)"}},
		{Name: "pay_time", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "create_time", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "update_time", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
	}
	// OrdersTable holds the schema information for the "orders" table.
	OrdersTable = &schema.Table{
		Name:       "orders",
		Columns:    OrdersColumns,
		PrimaryKey: []*schema.Column{OrdersColumns[0]},
	}
	// SettingsColumns holds the columns for the "settings" table.
	SettingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "detail", Type: field.TypeString},
		{Name: "cover", Type: field.TypeString},
	}
	// SettingsTable holds the schema information for the "settings" table.
	SettingsTable = &schema.Table{
		Name:       "settings",
		Columns:    SettingsColumns,
		PrimaryKey: []*schema.Column{SettingsColumns[0]},
	}
	// SlidersColumns holds the columns for the "sliders" table.
	SlidersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "content", Type: field.TypeString},
		{Name: "image_link", Type: field.TypeString},
		{Name: "create_at", Type: field.TypeTime},
		{Name: "update_at", Type: field.TypeTime},
		{Name: "is_valid", Type: field.TypeBool, Default: true},
		{Name: "priority", Type: field.TypeInt},
	}
	// SlidersTable holds the schema information for the "sliders" table.
	SlidersTable = &schema.Table{
		Name:       "sliders",
		Columns:    SlidersColumns,
		PrimaryKey: []*schema.Column{SlidersColumns[0]},
	}
	// TeachersColumns holds the columns for the "teachers" table.
	TeachersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "detail", Type: field.TypeString, Nullable: true, Size: 2147483647, SchemaType: map[string]string{"mysql": "text"}},
		{Name: "curriculum_vitae", Type: field.TypeString, Nullable: true, Size: 2147483647, SchemaType: map[string]string{"mysql": "text"}},
		{Name: "works", Type: field.TypeString, Nullable: true, Size: 2147483647, SchemaType: map[string]string{"mysql": "text"}},
		{Name: "skills", Type: field.TypeString, Nullable: true, Size: 2147483647, SchemaType: map[string]string{"mysql": "text"}},
		{Name: "name", Type: field.TypeString, SchemaType: map[string]string{"mysql": "varchar(64)"}},
		{Name: "level", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "int"}},
		{Name: "avator", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"mysql": "varchar(255)"}},
		{Name: "create_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "update_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
	}
	// TeachersTable holds the schema information for the "teachers" table.
	TeachersTable = &schema.Table{
		Name:       "teachers",
		Columns:    TeachersColumns,
		PrimaryKey: []*schema.Column{TeachersColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "username", Type: field.TypeString},
		{Name: "channel", Type: field.TypeString},
		{Name: "role_id", Type: field.TypeUint8},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// WalletsColumns holds the columns for the "wallets" table.
	WalletsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "gold_leaf", Type: field.TypeInt32, Nullable: true, Default: 0, SchemaType: map[string]string{"mysql": "int"}},
		{Name: "silver_leaf", Type: field.TypeInt32, Nullable: true, Default: 0, SchemaType: map[string]string{"mysql": "int"}},
		{Name: "frozen_gold_leaf", Type: field.TypeInt32, Nullable: true, Default: 0, SchemaType: map[string]string{"mysql": "int"}},
		{Name: "frozen_silver_leaf", Type: field.TypeInt32, Nullable: true, Default: 0, SchemaType: map[string]string{"mysql": "int"}},
		{Name: "user_id", Type: field.TypeInt32, Nullable: true, SchemaType: map[string]string{"mysql": "int"}},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
	}
	// WalletsTable holds the schema information for the "wallets" table.
	WalletsTable = &schema.Table{
		Name:       "wallets",
		Columns:    WalletsColumns,
		PrimaryKey: []*schema.Column{WalletsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AccountsTable,
		CategoriesTable,
		ChaptersTable,
		CoursesTable,
		LessonsTable,
		OrderInfosTable,
		OrdersTable,
		SettingsTable,
		SlidersTable,
		TeachersTable,
		UsersTable,
		WalletsTable,
	}
)

func init() {
	CategoriesTable.ForeignKeys[0].RefTable = CategoriesTable
	ChaptersTable.ForeignKeys[0].RefTable = CoursesTable
	CoursesTable.ForeignKeys[0].RefTable = CategoriesTable
	LessonsTable.ForeignKeys[0].RefTable = ChaptersTable
}
