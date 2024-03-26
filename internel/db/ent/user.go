// Code generated by ent, DO NOT EDIT.

package ent

import (
	"IM/internel/db/ent/user"
	"IM/internel/types/enums"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// User is the model entity for the User schema.
type User struct {
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
	// 用户昵称
	NickName string `json:"nick_name"`
	// 邮箱
	Email string `json:"email"`
	// 密码
	Password string `json:"password"`
	// 用户状态
	Status enums.UserStatus `json:"status"`
	// 最后在线时间
	LastOnlineAt time.Time `json:"last_online_at"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// SendMsg holds the value of the send_msg edge.
	SendMsg []*Msg `json:"send_msg,omitempty"`
	// ReceiveMsg holds the value of the receive_msg edge.
	ReceiveMsg []*Msg `json:"receive_msg,omitempty"`
	// OwnerUserFriend holds the value of the owner_user_friend edge.
	OwnerUserFriend []*Friend `json:"owner_user_friend,omitempty"`
	// FriendUserFriend holds the value of the friend_user_friend edge.
	FriendUserFriend []*Friend `json:"friend_user_friend,omitempty"`
	// UserGroup holds the value of the user_group edge.
	UserGroup []*Group `json:"user_group,omitempty"`
	// UserGroupMember holds the value of the user_group_member edge.
	UserGroupMember []*GroupMember `json:"user_group_member,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// SendMsgOrErr returns the SendMsg value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) SendMsgOrErr() ([]*Msg, error) {
	if e.loadedTypes[0] {
		return e.SendMsg, nil
	}
	return nil, &NotLoadedError{edge: "send_msg"}
}

// ReceiveMsgOrErr returns the ReceiveMsg value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ReceiveMsgOrErr() ([]*Msg, error) {
	if e.loadedTypes[1] {
		return e.ReceiveMsg, nil
	}
	return nil, &NotLoadedError{edge: "receive_msg"}
}

// OwnerUserFriendOrErr returns the OwnerUserFriend value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) OwnerUserFriendOrErr() ([]*Friend, error) {
	if e.loadedTypes[2] {
		return e.OwnerUserFriend, nil
	}
	return nil, &NotLoadedError{edge: "owner_user_friend"}
}

// FriendUserFriendOrErr returns the FriendUserFriend value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) FriendUserFriendOrErr() ([]*Friend, error) {
	if e.loadedTypes[3] {
		return e.FriendUserFriend, nil
	}
	return nil, &NotLoadedError{edge: "friend_user_friend"}
}

// UserGroupOrErr returns the UserGroup value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) UserGroupOrErr() ([]*Group, error) {
	if e.loadedTypes[4] {
		return e.UserGroup, nil
	}
	return nil, &NotLoadedError{edge: "user_group"}
}

// UserGroupMemberOrErr returns the UserGroupMember value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) UserGroupMemberOrErr() ([]*GroupMember, error) {
	if e.loadedTypes[5] {
		return e.UserGroupMember, nil
	}
	return nil, &NotLoadedError{edge: "user_group_member"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			values[i] = new(sql.NullInt64)
		case user.FieldNickName, user.FieldEmail, user.FieldPassword, user.FieldStatus:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt, user.FieldUpdatedAt, user.FieldDeletedAt, user.FieldLastOnlineAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int64(value.Int64)
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		case user.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				u.DeletedAt = value.Time
			}
		case user.FieldNickName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nick_name", values[i])
			} else if value.Valid {
				u.NickName = value.String
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case user.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				u.Status = enums.UserStatus(value.String)
			}
		case user.FieldLastOnlineAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_online_at", values[i])
			} else if value.Valid {
				u.LastOnlineAt = value.Time
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QuerySendMsg queries the "send_msg" edge of the User entity.
func (u *User) QuerySendMsg() *MsgQuery {
	return NewUserClient(u.config).QuerySendMsg(u)
}

// QueryReceiveMsg queries the "receive_msg" edge of the User entity.
func (u *User) QueryReceiveMsg() *MsgQuery {
	return NewUserClient(u.config).QueryReceiveMsg(u)
}

// QueryOwnerUserFriend queries the "owner_user_friend" edge of the User entity.
func (u *User) QueryOwnerUserFriend() *FriendQuery {
	return NewUserClient(u.config).QueryOwnerUserFriend(u)
}

// QueryFriendUserFriend queries the "friend_user_friend" edge of the User entity.
func (u *User) QueryFriendUserFriend() *FriendQuery {
	return NewUserClient(u.config).QueryFriendUserFriend(u)
}

// QueryUserGroup queries the "user_group" edge of the User entity.
func (u *User) QueryUserGroup() *GroupQuery {
	return NewUserClient(u.config).QueryUserGroup(u)
}

// QueryUserGroupMember queries the "user_group_member" edge of the User entity.
func (u *User) QueryUserGroupMember() *GroupMemberQuery {
	return NewUserClient(u.config).QueryUserGroupMember(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(u.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("nick_name=")
	builder.WriteString(u.NickName)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	builder.WriteString("password=")
	builder.WriteString(u.Password)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", u.Status))
	builder.WriteString(", ")
	builder.WriteString("last_online_at=")
	builder.WriteString(u.LastOnlineAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User
