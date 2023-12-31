// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"kubecit-service/ent/user"
	"kubecit-service/ent/vipinfo"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// VipInfo is the model entity for the VipInfo schema.
type VipInfo struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// 会员类型
	VipType int8 `json:"vip_type,omitempty"`
	// 会员权益生效时间
	StartAt time.Time `json:"start_at,omitempty"`
	// 会员权益失效时间
	ExpireAt time.Time `json:"expire_at,omitempty"`
	// 用户id
	UserID int `json:"user_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the VipInfoQuery when eager-loading is set.
	Edges        VipInfoEdges `json:"edges"`
	selectValues sql.SelectValues
}

// VipInfoEdges holds the relations/edges for other nodes in the graph.
type VipInfoEdges struct {
	// UserInfo holds the value of the user_info edge.
	UserInfo *User `json:"user_info,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserInfoOrErr returns the UserInfo value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e VipInfoEdges) UserInfoOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.UserInfo == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.UserInfo, nil
	}
	return nil, &NotLoadedError{edge: "user_info"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*VipInfo) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case vipinfo.FieldID, vipinfo.FieldVipType, vipinfo.FieldUserID:
			values[i] = new(sql.NullInt64)
		case vipinfo.FieldStartAt, vipinfo.FieldExpireAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the VipInfo fields.
func (vi *VipInfo) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case vipinfo.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			vi.ID = int(value.Int64)
		case vipinfo.FieldVipType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field vip_type", values[i])
			} else if value.Valid {
				vi.VipType = int8(value.Int64)
			}
		case vipinfo.FieldStartAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field start_at", values[i])
			} else if value.Valid {
				vi.StartAt = value.Time
			}
		case vipinfo.FieldExpireAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expire_at", values[i])
			} else if value.Valid {
				vi.ExpireAt = value.Time
			}
		case vipinfo.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				vi.UserID = int(value.Int64)
			}
		default:
			vi.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the VipInfo.
// This includes values selected through modifiers, order, etc.
func (vi *VipInfo) Value(name string) (ent.Value, error) {
	return vi.selectValues.Get(name)
}

// QueryUserInfo queries the "user_info" edge of the VipInfo entity.
func (vi *VipInfo) QueryUserInfo() *UserQuery {
	return NewVipInfoClient(vi.config).QueryUserInfo(vi)
}

// Update returns a builder for updating this VipInfo.
// Note that you need to call VipInfo.Unwrap() before calling this method if this VipInfo
// was returned from a transaction, and the transaction was committed or rolled back.
func (vi *VipInfo) Update() *VipInfoUpdateOne {
	return NewVipInfoClient(vi.config).UpdateOne(vi)
}

// Unwrap unwraps the VipInfo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (vi *VipInfo) Unwrap() *VipInfo {
	_tx, ok := vi.config.driver.(*txDriver)
	if !ok {
		panic("ent: VipInfo is not a transactional entity")
	}
	vi.config.driver = _tx.drv
	return vi
}

// String implements the fmt.Stringer.
func (vi *VipInfo) String() string {
	var builder strings.Builder
	builder.WriteString("VipInfo(")
	builder.WriteString(fmt.Sprintf("id=%v, ", vi.ID))
	builder.WriteString("vip_type=")
	builder.WriteString(fmt.Sprintf("%v", vi.VipType))
	builder.WriteString(", ")
	builder.WriteString("start_at=")
	builder.WriteString(vi.StartAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("expire_at=")
	builder.WriteString(vi.ExpireAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", vi.UserID))
	builder.WriteByte(')')
	return builder.String()
}

// VipInfos is a parsable slice of VipInfo.
type VipInfos []*VipInfo
