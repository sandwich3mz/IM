// Code generated by ent, DO NOT EDIT.

package ent

import (
	"IM/internel/db/ent/group"
	"IM/internel/db/ent/user"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Group is the model entity for the Group schema.
type Group struct {
	config `json:"-"`
	// ID of the ent.
	// 19 位雪花 ID
	ID int64 `json:"id,string"`
	// 创建时刻，带时区
	CreatedAt time.Time `json:"created_at"`
	// 更新时刻，带时区
	UpdatedAt time.Time `json:"updated_at"`
	// 软删除时刻，带时区
	DeletedAt time.Time `json:"deleted_at"`
	// 群聊名称
	GroupName string `json:"group_name"`
	// 群公告
	Notification string `json:"notification"`
	// 用户关系
	Relationship int8 `json:"relationship"`
	// 群主 ID
	OwnerUserID int64 `json:"owner_user_id,string"`
	// 头像
	FaceURL string `json:"face_url"`
	// 群员数量
	Count int64 `json:"count"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GroupQuery when eager-loading is set.
	Edges        GroupEdges `json:"edges"`
	selectValues sql.SelectValues
}

// GroupEdges holds the relations/edges for other nodes in the graph.
type GroupEdges struct {
	// OwnerUser holds the value of the owner_user edge.
	OwnerUser *User `json:"owner_user,omitempty"`
	// GroupMember holds the value of the group_member edge.
	GroupMember []*GroupMember `json:"group_member,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// OwnerUserOrErr returns the OwnerUser value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GroupEdges) OwnerUserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.OwnerUser == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.OwnerUser, nil
	}
	return nil, &NotLoadedError{edge: "owner_user"}
}

// GroupMemberOrErr returns the GroupMember value or an error if the edge
// was not loaded in eager-loading.
func (e GroupEdges) GroupMemberOrErr() ([]*GroupMember, error) {
	if e.loadedTypes[1] {
		return e.GroupMember, nil
	}
	return nil, &NotLoadedError{edge: "group_member"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Group) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case group.FieldID, group.FieldRelationship, group.FieldOwnerUserID, group.FieldCount:
			values[i] = new(sql.NullInt64)
		case group.FieldGroupName, group.FieldNotification, group.FieldFaceURL:
			values[i] = new(sql.NullString)
		case group.FieldCreatedAt, group.FieldUpdatedAt, group.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Group fields.
func (gr *Group) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case group.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			gr.ID = int64(value.Int64)
		case group.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				gr.CreatedAt = value.Time
			}
		case group.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				gr.UpdatedAt = value.Time
			}
		case group.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				gr.DeletedAt = value.Time
			}
		case group.FieldGroupName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field group_name", values[i])
			} else if value.Valid {
				gr.GroupName = value.String
			}
		case group.FieldNotification:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field notification", values[i])
			} else if value.Valid {
				gr.Notification = value.String
			}
		case group.FieldRelationship:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field relationship", values[i])
			} else if value.Valid {
				gr.Relationship = int8(value.Int64)
			}
		case group.FieldOwnerUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field owner_user_id", values[i])
			} else if value.Valid {
				gr.OwnerUserID = value.Int64
			}
		case group.FieldFaceURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field face_url", values[i])
			} else if value.Valid {
				gr.FaceURL = value.String
			}
		case group.FieldCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field count", values[i])
			} else if value.Valid {
				gr.Count = value.Int64
			}
		default:
			gr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Group.
// This includes values selected through modifiers, order, etc.
func (gr *Group) Value(name string) (ent.Value, error) {
	return gr.selectValues.Get(name)
}

// QueryOwnerUser queries the "owner_user" edge of the Group entity.
func (gr *Group) QueryOwnerUser() *UserQuery {
	return NewGroupClient(gr.config).QueryOwnerUser(gr)
}

// QueryGroupMember queries the "group_member" edge of the Group entity.
func (gr *Group) QueryGroupMember() *GroupMemberQuery {
	return NewGroupClient(gr.config).QueryGroupMember(gr)
}

// Update returns a builder for updating this Group.
// Note that you need to call Group.Unwrap() before calling this method if this Group
// was returned from a transaction, and the transaction was committed or rolled back.
func (gr *Group) Update() *GroupUpdateOne {
	return NewGroupClient(gr.config).UpdateOne(gr)
}

// Unwrap unwraps the Group entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gr *Group) Unwrap() *Group {
	_tx, ok := gr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Group is not a transactional entity")
	}
	gr.config.driver = _tx.drv
	return gr
}

// String implements the fmt.Stringer.
func (gr *Group) String() string {
	var builder strings.Builder
	builder.WriteString("Group(")
	builder.WriteString(fmt.Sprintf("id=%v, ", gr.ID))
	builder.WriteString("created_at=")
	builder.WriteString(gr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(gr.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(gr.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("group_name=")
	builder.WriteString(gr.GroupName)
	builder.WriteString(", ")
	builder.WriteString("notification=")
	builder.WriteString(gr.Notification)
	builder.WriteString(", ")
	builder.WriteString("relationship=")
	builder.WriteString(fmt.Sprintf("%v", gr.Relationship))
	builder.WriteString(", ")
	builder.WriteString("owner_user_id=")
	builder.WriteString(fmt.Sprintf("%v", gr.OwnerUserID))
	builder.WriteString(", ")
	builder.WriteString("face_url=")
	builder.WriteString(gr.FaceURL)
	builder.WriteString(", ")
	builder.WriteString("count=")
	builder.WriteString(fmt.Sprintf("%v", gr.Count))
	builder.WriteByte(')')
	return builder.String()
}

// Groups is a parsable slice of Group.
type Groups []*Group
