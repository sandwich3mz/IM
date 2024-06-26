// Code generated by ent, DO NOT EDIT.

package friend

import (
	"IM/internel/db/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Friend {
	return predicate.Friend(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Friend {
	return predicate.Friend(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Friend {
	return predicate.Friend(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Friend {
	return predicate.Friend(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Friend {
	return predicate.Friend(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Friend {
	return predicate.Friend(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Friend {
	return predicate.Friend(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldDeletedAt, v))
}

// OwnerUserID applies equality check predicate on the "owner_user_id" field. It's identical to OwnerUserIDEQ.
func OwnerUserID(v int64) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldOwnerUserID, v))
}

// FriendUserID applies equality check predicate on the "friend_user_id" field. It's identical to FriendUserIDEQ.
func FriendUserID(v int64) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldFriendUserID, v))
}

// Relationship applies equality check predicate on the "relationship" field. It's identical to RelationshipEQ.
func Relationship(v int8) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldRelationship, v))
}

// Remark applies equality check predicate on the "remark" field. It's identical to RemarkEQ.
func Remark(v string) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldRemark, v))
}

// GroupID applies equality check predicate on the "group_id" field. It's identical to GroupIDEQ.
func GroupID(v int64) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldGroupID, v))
}

// LastTalkAt applies equality check predicate on the "last_talk_at" field. It's identical to LastTalkAtEQ.
func LastTalkAt(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldLastTalkAt, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldLTE(FieldDeletedAt, v))
}

// OwnerUserIDEQ applies the EQ predicate on the "owner_user_id" field.
func OwnerUserIDEQ(v int64) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldOwnerUserID, v))
}

// OwnerUserIDNEQ applies the NEQ predicate on the "owner_user_id" field.
func OwnerUserIDNEQ(v int64) predicate.Friend {
	return predicate.Friend(sql.FieldNEQ(FieldOwnerUserID, v))
}

// OwnerUserIDIn applies the In predicate on the "owner_user_id" field.
func OwnerUserIDIn(vs ...int64) predicate.Friend {
	return predicate.Friend(sql.FieldIn(FieldOwnerUserID, vs...))
}

// OwnerUserIDNotIn applies the NotIn predicate on the "owner_user_id" field.
func OwnerUserIDNotIn(vs ...int64) predicate.Friend {
	return predicate.Friend(sql.FieldNotIn(FieldOwnerUserID, vs...))
}

// FriendUserIDEQ applies the EQ predicate on the "friend_user_id" field.
func FriendUserIDEQ(v int64) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldFriendUserID, v))
}

// FriendUserIDNEQ applies the NEQ predicate on the "friend_user_id" field.
func FriendUserIDNEQ(v int64) predicate.Friend {
	return predicate.Friend(sql.FieldNEQ(FieldFriendUserID, v))
}

// FriendUserIDIn applies the In predicate on the "friend_user_id" field.
func FriendUserIDIn(vs ...int64) predicate.Friend {
	return predicate.Friend(sql.FieldIn(FieldFriendUserID, vs...))
}

// FriendUserIDNotIn applies the NotIn predicate on the "friend_user_id" field.
func FriendUserIDNotIn(vs ...int64) predicate.Friend {
	return predicate.Friend(sql.FieldNotIn(FieldFriendUserID, vs...))
}

// RelationshipEQ applies the EQ predicate on the "relationship" field.
func RelationshipEQ(v int8) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldRelationship, v))
}

// RelationshipNEQ applies the NEQ predicate on the "relationship" field.
func RelationshipNEQ(v int8) predicate.Friend {
	return predicate.Friend(sql.FieldNEQ(FieldRelationship, v))
}

// RelationshipIn applies the In predicate on the "relationship" field.
func RelationshipIn(vs ...int8) predicate.Friend {
	return predicate.Friend(sql.FieldIn(FieldRelationship, vs...))
}

// RelationshipNotIn applies the NotIn predicate on the "relationship" field.
func RelationshipNotIn(vs ...int8) predicate.Friend {
	return predicate.Friend(sql.FieldNotIn(FieldRelationship, vs...))
}

// RelationshipGT applies the GT predicate on the "relationship" field.
func RelationshipGT(v int8) predicate.Friend {
	return predicate.Friend(sql.FieldGT(FieldRelationship, v))
}

// RelationshipGTE applies the GTE predicate on the "relationship" field.
func RelationshipGTE(v int8) predicate.Friend {
	return predicate.Friend(sql.FieldGTE(FieldRelationship, v))
}

// RelationshipLT applies the LT predicate on the "relationship" field.
func RelationshipLT(v int8) predicate.Friend {
	return predicate.Friend(sql.FieldLT(FieldRelationship, v))
}

// RelationshipLTE applies the LTE predicate on the "relationship" field.
func RelationshipLTE(v int8) predicate.Friend {
	return predicate.Friend(sql.FieldLTE(FieldRelationship, v))
}

// RemarkEQ applies the EQ predicate on the "remark" field.
func RemarkEQ(v string) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldRemark, v))
}

// RemarkNEQ applies the NEQ predicate on the "remark" field.
func RemarkNEQ(v string) predicate.Friend {
	return predicate.Friend(sql.FieldNEQ(FieldRemark, v))
}

// RemarkIn applies the In predicate on the "remark" field.
func RemarkIn(vs ...string) predicate.Friend {
	return predicate.Friend(sql.FieldIn(FieldRemark, vs...))
}

// RemarkNotIn applies the NotIn predicate on the "remark" field.
func RemarkNotIn(vs ...string) predicate.Friend {
	return predicate.Friend(sql.FieldNotIn(FieldRemark, vs...))
}

