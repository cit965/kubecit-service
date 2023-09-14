// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"kubecit-service/ent/orderinfos"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// OrderInfos is the model entity for the OrderInfos schema.
type OrderInfos struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// 	订单id
	OrderID int32 `json:"order_id,omitempty"`
	// 课程id
	CourseID int32 `json:"course_id,omitempty"`
	// 课程名称
	CourseName string `json:"course_name,omitempty"`
	// 课程价格(单位分)
	CoursePrice int32 `json:"course_price,omitempty"`
	// 课程描述
	CourseDescribe string `json:"course_describe,omitempty"`
	// 创建时间
	CreateTime time.Time `json:"create_time,omitempty"`
	// 更新时间
	UpdateTime   time.Time `json:"update_time,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OrderInfos) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case orderinfos.FieldID, orderinfos.FieldOrderID, orderinfos.FieldCourseID, orderinfos.FieldCoursePrice:
			values[i] = new(sql.NullInt64)
		case orderinfos.FieldCourseName, orderinfos.FieldCourseDescribe:
			values[i] = new(sql.NullString)
		case orderinfos.FieldCreateTime, orderinfos.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OrderInfos fields.
func (oi *OrderInfos) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case orderinfos.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			oi.ID = int(value.Int64)
		case orderinfos.FieldOrderID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order_id", values[i])
			} else if value.Valid {
				oi.OrderID = int32(value.Int64)
			}
		case orderinfos.FieldCourseID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field course_id", values[i])
			} else if value.Valid {
				oi.CourseID = int32(value.Int64)
			}
		case orderinfos.FieldCourseName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field course_name", values[i])
			} else if value.Valid {
				oi.CourseName = value.String
			}
		case orderinfos.FieldCoursePrice:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field course_price", values[i])
			} else if value.Valid {
				oi.CoursePrice = int32(value.Int64)
			}
		case orderinfos.FieldCourseDescribe:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field course_describe", values[i])
			} else if value.Valid {
				oi.CourseDescribe = value.String
			}
		case orderinfos.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				oi.CreateTime = value.Time
			}
		case orderinfos.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				oi.UpdateTime = value.Time
			}
		default:
			oi.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the OrderInfos.
// This includes values selected through modifiers, order, etc.
func (oi *OrderInfos) Value(name string) (ent.Value, error) {
	return oi.selectValues.Get(name)
}

// Update returns a builder for updating this OrderInfos.
// Note that you need to call OrderInfos.Unwrap() before calling this method if this OrderInfos
// was returned from a transaction, and the transaction was committed or rolled back.
func (oi *OrderInfos) Update() *OrderInfosUpdateOne {
	return NewOrderInfosClient(oi.config).UpdateOne(oi)
}

// Unwrap unwraps the OrderInfos entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (oi *OrderInfos) Unwrap() *OrderInfos {
	_tx, ok := oi.config.driver.(*txDriver)
	if !ok {
		panic("ent: OrderInfos is not a transactional entity")
	}
	oi.config.driver = _tx.drv
	return oi
}

// String implements the fmt.Stringer.
func (oi *OrderInfos) String() string {
	var builder strings.Builder
	builder.WriteString("OrderInfos(")
	builder.WriteString(fmt.Sprintf("id=%v, ", oi.ID))
	builder.WriteString("order_id=")
	builder.WriteString(fmt.Sprintf("%v", oi.OrderID))
	builder.WriteString(", ")
	builder.WriteString("course_id=")
	builder.WriteString(fmt.Sprintf("%v", oi.CourseID))
	builder.WriteString(", ")
	builder.WriteString("course_name=")
	builder.WriteString(oi.CourseName)
	builder.WriteString(", ")
	builder.WriteString("course_price=")
	builder.WriteString(fmt.Sprintf("%v", oi.CoursePrice))
	builder.WriteString(", ")
	builder.WriteString("course_describe=")
	builder.WriteString(oi.CourseDescribe)
	builder.WriteString(", ")
	builder.WriteString("create_time=")
	builder.WriteString(oi.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(oi.UpdateTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// OrderInfosSlice is a parsable slice of OrderInfos.
type OrderInfosSlice []*OrderInfos
