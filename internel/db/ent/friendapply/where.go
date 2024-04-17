// Code generated by ent, DO NOT EDIT.

package friendapply

import (
	"IM/internel/db/ent/predicate"
	"IM/internel/types/enums"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldEQ(FieldDeletedAt, v))
}

// FromUserID applies equality check predicate on the "from_user_id" field. It's identical to FromUserIDEQ.
func FromUserID(v int64) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldEQ(FieldFromUserID, v))
}

// FromNickname applies equality check predicate on the "from_nickname" field. It's identical to FromNicknameEQ.
func FromNickname(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldEQ(FieldFromNickname, v))
}

// FromFaceURL applies equality check predicate on the "from_face_url" field. It's identical to FromFaceURLEQ.
func FromFaceURL(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldEQ(FieldFromFaceURL, v))
}

// ToUserID applies equality check predicate on the "to_user_id" field. It's identical to ToUserIDEQ.
func ToUserID(v int64) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldEQ(FieldToUserID, v))
}

// ToFaceURL applies equality check predicate on the "to_face_url" field. It's identical to ToFaceURLEQ.
func ToFaceURL(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldEQ(FieldToFaceURL, v))
}

// ToNickname applies equality check predicate on the "to_nickname" field. It's identical to ToNicknameEQ.
func ToNickname(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldEQ(FieldToNickname, v))
}

// Result applies equality check predicate on the "result" field. It's identical to ResultEQ.
func Result(v enums.ApplyType) predicate.FriendApply {
	vc := string(v)
	return predicate.FriendApply(sql.FieldEQ(FieldResult, vc))
}

// ReqMsg applies equality check predicate on the "req_msg" field. It's identical to ReqMsgEQ.
func ReqMsg(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldEQ(FieldReqMsg, v))
}

// GroupID applies equality check predicate on the "group_id" field. It's identical to GroupIDEQ.
func GroupID(v int64) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldEQ(FieldGroupID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldLTE(FieldDeletedAt, v))
}

// FromUserIDEQ applies the EQ predicate on the "from_user_id" field.
func FromUserIDEQ(v int64) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldEQ(FieldFromUserID, v))
}

// FromUserIDNEQ applies the NEQ predicate on the "from_user_id" field.
func FromUserIDNEQ(v int64) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldNEQ(FieldFromUserID, v))
}

// FromUserIDIn applies the In predicate on the "from_user_id" field.
func FromUserIDIn(vs ...int64) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldIn(FieldFromUserID, vs...))
}

// FromUserIDNotIn applies the NotIn predicate on the "from_user_id" field.
func FromUserIDNotIn(vs ...int64) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldNotIn(FieldFromUserID, vs...))
}

// FromNicknameEQ applies the EQ predicate on the "from_nickname" field.
func FromNicknameEQ(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldEQ(FieldFromNickname, v))
}

// FromNicknameNEQ applies the NEQ predicate on the "from_nickname" field.
func FromNicknameNEQ(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldNEQ(FieldFromNickname, v))
}

// FromNicknameIn applies the In predicate on the "from_nickname" field.
func FromNicknameIn(vs ...string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldIn(FieldFromNickname, vs...))
}

// FromNicknameNotIn applies the NotIn predicate on the "from_nickname" field.
func FromNicknameNotIn(vs ...string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldNotIn(FieldFromNickname, vs...))
}

// FromNicknameGT applies the GT predicate on the "from_nickname" field.
func FromNicknameGT(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldGT(FieldFromNickname, v))
}

// FromNicknameGTE applies the GTE predicate on the "from_nickname" field.
func FromNicknameGTE(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldGTE(FieldFromNickname, v))
}

// FromNicknameLT applies the LT predicate on the "from_nickname" field.
func FromNicknameLT(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldLT(FieldFromNickname, v))
}

// FromNicknameLTE applies the LTE predicate on the "from_nickname" field.
func FromNicknameLTE(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldLTE(FieldFromNickname, v))
}

