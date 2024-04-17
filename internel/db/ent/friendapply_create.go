// Code generated by ent, DO NOT EDIT.

package ent

import (
	"IM/internel/db/ent/friendapply"
	"IM/internel/db/ent/user"
	"IM/internel/types/enums"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FriendApplyCreate is the builder for creating a FriendApply entity.
type FriendApplyCreate struct {
	config
	mutation *FriendApplyMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (fac *FriendApplyCreate) SetCreatedAt(t time.Time) *FriendApplyCreate {
	fac.mutation.SetCreatedAt(t)
	return fac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fac *FriendApplyCreate) SetNillableCreatedAt(t *time.Time) *FriendApplyCreate {
	if t != nil {
		fac.SetCreatedAt(*t)
	}
	return fac
}

// SetUpdatedAt sets the "updated_at" field.
func (fac *FriendApplyCreate) SetUpdatedAt(t time.Time) *FriendApplyCreate {
	fac.mutation.SetUpdatedAt(t)
	return fac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fac *FriendApplyCreate) SetNillableUpdatedAt(t *time.Time) *FriendApplyCreate {
	if t != nil {
		fac.SetUpdatedAt(*t)
	}
	return fac
}

// SetDeletedAt sets the "deleted_at" field.
func (fac *FriendApplyCreate) SetDeletedAt(t time.Time) *FriendApplyCreate {
	fac.mutation.SetDeletedAt(t)
	return fac
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (fac *FriendApplyCreate) SetNillableDeletedAt(t *time.Time) *FriendApplyCreate {
	if t != nil {
		fac.SetDeletedAt(*t)
	}
	return fac
}

// SetFromUserID sets the "from_user_id" field.
func (fac *FriendApplyCreate) SetFromUserID(i int64) *FriendApplyCreate {
	fac.mutation.SetFromUserID(i)
	return fac
}

// SetFromNickname sets the "from_nickname" field.
func (fac *FriendApplyCreate) SetFromNickname(s string) *FriendApplyCreate {
	fac.mutation.SetFromNickname(s)
	return fac
}

// SetFromFaceURL sets the "from_face_url" field.
func (fac *FriendApplyCreate) SetFromFaceURL(s string) *FriendApplyCreate {
	fac.mutation.SetFromFaceURL(s)
	return fac
}

// SetToUserID sets the "to_user_id" field.
func (fac *FriendApplyCreate) SetToUserID(i int64) *FriendApplyCreate {
	fac.mutation.SetToUserID(i)
	return fac
}

// SetToFaceURL sets the "to_face_url" field.
func (fac *FriendApplyCreate) SetToFaceURL(s string) *FriendApplyCreate {
	fac.mutation.SetToFaceURL(s)
	return fac
}

// SetToNickname sets the "to_nickname" field.
func (fac *FriendApplyCreate) SetToNickname(s string) *FriendApplyCreate {
	fac.mutation.SetToNickname(s)
	return fac
}

// SetResult sets the "result" field.
func (fac *FriendApplyCreate) SetResult(et enums.ApplyType) *FriendApplyCreate {
	fac.mutation.SetResult(et)
	return fac
}

// SetNillableResult sets the "result" field if the given value is not nil.
func (fac *FriendApplyCreate) SetNillableResult(et *enums.ApplyType) *FriendApplyCreate {
	if et != nil {
		fac.SetResult(*et)
	}
	return fac
}

// SetReqMsg sets the "req_msg" field.
func (fac *FriendApplyCreate) SetReqMsg(s string) *FriendApplyCreate {
	fac.mutation.SetReqMsg(s)
	return fac
}

// SetNillableReqMsg sets the "req_msg" field if the given value is not nil.
func (fac *FriendApplyCreate) SetNillableReqMsg(s *string) *FriendApplyCreate {
	if s != nil {
		fac.SetReqMsg(*s)
	}
	return fac
}

// SetGroupID sets the "group_id" field.
func (fac *FriendApplyCreate) SetGroupID(i int64) *FriendApplyCreate {
	fac.mutation.SetGroupID(i)
	return fac
}

// SetID sets the "id" field.
func (fac *FriendApplyCreate) SetID(i int64) *FriendApplyCreate {
	fac.mutation.SetID(i)
	return fac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (fac *FriendApplyCreate) SetNillableID(i *int64) *FriendApplyCreate {
	if i != nil {
		fac.SetID(*i)
	}
	return fac
}

// SetFromUser sets the "from_user" edge to the User entity.
func (fac *FriendApplyCreate) SetFromUser(u *User) *FriendApplyCreate {
	return fac.SetFromUserID(u.ID)
}

// SetToUser sets the "to_user" edge to the User entity.
func (fac *FriendApplyCreate) SetToUser(u *User) *FriendApplyCreate {
	return fac.SetToUserID(u.ID)
}

// Mutation returns the FriendApplyMutation object of the builder.
func (fac *FriendApplyCreate) Mutation() *FriendApplyMutation {
	return fac.mutation
}

// Save creates the FriendApply in the database.
func (fac *FriendApplyCreate) Save(ctx context.Context) (*FriendApply, error) {
	fac.defaults()
	return withHooks(ctx, fac.sqlSave, fac.mutation, fac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (fac *FriendApplyCreate) SaveX(ctx context.Context) *FriendApply {
	v, err := fac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fac *FriendApplyCreate) Exec(ctx context.Context) error {
	_, err := fac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fac *FriendApplyCreate) ExecX(ctx context.Context) {
	if err := fac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fac *FriendApplyCreate) defaults() {
	if _, ok := fac.mutation.CreatedAt(); !ok {
		v := friendapply.DefaultCreatedAt()
		fac.mutation.SetCreatedAt(v)
	}
	if _, ok := fac.mutation.UpdatedAt(); !ok {
		v := friendapply.DefaultUpdatedAt()
		fac.mutation.SetUpdatedAt(v)
	}
	if _, ok := fac.mutation.DeletedAt(); !ok {
		v := friendapply.DefaultDeletedAt
		fac.mutation.SetDeletedAt(v)
	}
	if _, ok := fac.mutation.Result(); !ok {
		v := friendapply.DefaultResult
		fac.mutation.SetResult(v)
	}
	if _, ok := fac.mutation.ReqMsg(); !ok {
		v := friendapply.DefaultReqMsg
		fac.mutation.SetReqMsg(v)
	}
	if _, ok := fac.mutation.ID(); !ok {
		v := friendapply.DefaultID()
		fac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fac *FriendApplyCreate) check() error {
	if _, ok := fac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "FriendApply.created_at"`)}
	}
	if _, ok := fac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "FriendApply.updated_at"`)}
	}
	if _, ok := fac.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "FriendApply.deleted_at"`)}
	}
	if _, ok := fac.mutation.FromUserID(); !ok {
		return &ValidationError{Name: "from_user_id", err: errors.New(`ent: missing required field "FriendApply.from_user_id"`)}
	}
	if _, ok := fac.mutation.FromNickname(); !ok {
		return &ValidationError{Name: "from_nickname", err: errors.New(`ent: missing required field "FriendApply.from_nickname"`)}
	}
	if _, ok := fac.mutation.FromFaceURL(); !ok {
		return &ValidationError{Name: "from_face_url", err: errors.New(`ent: missing required field "FriendApply.from_face_url"`)}
	}
	if _, ok := fac.mutation.ToUserID(); !ok {
		return &ValidationError{Name: "to_user_id", err: errors.New(`ent: missing required field "FriendApply.to_user_id"`)}
	}
	if _, ok := fac.mutation.ToFaceURL(); !ok {
		return &ValidationError{Name: "to_face_url", err: errors.New(`ent: missing required field "FriendApply.to_face_url"`)}
	}
	if _, ok := fac.mutation.ToNickname(); !ok {
		return &ValidationError{Name: "to_nickname", err: errors.New(`ent: missing required field "FriendApply.to_nickname"`)}
	}
	if _, ok := fac.mutation.Result(); !ok {
		return &ValidationError{Name: "result", err: errors.New(`ent: missing required field "FriendApply.result"`)}
	}
	if _, ok := fac.mutation.ReqMsg(); !ok {
		return &ValidationError{Name: "req_msg", err: errors.New(`ent: missing required field "FriendApply.req_msg"`)}
	}
	if _, ok := fac.mutation.GroupID(); !ok {
		return &ValidationError{Name: "group_id", err: errors.New(`ent: missing required field "FriendApply.group_id"`)}
	}
	if _, ok := fac.mutation.FromUserID(); !ok {
		return &ValidationError{Name: "from_user", err: errors.New(`ent: missing required edge "FriendApply.from_user"`)}
	}
	if _, ok := fac.mutation.ToUserID(); !ok {
		return &ValidationError{Name: "to_user", err: errors.New(`ent: missing required edge "FriendApply.to_user"`)}
	}
	return nil
}

func (fac *FriendApplyCreate) sqlSave(ctx context.Context) (*FriendApply, error) {
	if err := fac.check(); err != nil {
		return nil, err
	}
	_node, _spec := fac.createSpec()
	if err := sqlgraph.CreateNode(ctx, fac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	fac.mutation.id = &_node.ID
	fac.mutation.done = true
	return _node, nil
}

func (fac *FriendApplyCreate) createSpec() (*FriendApply, *sqlgraph.CreateSpec) {
	var (
		_node = &FriendApply{config: fac.config}
		_spec = sqlgraph.NewCreateSpec(friendapply.Table, sqlgraph.NewFieldSpec(friendapply.FieldID, field.TypeInt64))
	)
	if id, ok := fac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := fac.mutation.CreatedAt(); ok {
		_spec.SetField(friendapply.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := fac.mutation.UpdatedAt(); ok {
		_spec.SetField(friendapply.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := fac.mutation.DeletedAt(); ok {
		_spec.SetField(friendapply.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := fac.mutation.FromNickname(); ok {
		_spec.SetField(friendapply.FieldFromNickname, field.TypeString, value)
		_node.FromNickname = value
	}
	if value, ok := fac.mutation.FromFaceURL(); ok {
		_spec.SetField(friendapply.FieldFromFaceURL, field.TypeString, value)
		_node.FromFaceURL = value
	}
	if value, ok := fac.mutation.ToFaceURL(); ok {
		_spec.SetField(friendapply.FieldToFaceURL, field.TypeString, value)
		_node.ToFaceURL = value
	}
	if value, ok := fac.mutation.ToNickname(); ok {
		_spec.SetField(friendapply.FieldToNickname, field.TypeString, value)
		_node.ToNickname = value
	}
	if value, ok := fac.mutation.Result(); ok {
		_spec.SetField(friendapply.FieldResult, field.TypeString, value)
		_node.Result = value
	}
	if value, ok := fac.mutation.ReqMsg(); ok {
		_spec.SetField(friendapply.FieldReqMsg, field.TypeString, value)
		_node.ReqMsg = value
	}
	if value, ok := fac.mutation.GroupID(); ok {
		_spec.SetField(friendapply.FieldGroupID, field.TypeInt64, value)
		_node.GroupID = value
	}
	if nodes := fac.mutation.FromUserIDs(); len(nodes) > 0 {
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
		_node.FromUserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fac.mutation.ToUserIDs(); len(nodes) > 0 {
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
		_node.ToUserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// FriendApplyCreateBulk is the builder for creating many FriendApply entities in bulk.
type FriendApplyCreateBulk struct {
	config
	err      error
	builders []*FriendApplyCreate
}

// Save creates the FriendApply entities in the database.
func (facb *FriendApplyCreateBulk) Save(ctx context.Context) ([]*FriendApply, error) {
	if facb.err != nil {
		return nil, facb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(facb.builders))
	nodes := make([]*FriendApply, len(facb.builders))
	mutators := make([]Mutator, len(facb.builders))
	for i := range facb.builders {
		func(i int, root context.Context) {
			builder := facb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FriendApplyMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, facb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, facb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, facb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (facb *FriendApplyCreateBulk) SaveX(ctx context.Context) []*FriendApply {
	v, err := facb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (facb *FriendApplyCreateBulk) Exec(ctx context.Context) error {
	_, err := facb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (facb *FriendApplyCreateBulk) ExecX(ctx context.Context) {
	if err := facb.Exec(ctx); err != nil {
		panic(err)
	}
}
