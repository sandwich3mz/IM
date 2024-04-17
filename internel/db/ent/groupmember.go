// Code generated by ent, DO NOT EDIT.

package ent

import (
	"IM/internel/db/ent/group"
	"IM/internel/db/ent/groupmember"
	"IM/internel/db/ent/user"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// GroupMember is the model entity for the GroupMember schema.
type GroupMember struct {
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
	// 群 ID
	GroupID int64 `json:"group_id,string"`
	// 用户 ID
	UserID int64 `json:"user_id,string"`
	// 群成员昵称
	Nickname int8 `json:"nickname"`
	// 头像
	FaceURL string `json:"face_url"`
	// 群员数量
	Count int64 `json:"count"`
	// 备注
	Remark string `json:"remark"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GroupMemberQuery when eager-loading is set.
	Edges        GroupMemberEdges `json:"edges"`
	selectValues sql.SelectValues
}

// GroupMemberEdges holds the relations/edges for other nodes in the graph.
type GroupMemberEdges struct {
	// GroupUser holds the value of the group_user edge.
	GroupUser *User `json:"group_user,omitempty"`
	// MemberGroup holds the value of the member_group edge.
	MemberGroup *Group `json:"member_group,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// GroupUserOrErr returns the GroupUser value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GroupMemberEdges) GroupUserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.GroupUser == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.GroupUser, nil
	}
	return nil, &NotLoadedError{edge: "group_user"}
}

// MemberGroupOrErr returns the MemberGroup value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GroupMemberEdges) MemberGroupOrErr() (*Group, error) {
	if e.loadedTypes[1] {
		if e.MemberGroup == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: group.Label}
		}
		return e.MemberGroup, nil
	}
	return nil, &NotLoadedError{edge: "member_group"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GroupMember) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case groupmember.FieldID, groupmember.FieldGroupID, groupmember.FieldUserID, groupmember.FieldNickname, groupmember.FieldCount:
			values[i] = new(sql.NullInt64)
		case groupmember.FieldFaceURL, groupmember.FieldRemark:
			values[i] = new(sql.NullString)
		case groupmember.FieldCreatedAt, groupmember.FieldUpdatedAt, groupmember.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GroupMember fields.
func (gm *GroupMember) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case groupmember.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			gm.ID = int64(value.Int64)
		case groupmember.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				gm.CreatedAt = value.Time
			}
		case groupmember.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				gm.UpdatedAt = value.Time
			}
		case groupmember.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				gm.DeletedAt = value.Time
			}
		case groupmember.FieldGroupID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field group_id", values[i])
			} else if value.Valid {
				gm.GroupID = value.Int64
			}
		case groupmember.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				gm.UserID = value.Int64
			}
		case groupmember.FieldNickname:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field nickname", values[i])
			} else if value.Valid {
				gm.Nickname = int8(value.Int64)
			}
		case groupmember.FieldFaceURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field face_url", values[i])
			} else if value.Valid {
				gm.FaceURL = value.String
			}
		case groupmember.FieldCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field count", values[i])
			} else if value.Valid {
				gm.Count = value.Int64
			}
		case groupmember.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				gm.Remark = value.String
			}
		default:
			gm.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the GroupMember.
// This includes values selected through modifiers, order, etc.
func (gm *GroupMember) Value(name string) (ent.Value, error) {
	return gm.selectValues.Get(name)
}

// QueryGroupUser queries the "group_user" edge of the GroupMember entity.
func (gm *GroupMember) QueryGroupUser() *UserQuery {
	return NewGroupMemberClient(gm.config).QueryGroupUser(gm)
}

// QueryMemberGroup queries the "member_group" edge of the GroupMember entity.
func (gm *GroupMember) QueryMemberGroup() *GroupQuery {
	return NewGroupMemberClient(gm.config).QueryMemberGroup(gm)
}

// Update returns a builder for updating this GroupMember.
// Note that you need to call GroupMember.Unwrap() before calling this method if this GroupMember
// was returned from a transaction, and the transaction was committed or rolled back.
func (gm *GroupMember) Update() *GroupMemberUpdateOne {
	return NewGroupMemberClient(gm.config).UpdateOne(gm)
}

// Unwrap unwraps the GroupMember entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gm *GroupMember) Unwrap() *GroupMember {
	_tx, ok := gm.config.driver.(*txDriver)
	if !ok {
		panic("ent: GroupMember is not a transactional entity")
	}
	gm.config.driver = _tx.drv
	return gm
}

// String implements the fmt.Stringer.
func (gm *GroupMember) String() string {
	var builder strings.Builder
	builder.WriteString("GroupMember(")
	builder.WriteString(fmt.Sprintf("id=%v, ", gm.ID))
	builder.WriteString("created_at=")
	builder.WriteString(gm.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(gm.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(gm.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("group_id=")
	builder.WriteString(fmt.Sprintf("%v", gm.GroupID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", gm.UserID))
	builder.WriteString(", ")
	builder.WriteString("nickname=")
	builder.WriteString(fmt.Sprintf("%v", gm.Nickname))
	builder.WriteString(", ")
	builder.WriteString("face_url=")
	builder.WriteString(gm.FaceURL)
	builder.WriteString(", ")
	builder.WriteString("count=")
	builder.WriteString(fmt.Sprintf("%v", gm.Count))
	builder.WriteString(", ")
	builder.WriteString("remark=")
	builder.WriteString(gm.Remark)
	builder.WriteByte(')')
	return builder.String()
}

// GroupMembers is a parsable slice of GroupMember.
type GroupMembers []*GroupMember