// FromNicknameContains applies the Contains predicate on the "from_nickname" field.
func FromNicknameContains(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldContains(FieldFromNickname, v))
}

// FromNicknameHasPrefix applies the HasPrefix predicate on the "from_nickname" field.
func FromNicknameHasPrefix(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldHasPrefix(FieldFromNickname, v))
}

// FromNicknameHasSuffix applies the HasSuffix predicate on the "from_nickname" field.
func FromNicknameHasSuffix(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldHasSuffix(FieldFromNickname, v))
}

// FromNicknameEqualFold applies the EqualFold predicate on the "from_nickname" field.
func FromNicknameEqualFold(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldEqualFold(FieldFromNickname, v))
}

// FromNicknameContainsFold applies the ContainsFold predicate on the "from_nickname" field.
func FromNicknameContainsFold(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldContainsFold(FieldFromNickname, v))
}

// FromFaceURLEQ applies the EQ predicate on the "from_face_url" field.
func FromFaceURLEQ(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldEQ(FieldFromFaceURL, v))
}

// FromFaceURLNEQ applies the NEQ predicate on the "from_face_url" field.
func FromFaceURLNEQ(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldNEQ(FieldFromFaceURL, v))
}

// FromFaceURLIn applies the In predicate on the "from_face_url" field.
func FromFaceURLIn(vs ...string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldIn(FieldFromFaceURL, vs...))
}

// FromFaceURLNotIn applies the NotIn predicate on the "from_face_url" field.
func FromFaceURLNotIn(vs ...string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldNotIn(FieldFromFaceURL, vs...))
}

// FromFaceURLGT applies the GT predicate on the "from_face_url" field.
func FromFaceURLGT(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldGT(FieldFromFaceURL, v))
}

// FromFaceURLGTE applies the GTE predicate on the "from_face_url" field.
func FromFaceURLGTE(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldGTE(FieldFromFaceURL, v))
}

// FromFaceURLLT applies the LT predicate on the "from_face_url" field.
func FromFaceURLLT(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldLT(FieldFromFaceURL, v))
}

// FromFaceURLLTE applies the LTE predicate on the "from_face_url" field.
func FromFaceURLLTE(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldLTE(FieldFromFaceURL, v))
}

// FromFaceURLContains applies the Contains predicate on the "from_face_url" field.
func FromFaceURLContains(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldContains(FieldFromFaceURL, v))
}

// FromFaceURLHasPrefix applies the HasPrefix predicate on the "from_face_url" field.
func FromFaceURLHasPrefix(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldHasPrefix(FieldFromFaceURL, v))
}

// FromFaceURLHasSuffix applies the HasSuffix predicate on the "from_face_url" field.
func FromFaceURLHasSuffix(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldHasSuffix(FieldFromFaceURL, v))
}

// FromFaceURLEqualFold applies the EqualFold predicate on the "from_face_url" field.
func FromFaceURLEqualFold(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldEqualFold(FieldFromFaceURL, v))
}

// FromFaceURLContainsFold applies the ContainsFold predicate on the "from_face_url" field.
func FromFaceURLContainsFold(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldContainsFold(FieldFromFaceURL, v))
}

// ToUserIDEQ applies the EQ predicate on the "to_user_id" field.
func ToUserIDEQ(v int64) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldEQ(FieldToUserID, v))
}

// ToUserIDNEQ applies the NEQ predicate on the "to_user_id" field.
func ToUserIDNEQ(v int64) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldNEQ(FieldToUserID, v))
}

// ToUserIDIn applies the In predicate on the "to_user_id" field.
func ToUserIDIn(vs ...int64) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldIn(FieldToUserID, vs...))
}

// ToUserIDNotIn applies the NotIn predicate on the "to_user_id" field.
func ToUserIDNotIn(vs ...int64) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldNotIn(FieldToUserID, vs...))
}

// ToFaceURLEQ applies the EQ predicate on the "to_face_url" field.
func ToFaceURLEQ(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldEQ(FieldToFaceURL, v))
}

