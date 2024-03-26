// Code generated by ent, DO NOT EDIT.

package msg

import (
	"IM/internel/types/enums"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the msg type in the database.
	Label = "msg"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldSendAt holds the string denoting the send_at field in the database.
	FieldSendAt = "send_at"
	// FieldSessionType holds the string denoting the session_type field in the database.
	FieldSessionType = "session_type"
	// FieldSendID holds the string denoting the send_id field in the database.
	FieldSendID = "send_id"
	// FieldReceiveID holds the string denoting the receive_id field in the database.
	FieldReceiveID = "receive_id"
	// FieldContentType holds the string denoting the content_type field in the database.
	FieldContentType = "content_type"
	// FieldSeq holds the string denoting the seq field in the database.
	FieldSeq = "seq"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldTextElem holds the string denoting the text_elem field in the database.
	FieldTextElem = "text_elem"
	// EdgeSendUser holds the string denoting the send_user edge name in mutations.
	EdgeSendUser = "send_user"
	// EdgeReceiveUser holds the string denoting the receive_user edge name in mutations.
	EdgeReceiveUser = "receive_user"
	// Table holds the table name of the msg in the database.
	Table = "t_message"
	// SendUserTable is the table that holds the send_user relation/edge.
	SendUserTable = "t_message"
	// SendUserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	SendUserInverseTable = "t_user"
	// SendUserColumn is the table column denoting the send_user relation/edge.
	SendUserColumn = "send_id"
	// ReceiveUserTable is the table that holds the receive_user relation/edge.
	ReceiveUserTable = "t_message"
	// ReceiveUserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	ReceiveUserInverseTable = "t_user"
	// ReceiveUserColumn is the table column denoting the receive_user relation/edge.
	ReceiveUserColumn = "receive_id"
)

// Columns holds all SQL columns for msg fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldSendAt,
	FieldSessionType,
	FieldSendID,
	FieldReceiveID,
	FieldContentType,
	FieldSeq,
	FieldStatus,
	FieldTextElem,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultDeletedAt holds the default value on creation for the "deleted_at" field.
	DefaultDeletedAt time.Time
	// DefaultSendAt holds the default value on creation for the "send_at" field.
	DefaultSendAt time.Time
	// DefaultSeq holds the default value on creation for the "seq" field.
	DefaultSeq int32
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int64
)

const DefaultSessionType enums.SessionType = "unknown"

// SessionTypeValidator is a validator for the "session_type" field enum values. It is called by the builders before save.
func SessionTypeValidator(st enums.SessionType) error {
	switch st {
	case "unknown", "single", "group":
		return nil
	default:
		return fmt.Errorf("msg: invalid enum value for session_type field: %q", st)
	}
}

const DefaultContentType enums.MessageType = "unknown"

// ContentTypeValidator is a validator for the "content_type" field enum values. It is called by the builders before save.
func ContentTypeValidator(ct enums.MessageType) error {
	switch ct {
	case "unknown", "textMessage":
		return nil
	default:
		return fmt.Errorf("msg: invalid enum value for content_type field: %q", ct)
	}
}

const DefaultStatus enums.MessageStatus = "unknown"

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s enums.MessageStatus) error {
	switch s {
	case "unknown", "sending", "succeed", "failed":
		return nil
	default:
		return fmt.Errorf("msg: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the Msg queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// BySendAt orders the results by the send_at field.
func BySendAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSendAt, opts...).ToFunc()
}

// BySessionType orders the results by the session_type field.
func BySessionType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSessionType, opts...).ToFunc()
}

// BySendID orders the results by the send_id field.
func BySendID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSendID, opts...).ToFunc()
}

// ByReceiveID orders the results by the receive_id field.
func ByReceiveID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReceiveID, opts...).ToFunc()
}

// ByContentType orders the results by the content_type field.
func ByContentType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContentType, opts...).ToFunc()
}

// BySeq orders the results by the seq field.
func BySeq(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSeq, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByTextElem orders the results by the text_elem field.
func ByTextElem(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTextElem, opts...).ToFunc()
}

// BySendUserField orders the results by send_user field.
func BySendUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSendUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByReceiveUserField orders the results by receive_user field.
func ByReceiveUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newReceiveUserStep(), sql.OrderByField(field, opts...))
	}
}
func newSendUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SendUserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, SendUserTable, SendUserColumn),
	)
}
func newReceiveUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ReceiveUserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ReceiveUserTable, ReceiveUserColumn),
	)
}