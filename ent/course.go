// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"kubecit-service/ent/category"
	"kubecit-service/ent/course"
	"kubecit-service/ent/teacher"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Course is the model entity for the Course schema.
type Course struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Level holds the value of the "level" field.
	Level int32 `json:"level,omitempty"`
	// 修改时间
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Detail holds the value of the "detail" field.
	Detail string `json:"detail,omitempty"`
	// Cover holds the value of the "cover" field.
	Cover string `json:"cover,omitempty"`
	// Price holds the value of the "price" field.
	Price int32 `json:"price,omitempty"`
	// Tags holds the value of the "tags" field.
	Tags string `json:"tags,omitempty"`
	// 创建时间
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Status holds the value of the "status" field.
	Status int32 `json:"status,omitempty"`
	// CategoryID holds the value of the "category_id" field.
	CategoryID int `json:"category_id,omitempty"`
	// Score holds the value of the "score" field.
	Score int32 `json:"score,omitempty"`
	// Duration holds the value of the "duration" field.
	Duration int32 `json:"duration,omitempty"`
	// People holds the value of the "people" field.
	People int32 `json:"people,omitempty"`
	// TeacherID holds the value of the "teacher_id" field.
	TeacherID int `json:"teacher_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CourseQuery when eager-loading is set.
	Edges        CourseEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CourseEdges holds the relations/edges for other nodes in the graph.
type CourseEdges struct {
	// Owner holds the value of the owner edge.
	Owner *Category `json:"owner,omitempty"`
	// Chapters holds the value of the chapters edge.
	Chapters []*Chapter `json:"chapters,omitempty"`
	// Teacher holds the value of the teacher edge.
	Teacher *Teacher `json:"teacher,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CourseEdges) OwnerOrErr() (*Category, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: category.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// ChaptersOrErr returns the Chapters value or an error if the edge
// was not loaded in eager-loading.
func (e CourseEdges) ChaptersOrErr() ([]*Chapter, error) {
	if e.loadedTypes[1] {
		return e.Chapters, nil
	}
	return nil, &NotLoadedError{edge: "chapters"}
}

// TeacherOrErr returns the Teacher value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CourseEdges) TeacherOrErr() (*Teacher, error) {
	if e.loadedTypes[2] {
		if e.Teacher == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: teacher.Label}
		}
		return e.Teacher, nil
	}
	return nil, &NotLoadedError{edge: "teacher"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Course) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case course.FieldID, course.FieldLevel, course.FieldPrice, course.FieldStatus, course.FieldCategoryID, course.FieldScore, course.FieldDuration, course.FieldPeople, course.FieldTeacherID:
			values[i] = new(sql.NullInt64)
		case course.FieldName, course.FieldDetail, course.FieldCover, course.FieldTags:
			values[i] = new(sql.NullString)
		case course.FieldUpdatedAt, course.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Course fields.
func (c *Course) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case course.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case course.FieldLevel:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field level", values[i])
			} else if value.Valid {
				c.Level = int32(value.Int64)
			}
		case course.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case course.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case course.FieldDetail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field detail", values[i])
			} else if value.Valid {
				c.Detail = value.String
			}
		case course.FieldCover:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cover", values[i])
			} else if value.Valid {
				c.Cover = value.String
			}
		case course.FieldPrice:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value.Valid {
				c.Price = int32(value.Int64)
			}
		case course.FieldTags:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tags", values[i])
			} else if value.Valid {
				c.Tags = value.String
			}
		case course.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case course.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				c.Status = int32(value.Int64)
			}
		case course.FieldCategoryID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field category_id", values[i])
			} else if value.Valid {
				c.CategoryID = int(value.Int64)
			}
		case course.FieldScore:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field score", values[i])
			} else if value.Valid {
				c.Score = int32(value.Int64)
			}
		case course.FieldDuration:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field duration", values[i])
			} else if value.Valid {
				c.Duration = int32(value.Int64)
			}
		case course.FieldPeople:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field people", values[i])
			} else if value.Valid {
				c.People = int32(value.Int64)
			}
		case course.FieldTeacherID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field teacher_id", values[i])
			} else if value.Valid {
				c.TeacherID = int(value.Int64)
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Course.
// This includes values selected through modifiers, order, etc.
func (c *Course) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the Course entity.
func (c *Course) QueryOwner() *CategoryQuery {
	return NewCourseClient(c.config).QueryOwner(c)
}

// QueryChapters queries the "chapters" edge of the Course entity.
func (c *Course) QueryChapters() *ChapterQuery {
	return NewCourseClient(c.config).QueryChapters(c)
}

// QueryTeacher queries the "teacher" edge of the Course entity.
func (c *Course) QueryTeacher() *TeacherQuery {
	return NewCourseClient(c.config).QueryTeacher(c)
}

// Update returns a builder for updating this Course.
// Note that you need to call Course.Unwrap() before calling this method if this Course
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Course) Update() *CourseUpdateOne {
	return NewCourseClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Course entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Course) Unwrap() *Course {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Course is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Course) String() string {
	var builder strings.Builder
	builder.WriteString("Course(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("level=")
	builder.WriteString(fmt.Sprintf("%v", c.Level))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("detail=")
	builder.WriteString(c.Detail)
	builder.WriteString(", ")
	builder.WriteString("cover=")
	builder.WriteString(c.Cover)
	builder.WriteString(", ")
	builder.WriteString("price=")
	builder.WriteString(fmt.Sprintf("%v", c.Price))
	builder.WriteString(", ")
	builder.WriteString("tags=")
	builder.WriteString(c.Tags)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", c.Status))
	builder.WriteString(", ")
	builder.WriteString("category_id=")
	builder.WriteString(fmt.Sprintf("%v", c.CategoryID))
	builder.WriteString(", ")
	builder.WriteString("score=")
	builder.WriteString(fmt.Sprintf("%v", c.Score))
	builder.WriteString(", ")
	builder.WriteString("duration=")
	builder.WriteString(fmt.Sprintf("%v", c.Duration))
	builder.WriteString(", ")
	builder.WriteString("people=")
	builder.WriteString(fmt.Sprintf("%v", c.People))
	builder.WriteString(", ")
	builder.WriteString("teacher_id=")
	builder.WriteString(fmt.Sprintf("%v", c.TeacherID))
	builder.WriteByte(')')
	return builder.String()
}

// Courses is a parsable slice of Course.
type Courses []*Course