// ToFaceURLNEQ applies the NEQ predicate on the "to_face_url" field.
func ToFaceURLNEQ(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldNEQ(FieldToFaceURL, v))
}

// ToFaceURLIn applies the In predicate on the "to_face_url" field.
func ToFaceURLIn(vs ...string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldIn(FieldToFaceURL, vs...))
}

// ToFaceURLNotIn applies the NotIn predicate on the "to_face_url" field.
func ToFaceURLNotIn(vs ...string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldNotIn(FieldToFaceURL, vs...))
}

// ToFaceURLGT applies the GT predicate on the "to_face_url" field.
func ToFaceURLGT(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldGT(FieldToFaceURL, v))
}

// ToFaceURLGTE applies the GTE predicate on the "to_face_url" field.
func ToFaceURLGTE(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldGTE(FieldToFaceURL, v))
}

// ToFaceURLLT applies the LT predicate on the "to_face_url" field.
func ToFaceURLLT(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldLT(FieldToFaceURL, v))
}

// ToFaceURLLTE applies the LTE predicate on the "to_face_url" field.
func ToFaceURLLTE(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldLTE(FieldToFaceURL, v))
}

// ToFaceURLContains applies the Contains predicate on the "to_face_url" field.
func ToFaceURLContains(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldContains(FieldToFaceURL, v))
}

// ToFaceURLHasPrefix applies the HasPrefix predicate on the "to_face_url" field.
func ToFaceURLHasPrefix(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldHasPrefix(FieldToFaceURL, v))
}

// ToFaceURLHasSuffix applies the HasSuffix predicate on the "to_face_url" field.
func ToFaceURLHasSuffix(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldHasSuffix(FieldToFaceURL, v))
}

// ToFaceURLEqualFold applies the EqualFold predicate on the "to_face_url" field.
func ToFaceURLEqualFold(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldEqualFold(FieldToFaceURL, v))
}

// ToFaceURLContainsFold applies the ContainsFold predicate on the "to_face_url" field.
func ToFaceURLContainsFold(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldContainsFold(FieldToFaceURL, v))
}

// ToNicknameEQ applies the EQ predicate on the "to_nickname" field.
func ToNicknameEQ(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldEQ(FieldToNickname, v))
}

// ToNicknameNEQ applies the NEQ predicate on the "to_nickname" field.
func ToNicknameNEQ(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldNEQ(FieldToNickname, v))
}

// ToNicknameIn applies the In predicate on the "to_nickname" field.
func ToNicknameIn(vs ...string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldIn(FieldToNickname, vs...))
}

// ToNicknameNotIn applies the NotIn predicate on the "to_nickname" field.
func ToNicknameNotIn(vs ...string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldNotIn(FieldToNickname, vs...))
}

// ToNicknameGT applies the GT predicate on the "to_nickname" field.
func ToNicknameGT(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldGT(FieldToNickname, v))
}

// ToNicknameGTE applies the GTE predicate on the "to_nickname" field.
func ToNicknameGTE(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldGTE(FieldToNickname, v))
}

// ToNicknameLT applies the LT predicate on the "to_nickname" field.
func ToNicknameLT(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldLT(FieldToNickname, v))
}

// ToNicknameLTE applies the LTE predicate on the "to_nickname" field.
func ToNicknameLTE(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldLTE(FieldToNickname, v))
}

// ToNicknameContains applies the Contains predicate on the "to_nickname" field.
func ToNicknameContains(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldContains(FieldToNickname, v))
}

// ToNicknameHasPrefix applies the HasPrefix predicate on the "to_nickname" field.
func ToNicknameHasPrefix(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldHasPrefix(FieldToNickname, v))
}

// ToNicknameHasSuffix applies the HasSuffix predicate on the "to_nickname" field.
func ToNicknameHasSuffix(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldHasSuffix(FieldToNickname, v))
}

// ToNicknameEqualFold applies the EqualFold predicate on the "to_nickname" field.
func ToNicknameEqualFold(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldEqualFold(FieldToNickname, v))
}