// RemarkGT applies the GT predicate on the "remark" field.
func RemarkGT(v string) predicate.Friend {
	return predicate.Friend(sql.FieldGT(FieldRemark, v))
}

// RemarkGTE applies the GTE predicate on the "remark" field.
func RemarkGTE(v string) predicate.Friend {
	return predicate.Friend(sql.FieldGTE(FieldRemark, v))
}

// RemarkLT applies the LT predicate on the "remark" field.
func RemarkLT(v string) predicate.Friend {
	return predicate.Friend(sql.FieldLT(FieldRemark, v))
}

// RemarkLTE applies the LTE predicate on the "remark" field.
func RemarkLTE(v string) predicate.Friend {
	return predicate.Friend(sql.FieldLTE(FieldRemark, v))
}

// RemarkContains applies the Contains predicate on the "remark" field.
func RemarkContains(v string) predicate.Friend {
	return predicate.Friend(sql.FieldContains(FieldRemark, v))
}

// RemarkHasPrefix applies the HasPrefix predicate on the "remark" field.
func RemarkHasPrefix(v string) predicate.Friend {
	return predicate.Friend(sql.FieldHasPrefix(FieldRemark, v))
}

// RemarkHasSuffix applies the HasSuffix predicate on the "remark" field.
func RemarkHasSuffix(v string) predicate.Friend {
	return predicate.Friend(sql.FieldHasSuffix(FieldRemark, v))
}

// RemarkEqualFold applies the EqualFold predicate on the "remark" field.
func RemarkEqualFold(v string) predicate.Friend {
	return predicate.Friend(sql.FieldEqualFold(FieldRemark, v))
}

// RemarkContainsFold applies the ContainsFold predicate on the "remark" field.
func RemarkContainsFold(v string) predicate.Friend {
	return predicate.Friend(sql.FieldContainsFold(FieldRemark, v))
}

// GroupIDEQ applies the EQ predicate on the "group_id" field.
func GroupIDEQ(v int64) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldGroupID, v))
}

// GroupIDNEQ applies the NEQ predicate on the "group_id" field.
func GroupIDNEQ(v int64) predicate.Friend {
	return predicate.Friend(sql.FieldNEQ(FieldGroupID, v))
}

// GroupIDIn applies the In predicate on the "group_id" field.
func GroupIDIn(vs ...int64) predicate.Friend {
	return predicate.Friend(sql.FieldIn(FieldGroupID, vs...))
}

// GroupIDNotIn applies the NotIn predicate on the "group_id" field.
func GroupIDNotIn(vs ...int64) predicate.Friend {
	return predicate.Friend(sql.FieldNotIn(FieldGroupID, vs...))
}

// LastTalkAtEQ applies the EQ predicate on the "last_talk_at" field.
func LastTalkAtEQ(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldLastTalkAt, v))
}

// LastTalkAtNEQ applies the NEQ predicate on the "last_talk_at" field.
func LastTalkAtNEQ(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldNEQ(FieldLastTalkAt, v))
}

// LastTalkAtIn applies the In predicate on the "last_talk_at" field.
func LastTalkAtIn(vs ...time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldIn(FieldLastTalkAt, vs...))
}

// LastTalkAtNotIn applies the NotIn predicate on the "last_talk_at" field.
func LastTalkAtNotIn(vs ...time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldNotIn(FieldLastTalkAt, vs...))
}

// LastTalkAtGT applies the GT predicate on the "last_talk_at" field.
func LastTalkAtGT(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldGT(FieldLastTalkAt, v))
}

// LastTalkAtGTE applies the GTE predicate on the "last_talk_at" field.
func LastTalkAtGTE(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldGTE(FieldLastTalkAt, v))
}

// LastTalkAtLT applies the LT predicate on the "last_talk_at" field.
func LastTalkAtLT(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldLT(FieldLastTalkAt, v))
}

// LastTalkAtLTE applies the LTE predicate on the "last_talk_at" field.
func LastTalkAtLTE(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldLTE(FieldLastTalkAt, v))
}

// HasOwnerUser applies the HasEdge predicate on the "owner_user" edge.
func HasOwnerUser() predicate.Friend {
	return predicate.Friend(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerUserTable, OwnerUserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerUserWith applies the HasEdge predicate on the "owner_user" edge with a given conditions (other predicates).
func HasOwnerUserWith(preds ...predicate.User) predicate.Friend {
	return predicate.Friend(func(s *sql.Selector) {
		step := newOwnerUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFriendUser applies the HasEdge predicate on the "friend_user" edge.
func HasFriendUser() predicate.Friend {
	return predicate.Friend(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, FriendUserTable, FriendUserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFriendUserWith applies the HasEdge predicate on the "friend_user" edge with a given conditions (other predicates).
func HasFriendUserWith(preds ...predicate.User) predicate.Friend {
	return predicate.Friend(func(s *sql.Selector) {
		step := newFriendUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFriendGroupFriend applies the HasEdge predicate on the "friend_group_friend" edge.
func HasFriendGroupFriend() predicate.Friend {
	return predicate.Friend(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, FriendGroupFriendTable, FriendGroupFriendColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFriendGroupFriendWith applies the HasEdge predicate on the "friend_group_friend" edge with a given conditions (other predicates).
func HasFriendGroupFriendWith(preds ...predicate.FriendGroup) predicate.Friend {
	return predicate.Friend(func(s *sql.Selector) {
		step := newFriendGroupFriendStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Friend) predicate.Friend {
	return predicate.Friend(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Friend) predicate.Friend {
	return predicate.Friend(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Friend) predicate.Friend {
	return predicate.Friend(sql.NotPredicates(p))
}
