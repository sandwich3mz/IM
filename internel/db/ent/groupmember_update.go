// Code generated by ent, DO NOT EDIT.

package ent

import (
	"IM/internel/db/ent/group"
	"IM/internel/db/ent/groupmember"
	"IM/internel/db/ent/predicate"
	"IM/internel/db/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GroupMemberUpdate is the builder for updating GroupMember entities.
type GroupMemberUpdate struct {
	config
	hooks    []Hook
	mutation *GroupMemberMutation
}

// Where appends a list predicates to the GroupMemberUpdate builder.
func (gmu *GroupMemberUpdate) Where(ps ...predicate.GroupMember) *GroupMemberUpdate {
	gmu.mutation.Where(ps...)
	return gmu
}

// SetUpdatedAt sets the "updated_at" field.
func (gmu *GroupMemberUpdate) SetUpdatedAt(t time.Time) *GroupMemberUpdate {
	gmu.mutation.SetUpdatedAt(t)
	return gmu
}

// SetDeletedAt sets the "deleted_at" field.
func (gmu *GroupMemberUpdate) SetDeletedAt(t time.Time) *GroupMemberUpdate {
	gmu.mutation.SetDeletedAt(t)
	return gmu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gmu *GroupMemberUpdate) SetNillableDeletedAt(t *time.Time) *GroupMemberUpdate {
	if t != nil {
		gmu.SetDeletedAt(*t)
	}
	return gmu
}

// SetGroupID sets the "group_id" field.
func (gmu *GroupMemberUpdate) SetGroupID(i int64) *GroupMemberUpdate {
	gmu.mutation.SetGroupID(i)
	return gmu
}

// SetNillableGroupID sets the "group_id" field if the given value is not nil.
func (gmu *GroupMemberUpdate) SetNillableGroupID(i *int64) *GroupMemberUpdate {
	if i != nil {
		gmu.SetGroupID(*i)
	}
	return gmu
}

// SetUserID sets the "user_id" field.
func (gmu *GroupMemberUpdate) SetUserID(i int64) *GroupMemberUpdate {
	gmu.mutation.SetUserID(i)
	return gmu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (gmu *GroupMemberUpdate) SetNillableUserID(i *int64) *GroupMemberUpdate {
	if i != nil {
		gmu.SetUserID(*i)
	}
	return gmu
}

// SetNickname sets the "nickname" field.
func (gmu *GroupMemberUpdate) SetNickname(i int8) *GroupMemberUpdate {
	gmu.mutation.ResetNickname()
	gmu.mutation.SetNickname(i)
	return gmu
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (gmu *GroupMemberUpdate) SetNillableNickname(i *int8) *GroupMemberUpdate {
	if i != nil {
		gmu.SetNickname(*i)
	}
	return gmu
}

// AddNickname adds i to the "nickname" field.
func (gmu *GroupMemberUpdate) AddNickname(i int8) *GroupMemberUpdate {
	gmu.mutation.AddNickname(i)
	return gmu
}

// SetFaceURL sets the "face_url" field.
func (gmu *GroupMemberUpdate) SetFaceURL(s string) *GroupMemberUpdate {
	gmu.mutation.SetFaceURL(s)
	return gmu
}

// SetNillableFaceURL sets the "face_url" field if the given value is not nil.
func (gmu *GroupMemberUpdate) SetNillableFaceURL(s *string) *GroupMemberUpdate {
	if s != nil {
		gmu.SetFaceURL(*s)
	}
	return gmu
}

// SetCount sets the "count" field.
func (gmu *GroupMemberUpdate) SetCount(i int64) *GroupMemberUpdate {
	gmu.mutation.ResetCount()
	gmu.mutation.SetCount(i)
	return gmu
}

// SetNillableCount sets the "count" field if the given value is not nil.
func (gmu *GroupMemberUpdate) SetNillableCount(i *int64) *GroupMemberUpdate {
	if i != nil {
		gmu.SetCount(*i)
	}
	return gmu
}

// AddCount adds i to the "count" field.
func (gmu *GroupMemberUpdate) AddCount(i int64) *GroupMemberUpdate {
	gmu.mutation.AddCount(i)
	return gmu
}

// SetRemark sets the "remark" field.
func (gmu *GroupMemberUpdate) SetRemark(s string) *GroupMemberUpdate {
	gmu.mutation.SetRemark(s)
	return gmu
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (gmu *GroupMemberUpdate) SetNillableRemark(s *string) *GroupMemberUpdate {
	if s != nil {
		gmu.SetRemark(*s)
	}
	return gmu
}

// SetGroupUserID sets the "group_user" edge to the User entity by ID.
func (gmu *GroupMemberUpdate) SetGroupUserID(id int64) *GroupMemberUpdate {
	gmu.mutation.SetGroupUserID(id)
	return gmu
}

// SetGroupUser sets the "group_user" edge to the User entity.
func (gmu *GroupMemberUpdate) SetGroupUser(u *User) *GroupMemberUpdate {
	return gmu.SetGroupUserID(u.ID)
}

// SetMemberGroupID sets the "member_group" edge to the Group entity by ID.
func (gmu *GroupMemberUpdate) SetMemberGroupID(id int64) *GroupMemberUpdate {
	gmu.mutation.SetMemberGroupID(id)
	return gmu
}

// SetMemberGroup sets the "member_group" edge to the Group entity.
func (gmu *GroupMemberUpdate) SetMemberGroup(g *Group) *GroupMemberUpdate {
	return gmu.SetMemberGroupID(g.ID)
}

// Mutation returns the GroupMemberMutation object of the builder.
func (gmu *GroupMemberUpdate) Mutation() *GroupMemberMutation {
	return gmu.mutation
}

// ClearGroupUser clears the "group_user" edge to the User entity.
func (gmu *GroupMemberUpdate) ClearGroupUser() *GroupMemberUpdate {
	gmu.mutation.ClearGroupUser()
	return gmu
}

// ClearMemberGroup clears the "member_group" edge to the Group entity.
func (gmu *GroupMemberUpdate) ClearMemberGroup() *GroupMemberUpdate {
	gmu.mutation.ClearMemberGroup()
	return gmu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gmu *GroupMemberUpdate) Save(ctx context.Context) (int, error) {
	gmu.defaults()
	return withHooks(ctx, gmu.sqlSave, gmu.mutation, gmu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gmu *GroupMemberUpdate) SaveX(ctx context.Context) int {
	affected, err := gmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gmu *GroupMemberUpdate) Exec(ctx context.Context) error {
	_, err := gmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gmu *GroupMemberUpdate) ExecX(ctx context.Context) {
	if err := gmu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gmu *GroupMemberUpdate) defaults() {
	if _, ok := gmu.mutation.UpdatedAt(); !ok {
		v := groupmember.UpdateDefaultUpdatedAt()
		gmu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gmu *GroupMemberUpdate) check() error {
	if _, ok := gmu.mutation.GroupUserID(); gmu.mutation.GroupUserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "GroupMember.group_user"`)
	}
	if _, ok := gmu.mutation.MemberGroupID(); gmu.mutation.MemberGroupCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "GroupMember.member_group"`)
	}
	return nil
}

func (gmu *GroupMemberUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := gmu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(groupmember.Table, groupmember.Columns, sqlgraph.NewFieldSpec(groupmember.FieldID, field.TypeInt64))
	if ps := gmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gmu.mutation.UpdatedAt(); ok {
		_spec.SetField(groupmember.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := gmu.mutation.DeletedAt(); ok {
		_spec.SetField(groupmember.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := gmu.mutation.Nickname(); ok {
		_spec.SetField(groupmember.FieldNickname, field.TypeInt8, value)
	}
	if value, ok := gmu.mutation.AddedNickname(); ok {
		_spec.AddField(groupmember.FieldNickname, field.TypeInt8, value)
	}
	if value, ok := gmu.mutation.FaceURL(); ok {
		_spec.SetField(groupmember.FieldFaceURL, field.TypeString, value)
	}
	if value, ok := gmu.mutation.Count(); ok {
		_spec.SetField(groupmember.FieldCount, field.TypeInt64, value)
	}
	if value, ok := gmu.mutation.AddedCount(); ok {
		_spec.AddField(groupmember.FieldCount, field.TypeInt64, value)
	}
	if value, ok := gmu.mutation.Remark(); ok {
		_spec.SetField(groupmember.FieldRemark, field.TypeString, value)
	}
	if gmu.mutation.GroupUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   groupmember.GroupUserTable,
			Columns: []string{groupmember.GroupUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gmu.mutation.GroupUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   groupmember.GroupUserTable,
			Columns: []string{groupmember.GroupUserColumn},
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
	if gmu.mutation.MemberGroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   groupmember.MemberGroupTable,
			Columns: []string{groupmember.MemberGroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gmu.mutation.MemberGroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   groupmember.MemberGroupTable,
			Columns: []string{groupmember.MemberGroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{groupmember.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	gmu.mutation.done = true
	return n, nil
}

// GroupMemberUpdateOne is the builder for updating a single GroupMember entity.
type GroupMemberUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GroupMemberMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (gmuo *GroupMemberUpdateOne) SetUpdatedAt(t time.Time) *GroupMemberUpdateOne {
	gmuo.mutation.SetUpdatedAt(t)
	return gmuo
}

// SetDeletedAt sets the "deleted_at" field.
func (gmuo *GroupMemberUpdateOne) SetDeletedAt(t time.Time) *GroupMemberUpdateOne {
	gmuo.mutation.SetDeletedAt(t)
	return gmuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gmuo *GroupMemberUpdateOne) SetNillableDeletedAt(t *time.Time) *GroupMemberUpdateOne {
	if t != nil {
		gmuo.SetDeletedAt(*t)
	}
	return gmuo
}

// SetGroupID sets the "group_id" field.
func (gmuo *GroupMemberUpdateOne) SetGroupID(i int64) *GroupMemberUpdateOne {
	gmuo.mutation.SetGroupID(i)
	return gmuo
}

// SetNillableGroupID sets the "group_id" field if the given value is not nil.
func (gmuo *GroupMemberUpdateOne) SetNillableGroupID(i *int64) *GroupMemberUpdateOne {
	if i != nil {
		gmuo.SetGroupID(*i)
	}
	return gmuo
}

// SetUserID sets the "user_id" field.
func (gmuo *GroupMemberUpdateOne) SetUserID(i int64) *GroupMemberUpdateOne {
	gmuo.mutation.SetUserID(i)
	return gmuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (gmuo *GroupMemberUpdateOne) SetNillableUserID(i *int64) *GroupMemberUpdateOne {
	if i != nil {
		gmuo.SetUserID(*i)
	}
	return gmuo
}

// SetNickname sets the "nickname" field.
func (gmuo *GroupMemberUpdateOne) SetNickname(i int8) *GroupMemberUpdateOne {
	gmuo.mutation.ResetNickname()
	gmuo.mutation.SetNickname(i)
	return gmuo
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (gmuo *GroupMemberUpdateOne) SetNillableNickname(i *int8) *GroupMemberUpdateOne {
	if i != nil {
		gmuo.SetNickname(*i)
	}
	return gmuo
}

// AddNickname adds i to the "nickname" field.
func (gmuo *GroupMemberUpdateOne) AddNickname(i int8) *GroupMemberUpdateOne {
	gmuo.mutation.AddNickname(i)
	return gmuo
}

// SetFaceURL sets the "face_url" field.
func (gmuo *GroupMemberUpdateOne) SetFaceURL(s string) *GroupMemberUpdateOne {
	gmuo.mutation.SetFaceURL(s)
	return gmuo
}

// SetNillableFaceURL sets the "face_url" field if the given value is not nil.
func (gmuo *GroupMemberUpdateOne) SetNillableFaceURL(s *string) *GroupMemberUpdateOne {
	if s != nil {
		gmuo.SetFaceURL(*s)
	}
	return gmuo
}

// SetCount sets the "count" field.
func (gmuo *GroupMemberUpdateOne) SetCount(i int64) *GroupMemberUpdateOne {
	gmuo.mutation.ResetCount()
	gmuo.mutation.SetCount(i)
	return gmuo
}

// SetNillableCount sets the "count" field if the given value is not nil.
func (gmuo *GroupMemberUpdateOne) SetNillableCount(i *int64) *GroupMemberUpdateOne {
	if i != nil {
		gmuo.SetCount(*i)
	}
	return gmuo
}

// AddCount adds i to the "count" field.
func (gmuo *GroupMemberUpdateOne) AddCount(i int64) *GroupMemberUpdateOne {
	gmuo.mutation.AddCount(i)
	return gmuo
}

// SetRemark sets the "remark" field.
func (gmuo *GroupMemberUpdateOne) SetRemark(s string) *GroupMemberUpdateOne {
	gmuo.mutation.SetRemark(s)
	return gmuo
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (gmuo *GroupMemberUpdateOne) SetNillableRemark(s *string) *GroupMemberUpdateOne {
	if s != nil {
		gmuo.SetRemark(*s)
	}
	return gmuo
}

// SetGroupUserID sets the "group_user" edge to the User entity by ID.
func (gmuo *GroupMemberUpdateOne) SetGroupUserID(id int64) *GroupMemberUpdateOne {
	gmuo.mutation.SetGroupUserID(id)
	return gmuo
}

// SetGroupUser sets the "group_user" edge to the User entity.
func (gmuo *GroupMemberUpdateOne) SetGroupUser(u *User) *GroupMemberUpdateOne {
	return gmuo.SetGroupUserID(u.ID)
}

// SetMemberGroupID sets the "member_group" edge to the Group entity by ID.
func (gmuo *GroupMemberUpdateOne) SetMemberGroupID(id int64) *GroupMemberUpdateOne {
	gmuo.mutation.SetMemberGroupID(id)
	return gmuo
}

// SetMemberGroup sets the "member_group" edge to the Group entity.
func (gmuo *GroupMemberUpdateOne) SetMemberGroup(g *Group) *GroupMemberUpdateOne {
	return gmuo.SetMemberGroupID(g.ID)
}

// Mutation returns the GroupMemberMutation object of the builder.
func (gmuo *GroupMemberUpdateOne) Mutation() *GroupMemberMutation {
	return gmuo.mutation
}

// ClearGroupUser clears the "group_user" edge to the User entity.
func (gmuo *GroupMemberUpdateOne) ClearGroupUser() *GroupMemberUpdateOne {
	gmuo.mutation.ClearGroupUser()
	return gmuo
}

// ClearMemberGroup clears the "member_group" edge to the Group entity.
func (gmuo *GroupMemberUpdateOne) ClearMemberGroup() *GroupMemberUpdateOne {
	gmuo.mutation.ClearMemberGroup()
	return gmuo
}

// Where appends a list predicates to the GroupMemberUpdate builder.
func (gmuo *GroupMemberUpdateOne) Where(ps ...predicate.GroupMember) *GroupMemberUpdateOne {
	gmuo.mutation.Where(ps...)
	return gmuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (gmuo *GroupMemberUpdateOne) Select(field string, fields ...string) *GroupMemberUpdateOne {
	gmuo.fields = append([]string{field}, fields...)
	return gmuo
}

// Save executes the query and returns the updated GroupMember entity.
func (gmuo *GroupMemberUpdateOne) Save(ctx context.Context) (*GroupMember, error) {
	gmuo.defaults()
	return withHooks(ctx, gmuo.sqlSave, gmuo.mutation, gmuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gmuo *GroupMemberUpdateOne) SaveX(ctx context.Context) *GroupMember {
	node, err := gmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (gmuo *GroupMemberUpdateOne) Exec(ctx context.Context) error {
	_, err := gmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gmuo *GroupMemberUpdateOne) ExecX(ctx context.Context) {
	if err := gmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gmuo *GroupMemberUpdateOne) defaults() {
	if _, ok := gmuo.mutation.UpdatedAt(); !ok {
		v := groupmember.UpdateDefaultUpdatedAt()
		gmuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gmuo *GroupMemberUpdateOne) check() error {
	if _, ok := gmuo.mutation.GroupUserID(); gmuo.mutation.GroupUserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "GroupMember.group_user"`)
	}
	if _, ok := gmuo.mutation.MemberGroupID(); gmuo.mutation.MemberGroupCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "GroupMember.member_group"`)
	}
	return nil
}

func (gmuo *GroupMemberUpdateOne) sqlSave(ctx context.Context) (_node *GroupMember, err error) {
	if err := gmuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(groupmember.Table, groupmember.Columns, sqlgraph.NewFieldSpec(groupmember.FieldID, field.TypeInt64))
	id, ok := gmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "GroupMember.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := gmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, groupmember.FieldID)
		for _, f := range fields {
			if !groupmember.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != groupmember.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := gmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gmuo.mutation.UpdatedAt(); ok {
		_spec.SetField(groupmember.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := gmuo.mutation.DeletedAt(); ok {
		_spec.SetField(groupmember.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := gmuo.mutation.Nickname(); ok {
		_spec.SetField(groupmember.FieldNickname, field.TypeInt8, value)
	}
	if value, ok := gmuo.mutation.AddedNickname(); ok {
		_spec.AddField(groupmember.FieldNickname, field.TypeInt8, value)
	}
	if value, ok := gmuo.mutation.FaceURL(); ok {
		_spec.SetField(groupmember.FieldFaceURL, field.TypeString, value)
	}
	if value, ok := gmuo.mutation.Count(); ok {
		_spec.SetField(groupmember.FieldCount, field.TypeInt64, value)
	}
	if value, ok := gmuo.mutation.AddedCount(); ok {
		_spec.AddField(groupmember.FieldCount, field.TypeInt64, value)
	}
	if value, ok := gmuo.mutation.Remark(); ok {
		_spec.SetField(groupmember.FieldRemark, field.TypeString, value)
	}
	if gmuo.mutation.GroupUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   groupmember.GroupUserTable,
			Columns: []string{groupmember.GroupUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gmuo.mutation.GroupUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   groupmember.GroupUserTable,
			Columns: []string{groupmember.GroupUserColumn},
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
	if gmuo.mutation.MemberGroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   groupmember.MemberGroupTable,
			Columns: []string{groupmember.MemberGroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gmuo.mutation.MemberGroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   groupmember.MemberGroupTable,
			Columns: []string{groupmember.MemberGroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &GroupMember{config: gmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, gmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{groupmember.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	gmuo.mutation.done = true
	return _node, nil
}