// ToNicknameContainsFold applies the ContainsFold predicate on the "to_nickname" field.
func ToNicknameContainsFold(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldContainsFold(FieldToNickname, v))
}

// ResultEQ applies the EQ predicate on the "result" field.
func ResultEQ(v enums.ApplyType) predicate.FriendApply {
	vc := string(v)
	return predicate.FriendApply(sql.FieldEQ(FieldResult, vc))
}

// ResultNEQ applies the NEQ predicate on the "result" field.
func ResultNEQ(v enums.ApplyType) predicate.FriendApply {
	vc := string(v)
	return predicate.FriendApply(sql.FieldNEQ(FieldResult, vc))
}

// ResultIn applies the In predicate on the "result" field.
func ResultIn(vs ...enums.ApplyType) predicate.FriendApply {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.FriendApply(sql.FieldIn(FieldResult, v...))
}

// ResultNotIn applies the NotIn predicate on the "result" field.
func ResultNotIn(vs ...enums.ApplyType) predicate.FriendApply {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.FriendApply(sql.FieldNotIn(FieldResult, v...))
}

// ResultGT applies the GT predicate on the "result" field.
func ResultGT(v enums.ApplyType) predicate.FriendApply {
	vc := string(v)
	return predicate.FriendApply(sql.FieldGT(FieldResult, vc))
}

// ResultGTE applies the GTE predicate on the "result" field.
func ResultGTE(v enums.ApplyType) predicate.FriendApply {
	vc := string(v)
	return predicate.FriendApply(sql.FieldGTE(FieldResult, vc))
}

// ResultLT applies the LT predicate on the "result" field.
func ResultLT(v enums.ApplyType) predicate.FriendApply {
	vc := string(v)
	return predicate.FriendApply(sql.FieldLT(FieldResult, vc))
}

// ResultLTE applies the LTE predicate on the "result" field.
func ResultLTE(v enums.ApplyType) predicate.FriendApply {
	vc := string(v)
	return predicate.FriendApply(sql.FieldLTE(FieldResult, vc))
}

// ResultContains applies the Contains predicate on the "result" field.
func ResultContains(v enums.ApplyType) predicate.FriendApply {
	vc := string(v)
	return predicate.FriendApply(sql.FieldContains(FieldResult, vc))
}

// ResultHasPrefix applies the HasPrefix predicate on the "result" field.
func ResultHasPrefix(v enums.ApplyType) predicate.FriendApply {
	vc := string(v)
	return predicate.FriendApply(sql.FieldHasPrefix(FieldResult, vc))
}

// ResultHasSuffix applies the HasSuffix predicate on the "result" field.
func ResultHasSuffix(v enums.ApplyType) predicate.FriendApply {
	vc := string(v)
	return predicate.FriendApply(sql.FieldHasSuffix(FieldResult, vc))
}

// ResultEqualFold applies the EqualFold predicate on the "result" field.
func ResultEqualFold(v enums.ApplyType) predicate.FriendApply {
	vc := string(v)
	return predicate.FriendApply(sql.FieldEqualFold(FieldResult, vc))
}

// ResultContainsFold applies the ContainsFold predicate on the "result" field.
func ResultContainsFold(v enums.ApplyType) predicate.FriendApply {
	vc := string(v)
	return predicate.FriendApply(sql.FieldContainsFold(FieldResult, vc))
}

// ReqMsgEQ applies the EQ predicate on the "req_msg" field.
func ReqMsgEQ(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldEQ(FieldReqMsg, v))
}

// ReqMsgNEQ applies the NEQ predicate on the "req_msg" field.
func ReqMsgNEQ(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldNEQ(FieldReqMsg, v))
}

// ReqMsgIn applies the In predicate on the "req_msg" field.
func ReqMsgIn(vs ...string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldIn(FieldReqMsg, vs...))
}

// ReqMsgNotIn applies the NotIn predicate on the "req_msg" field.
func ReqMsgNotIn(vs ...string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldNotIn(FieldReqMsg, vs...))
}

