// Code generated by ent, DO NOT EDIT.

package ent

import (
	"IM/internel/db/ent/friendapply"
	"IM/internel/db/ent/predicate"
	"IM/internel/db/ent/user"
	"IM/internel/types/enums"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FriendApplyUpdate is the builder for updating FriendApply entities.
type FriendApplyUpdate struct {
	config
	hooks    []Hook
	mutation *FriendApplyMutation
}

// Where appends a list predicates to the FriendApplyUpdate builder.
func (fau *FriendApplyUpdate) Where(ps ...predicate.FriendApply) *FriendApplyUpdate {
	fau.mutation.Where(ps...)
	return fau
}

// SetUpdatedAt sets the "updated_at" field.
func (fau *FriendApplyUpdate) SetUpdatedAt(t time.Time) *FriendApplyUpdate {
	fau.mutation.SetUpdatedAt(t)
	return fau
}

// SetDeletedAt sets the "deleted_at" field.
func (fau *FriendApplyUpdate) SetDeletedAt(t time.Time) *FriendApplyUpdate {
	fau.mutation.SetDeletedAt(t)
	return fau
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (fau *FriendApplyUpdate) SetNillableDeletedAt(t *time.Time) *FriendApplyUpdate {
	if t != nil {
		fau.SetDeletedAt(*t)
	}
	return fau
}

// SetFromUserID sets the "from_user_id" field.
func (fau *FriendApplyUpdate) SetFromUserID(i int64) *FriendApplyUpdate {
	fau.mutation.SetFromUserID(i)
	return fau
}

// SetNillableFromUserID sets the "from_user_id" field if the given value is not nil.
func (fau *FriendApplyUpdate) SetNillableFromUserID(i *int64) *FriendApplyUpdate {
	if i != nil {
		fau.SetFromUserID(*i)
	}
	return fau
}

// SetFromNickname sets the "from_nickname" field.
func (fau *FriendApplyUpdate) SetFromNickname(s string) *FriendApplyUpdate {
	fau.mutation.SetFromNickname(s)
	return fau
}

// SetNillableFromNickname sets the "from_nickname" field if the given value is not nil.
func (fau *FriendApplyUpdate) SetNillableFromNickname(s *string) *FriendApplyUpdate {
	if s != nil {
		fau.SetFromNickname(*s)
	}
	return fau
}

// SetFromFaceURL sets the "from_face_url" field.
func (fau *FriendApplyUpdate) SetFromFaceURL(s string) *FriendApplyUpdate {
	fau.mutation.SetFromFaceURL(s)
	return fau
}

// SetNillableFromFaceURL sets the "from_face_url" field if the given value is not nil.
func (fau *FriendApplyUpdate) SetNillableFromFaceURL(s *string) *FriendApplyUpdate {
	if s != nil {
		fau.SetFromFaceURL(*s)
	}
	return fau
}

// SetToUserID sets the "to_user_id" field.
func (fau *FriendApplyUpdate) SetToUserID(i int64) *FriendApplyUpdate {
	fau.mutation.SetToUserID(i)
	return fau
}

// SetNillableToUserID sets the "to_user_id" field if the given value is not nil.
func (fau *FriendApplyUpdate) SetNillableToUserID(i *int64) *FriendApplyUpdate {
	if i != nil {
		fau.SetToUserID(*i)
	}
	return fau
}

// SetToFaceURL sets the "to_face_url" field.
func (fau *FriendApplyUpdate) SetToFaceURL(s string) *FriendApplyUpdate {
	fau.mutation.SetToFaceURL(s)
	return fau
}

// SetNillableToFaceURL sets the "to_face_url" field if the given value is not nil.
func (fau *FriendApplyUpdate) SetNillableToFaceURL(s *string) *FriendApplyUpdate {
	if s != nil {
		fau.SetToFaceURL(*s)
	}
	return fau
}

// SetToNickname sets the "to_nickname" field.
func (fau *FriendApplyUpdate) SetToNickname(s string) *FriendApplyUpdate {
	fau.mutation.SetToNickname(s)
	return fau
}

// SetNillableToNickname sets the "to_nickname" field if the given value is not nil.
func (fau *FriendApplyUpdate) SetNillableToNickname(s *string) *FriendApplyUpdate {
	if s != nil {
		fau.SetToNickname(*s)
	}
	return fau
}

// SetResult sets the "result" field.
func (fau *FriendApplyUpdate) SetResult(et enums.ApplyType) *FriendApplyUpdate {
	fau.mutation.SetResult(et)
	return fau
}

// SetNillableResult sets the "result" field if the given value is not nil.
func (fau *FriendApplyUpdate) SetNillableResult(et *enums.ApplyType) *FriendApplyUpdate {
	if et != nil {
		fau.SetResult(*et)
	}
	return fau
}

// SetReqMsg sets the "req_msg" field.
func (fau *FriendApplyUpdate) SetReqMsg(s string) *FriendApplyUpdate {
	fau.mutation.SetReqMsg(s)
	return fau
}

// SetNillableReqMsg sets the "req_msg" field if the given value is not nil.
func (fau *FriendApplyUpdate) SetNillableReqMsg(s *string) *FriendApplyUpdate {
	if s != nil {
		fau.SetReqMsg(*s)
	}
	return fau
}

// SetGroupID sets the "group_id" field.
func (fau *FriendApplyUpdate) SetGroupID(i int64) *FriendApplyUpdate {
	fau.mutation.ResetGroupID()
	fau.mutation.SetGroupID(i)
	return fau
}

// SetNillableGroupID sets the "group_id" field if the given value is not nil.
func (fau *FriendApplyUpdate) SetNillableGroupID(i *int64) *FriendApplyUpdate {
	if i != nil {
		fau.SetGroupID(*i)
	}
	return fau
}

// AddGroupID adds i to the "group_id" field.
func (fau *FriendApplyUpdate) AddGroupID(i int64) *FriendApplyUpdate {
	fau.mutation.AddGroupID(i)
	return fau
}

// SetFromUser sets the "from_user" edge to the User entity.
func (fau *FriendApplyUpdate) SetFromUser(u *User) *FriendApplyUpdate {
	return fau.SetFromUserID(u.ID)
}

// SetToUser sets the "to_user" edge to the User entity.
func (fau *FriendApplyUpdate) SetToUser(u *User) *FriendApplyUpdate {
	return fau.SetToUserID(u.ID)
}

// Mutation returns the FriendApplyMutation object of the builder.
func (fau *FriendApplyUpdate) Mutation() *FriendApplyMutation {
	return fau.mutation
}

// ClearFromUser clears the "from_user" edge to the User entity.
func (fau *FriendApplyUpdate) ClearFromUser() *FriendApplyUpdate {
	fau.mutation.ClearFromUser()
	return fau
}

// ClearToUser clears the "to_user" edge to the User entity.
func (fau *FriendApplyUpdate) ClearToUser() *FriendApplyUpdate {
	fau.mutation.ClearToUser()
	return fau
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fau *FriendApplyUpdate) Save(ctx context.Context) (int, error) {
	fau.defaults()
	return withHooks(ctx, fau.sqlSave, fau.mutation, fau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fau *FriendApplyUpdate) SaveX(ctx context.Context) int {
	affected, err := fau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fau *FriendApplyUpdate) Exec(ctx context.Context) error {
	_, err := fau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fau *FriendApplyUpdate) ExecX(ctx context.Context) {
	if err := fau.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fau *FriendApplyUpdate) defaults() {
	if _, ok := fau.mutation.UpdatedAt(); !ok {
		v := friendapply.UpdateDefaultUpdatedAt()
		fau.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fau *FriendApplyUpdate) check() error {
	if _, ok := fau.mutation.FromUserID(); fau.mutation.FromUserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "FriendApply.from_user"`)
	}
	if _, ok := fau.mutation.ToUserID(); fau.mutation.ToUserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "FriendApply.to_user"`)
	}
	return nil
}

func (fau *FriendApplyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := fau.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(friendapply.Table, friendapply.Columns, sqlgraph.NewFieldSpec(friendapply.FieldID, field.TypeInt64))
	if ps := fau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fau.mutation.UpdatedAt(); ok {
		_spec.SetField(friendapply.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := fau.mutation.DeletedAt(); ok {
		_spec.SetField(friendapply.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := fau.mutation.FromNickname(); ok {
		_spec.SetField(friendapply.FieldFromNickname, field.TypeString, value)
	}
	if value, ok := fau.mutation.FromFaceURL(); ok {
		_spec.SetField(friendapply.FieldFromFaceURL, field.TypeString, value)
	}
	if value, ok := fau.mutation.ToFaceURL(); ok {
		_spec.SetField(friendapply.FieldToFaceURL, field.TypeString, value)
	}
	if value, ok := fau.mutation.ToNickname(); ok {
		_spec.SetField(friendapply.FieldToNickname, field.TypeString, value)
	}
	if value, ok := fau.mutation.Result(); ok {
		_spec.SetField(friendapply.FieldResult, field.TypeString, value)
	}
	if value, ok := fau.mutation.ReqMsg(); ok {
		_spec.SetField(friendapply.FieldReqMsg, field.TypeString, value)
	}
	if value, ok := fau.mutation.GroupID(); ok {
		_spec.SetField(friendapply.FieldGroupID, field.TypeInt64, value)
	}
	if value, ok := fau.mutation.AddedGroupID(); ok {
		_spec.AddField(friendapply.FieldGroupID, field.TypeInt64, value)
	}
	if fau.mutation.FromUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   friendapply.FromUserTable,
			Columns: []string{friendapply.FromUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fau.mutation.FromUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   friendapply.FromUserTable,
			Columns: []string{friendapply.FromUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fau.mutation.ToUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   friendapply.ToUserTable,
			Columns: []string{friendapply.ToUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fau.mutation.ToUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   friendapply.ToUserTable,
			Columns: []string{friendapply.ToUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{friendapply.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	fau.mutation.done = true
	return n, nil
}

// FriendApplyUpdateOne is the builder for updating a single FriendApply entity.
type FriendApplyUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FriendApplyMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (fauo *FriendApplyUpdateOne) SetUpdatedAt(t time.Time) *FriendApplyUpdateOne {
	fauo.mutation.SetUpdatedAt(t)
	return fauo
}

// SetDeletedAt sets the "deleted_at" field.
func (fauo *FriendApplyUpdateOne) SetDeletedAt(t time.Time) *FriendApplyUpdateOne {
	fauo.mutation.SetDeletedAt(t)
	return fauo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (fauo *FriendApplyUpdateOne) SetNillableDeletedAt(t *time.Time) *FriendApplyUpdateOne {
	if t != nil {
		fauo.SetDeletedAt(*t)
	}
	return fauo
}

// SetFromUserID sets the "from_user_id" field.
func (fauo *FriendApplyUpdateOne) SetFromUserID(i int64) *FriendApplyUpdateOne {
	fauo.mutation.SetFromUserID(i)
	return fauo
}

// SetNillableFromUserID sets the "from_user_id" field if the given value is not nil.
func (fauo *FriendApplyUpdateOne) SetNillableFromUserID(i *int64) *FriendApplyUpdateOne {
	if i != nil {
		fauo.SetFromUserID(*i)
	}
	return fauo
}

// SetFromNickname sets the "from_nickname" field.
func (fauo *FriendApplyUpdateOne) SetFromNickname(s string) *FriendApplyUpdateOne {
	fauo.mutation.SetFromNickname(s)
	return fauo
}

// SetNillableFromNickname sets the "from_nickname" field if the given value is not nil.
func (fauo *FriendApplyUpdateOne) SetNillableFromNickname(s *string) *FriendApplyUpdateOne {
	if s != nil {
		fauo.SetFromNickname(*s)
	}
	return fauo
}

// SetFromFaceURL sets the "from_face_url" field.
func (fauo *FriendApplyUpdateOne) SetFromFaceURL(s string) *FriendApplyUpdateOne {
	fauo.mutation.SetFromFaceURL(s)
	return fauo
}

// SetNillableFromFaceURL sets the "from_face_url" field if the given value is not nil.
func (fauo *FriendApplyUpdateOne) SetNillableFromFaceURL(s *string) *FriendApplyUpdateOne {
	if s != nil {
		fauo.SetFromFaceURL(*s)
	}
	return fauo
}

// SetToUserID sets the "to_user_id" field.
func (fauo *FriendApplyUpdateOne) SetToUserID(i int64) *FriendApplyUpdateOne {
	fauo.mutation.SetToUserID(i)
	return fauo
}

// SetNillableToUserID sets the "to_user_id" field if the given value is not nil.
func (fauo *FriendApplyUpdateOne) SetNillableToUserID(i *int64) *FriendApplyUpdateOne {
	if i != nil {
		fauo.SetToUserID(*i)
	}
	return fauo
}

// SetToFaceURL sets the "to_face_url" field.
func (fauo *FriendApplyUpdateOne) SetToFaceURL(s string) *FriendApplyUpdateOne {
	fauo.mutation.SetToFaceURL(s)
	return fauo
}

// SetNillableToFaceURL sets the "to_face_url" field if the given value is not nil.
func (fauo *FriendApplyUpdateOne) SetNillableToFaceURL(s *string) *FriendApplyUpdateOne {
	if s != nil {
		fauo.SetToFaceURL(*s)
	}
	return fauo
}

// SetToNickname sets the "to_nickname" field.
func (fauo *FriendApplyUpdateOne) SetToNickname(s string) *FriendApplyUpdateOne {
	fauo.mutation.SetToNickname(s)
	return fauo
}

// SetNillableToNickname sets the "to_nickname" field if the given value is not nil.
func (fauo *FriendApplyUpdateOne) SetNillableToNickname(s *string) *FriendApplyUpdateOne {
	if s != nil {
		fauo.SetToNickname(*s)
	}
	return fauo
}

// SetResult sets the "result" field.
func (fauo *FriendApplyUpdateOne) SetResult(et enums.ApplyType) *FriendApplyUpdateOne {
	fauo.mutation.SetResult(et)
	return fauo
}

// SetNillableResult sets the "result" field if the given value is not nil.
func (fauo *FriendApplyUpdateOne) SetNillableResult(et *enums.ApplyType) *FriendApplyUpdateOne {
	if et != nil {
		fauo.SetResult(*et)
	}
	return fauo
}

// SetReqMsg sets the "req_msg" field.
func (fauo *FriendApplyUpdateOne) SetReqMsg(s string) *FriendApplyUpdateOne {
	fauo.mutation.SetReqMsg(s)
	return fauo
}

// SetNillableReqMsg sets the "req_msg" field if the given value is not nil.
func (fauo *FriendApplyUpdateOne) SetNillableReqMsg(s *string) *FriendApplyUpdateOne {
	if s != nil {
		fauo.SetReqMsg(*s)
	}
	return fauo
}

// SetGroupID sets the "group_id" field.
func (fauo *FriendApplyUpdateOne) SetGroupID(i int64) *FriendApplyUpdateOne {
	fauo.mutation.ResetGroupID()
	fauo.mutation.SetGroupID(i)
	return fauo
}

// SetNillableGroupID sets the "group_id" field if the given value is not nil.
func (fauo *FriendApplyUpdateOne) SetNillableGroupID(i *int64) *FriendApplyUpdateOne {
	if i != nil {
		fauo.SetGroupID(*i)
	}
	return fauo
}

// AddGroupID adds i to the "group_id" field.
func (fauo *FriendApplyUpdateOne) AddGroupID(i int64) *FriendApplyUpdateOne {
	fauo.mutation.AddGroupID(i)
	return fauo
}

// SetFromUser sets the "from_user" edge to the User entity.
func (fauo *FriendApplyUpdateOne) SetFromUser(u *User) *FriendApplyUpdateOne {
	return fauo.SetFromUserID(u.ID)
}

// SetToUser sets the "to_user" edge to the User entity.
func (fauo *FriendApplyUpdateOne) SetToUser(u *User) *FriendApplyUpdateOne {
	return fauo.SetToUserID(u.ID)
}

// Mutation returns the FriendApplyMutation object of the builder.
func (fauo *FriendApplyUpdateOne) Mutation() *FriendApplyMutation {
	return fauo.mutation
}

// ClearFromUser clears the "from_user" edge to the User entity.
func (fauo *FriendApplyUpdateOne) ClearFromUser() *FriendApplyUpdateOne {
	fauo.mutation.ClearFromUser()
	return fauo
}

// ClearToUser clears the "to_user" edge to the User entity.
func (fauo *FriendApplyUpdateOne) ClearToUser() *FriendApplyUpdateOne {
	fauo.mutation.ClearToUser()
	return fauo
}

// Where appends a list predicates to the FriendApplyUpdate builder.
func (fauo *FriendApplyUpdateOne) Where(ps ...predicate.FriendApply) *FriendApplyUpdateOne {
	fauo.mutation.Where(ps...)
	return fauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fauo *FriendApplyUpdateOne) Select(field string, fields ...string) *FriendApplyUpdateOne {
	fauo.fields = append([]string{field}, fields...)
	return fauo
}

// Save executes the query and returns the updated FriendApply entity.
func (fauo *FriendApplyUpdateOne) Save(ctx context.Context) (*FriendApply, error) {
	fauo.defaults()
	return withHooks(ctx, fauo.sqlSave, fauo.mutation, fauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fauo *FriendApplyUpdateOne) SaveX(ctx context.Context) *FriendApply {
	node, err := fauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fauo *FriendApplyUpdateOne) Exec(ctx context.Context) error {
	_, err := fauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fauo *FriendApplyUpdateOne) ExecX(ctx context.Context) {
	if err := fauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fauo *FriendApplyUpdateOne) defaults() {
	if _, ok := fauo.mutation.UpdatedAt(); !ok {
		v := friendapply.UpdateDefaultUpdatedAt()
		fauo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fauo *FriendApplyUpdateOne) check() error {
	if _, ok := fauo.mutation.FromUserID(); fauo.mutation.FromUserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "FriendApply.from_user"`)
	}
	if _, ok := fauo.mutation.ToUserID(); fauo.mutation.ToUserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "FriendApply.to_user"`)
	}
	return nil
}

func (fauo *FriendApplyUpdateOne) sqlSave(ctx context.Context) (_node *FriendApply, err error) {
	if err := fauo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(friendapply.Table, friendapply.Columns, sqlgraph.NewFieldSpec(friendapply.FieldID, field.TypeInt64))
	id, ok := fauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "FriendApply.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, friendapply.FieldID)
		for _, f := range fields {
			if !friendapply.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != friendapply.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fauo.mutation.UpdatedAt(); ok {
		_spec.SetField(friendapply.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := fauo.mutation.DeletedAt(); ok {
		_spec.SetField(friendapply.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := fauo.mutation.FromNickname(); ok {
		_spec.SetField(friendapply.FieldFromNickname, field.TypeString, value)
	}
	if value, ok := fauo.mutation.FromFaceURL(); ok {
		_spec.SetField(friendapply.FieldFromFaceURL, field.TypeString, value)
	}
	if value, ok := fauo.mutation.ToFaceURL(); ok {
		_spec.SetField(friendapply.FieldToFaceURL, field.TypeString, value)
	}
	if value, ok := fauo.mutation.ToNickname(); ok {
		_spec.SetField(friendapply.FieldToNickname, field.TypeString, value)
	}
	if value, ok := fauo.mutation.Result(); ok {
		_spec.SetField(friendapply.FieldResult, field.TypeString, value)
	}
	if value, ok := fauo.mutation.ReqMsg(); ok {
		_spec.SetField(friendapply.FieldReqMsg, field.TypeString, value)
	}
	if value, ok := fauo.mutation.GroupID(); ok {
		_spec.SetField(friendapply.FieldGroupID, field.TypeInt64, value)
	}
	if value, ok := fauo.mutation.AddedGroupID(); ok {
		_spec.AddField(friendapply.FieldGroupID, field.TypeInt64, value)
	}
	if fauo.mutation.FromUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   friendapply.FromUserTable,
			Columns: []string{friendapply.FromUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fauo.mutation.FromUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   friendapply.FromUserTable,
			Columns: []string{friendapply.FromUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fauo.mutation.ToUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   friendapply.ToUserTable,
			Columns: []string{friendapply.ToUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fauo.mutation.ToUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   friendapply.ToUserTable,
			Columns: []string{friendapply.ToUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &FriendApply{config: fauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{friendapply.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	fauo.mutation.done = true
	return _node, nil
}