// ReqMsgGT applies the GT predicate on the "req_msg" field.
func ReqMsgGT(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldGT(FieldReqMsg, v))
}

// ReqMsgGTE applies the GTE predicate on the "req_msg" field.
func ReqMsgGTE(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldGTE(FieldReqMsg, v))
}

// ReqMsgLT applies the LT predicate on the "req_msg" field.
func ReqMsgLT(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldLT(FieldReqMsg, v))
}

// ReqMsgLTE applies the LTE predicate on the "req_msg" field.
func ReqMsgLTE(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldLTE(FieldReqMsg, v))
}

// ReqMsgContains applies the Contains predicate on the "req_msg" field.
func ReqMsgContains(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldContains(FieldReqMsg, v))
}

// ReqMsgHasPrefix applies the HasPrefix predicate on the "req_msg" field.
func ReqMsgHasPrefix(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldHasPrefix(FieldReqMsg, v))
}

// ReqMsgHasSuffix applies the HasSuffix predicate on the "req_msg" field.
func ReqMsgHasSuffix(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldHasSuffix(FieldReqMsg, v))
}

// ReqMsgEqualFold applies the EqualFold predicate on the "req_msg" field.
func ReqMsgEqualFold(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldEqualFold(FieldReqMsg, v))
}

// ReqMsgContainsFold applies the ContainsFold predicate on the "req_msg" field.
func ReqMsgContainsFold(v string) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldContainsFold(FieldReqMsg, v))
}

// GroupIDEQ applies the EQ predicate on the "group_id" field.
func GroupIDEQ(v int64) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldEQ(FieldGroupID, v))
}

// GroupIDNEQ applies the NEQ predicate on the "group_id" field.
func GroupIDNEQ(v int64) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldNEQ(FieldGroupID, v))
}

// GroupIDIn applies the In predicate on the "group_id" field.
func GroupIDIn(vs ...int64) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldIn(FieldGroupID, vs...))
}

// GroupIDNotIn applies the NotIn predicate on the "group_id" field.
func GroupIDNotIn(vs ...int64) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldNotIn(FieldGroupID, vs...))
}

// GroupIDGT applies the GT predicate on the "group_id" field.
func GroupIDGT(v int64) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldGT(FieldGroupID, v))
}

// GroupIDGTE applies the GTE predicate on the "group_id" field.
func GroupIDGTE(v int64) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldGTE(FieldGroupID, v))
}

// GroupIDLT applies the LT predicate on the "group_id" field.
func GroupIDLT(v int64) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldLT(FieldGroupID, v))
}

// GroupIDLTE applies the LTE predicate on the "group_id" field.
func GroupIDLTE(v int64) predicate.FriendApply {
	return predicate.FriendApply(sql.FieldLTE(FieldGroupID, v))
}

// HasFromUser applies the HasEdge predicate on the "from_user" edge.
func HasFromUser() predicate.FriendApply {
	return predicate.FriendApply(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, FromUserTable, FromUserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFromUserWith applies the HasEdge predicate on the "from_user" edge with a given conditions (other predicates).
func HasFromUserWith(preds ...predicate.User) predicate.FriendApply {
	return predicate.FriendApply(func(s *sql.Selector) {
		step := newFromUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasToUser applies the HasEdge predicate on the "to_user" edge.
func HasToUser() predicate.FriendApply {
	return predicate.FriendApply(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ToUserTable, ToUserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasToUserWith applies the HasEdge predicate on the "to_user" edge with a given conditions (other predicates).
func HasToUserWith(preds ...predicate.User) predicate.FriendApply {
	return predicate.FriendApply(func(s *sql.Selector) {
		step := newToUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.FriendApply) predicate.FriendApply {
	return predicate.FriendApply(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.FriendApply) predicate.FriendApply {
	return predicate.FriendApply(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.FriendApply) predicate.FriendApply {
	return predicate.FriendApply(sql.NotPredicates(p))
}