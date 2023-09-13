// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kubecit-service/ent/setting"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SettingCreate is the builder for creating a Setting entity.
type SettingCreate struct {
	config
	mutation *SettingMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (sc *SettingCreate) SetName(s string) *SettingCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetDetail sets the "detail" field.
func (sc *SettingCreate) SetDetail(s string) *SettingCreate {
	sc.mutation.SetDetail(s)
	return sc
}

// SetCover sets the "cover" field.
func (sc *SettingCreate) SetCover(s string) *SettingCreate {
	sc.mutation.SetCover(s)
	return sc
}

// Mutation returns the SettingMutation object of the builder.
func (sc *SettingCreate) Mutation() *SettingMutation {
	return sc.mutation
}

// Save creates the Setting in the database.
func (sc *SettingCreate) Save(ctx context.Context) (*Setting, error) {
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SettingCreate) SaveX(ctx context.Context) *Setting {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SettingCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SettingCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SettingCreate) check() error {
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Setting.name"`)}
	}
	if _, ok := sc.mutation.Detail(); !ok {
		return &ValidationError{Name: "detail", err: errors.New(`ent: missing required field "Setting.detail"`)}
	}
	if _, ok := sc.mutation.Cover(); !ok {
		return &ValidationError{Name: "cover", err: errors.New(`ent: missing required field "Setting.cover"`)}
	}
	return nil
}

func (sc *SettingCreate) sqlSave(ctx context.Context) (*Setting, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SettingCreate) createSpec() (*Setting, *sqlgraph.CreateSpec) {
	var (
		_node = &Setting{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(setting.Table, sqlgraph.NewFieldSpec(setting.FieldID, field.TypeInt))
	)
	if value, ok := sc.mutation.Name(); ok {
		_spec.SetField(setting.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := sc.mutation.Detail(); ok {
		_spec.SetField(setting.FieldDetail, field.TypeString, value)
		_node.Detail = value
	}
	if value, ok := sc.mutation.Cover(); ok {
		_spec.SetField(setting.FieldCover, field.TypeString, value)
		_node.Cover = value
	}
	return _node, _spec
}

// SettingCreateBulk is the builder for creating many Setting entities in bulk.
type SettingCreateBulk struct {
	config
	builders []*SettingCreate
}

// Save creates the Setting entities in the database.
func (scb *SettingCreateBulk) Save(ctx context.Context) ([]*Setting, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Setting, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SettingMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SettingCreateBulk) SaveX(ctx context.Context) []*Setting {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SettingCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SettingCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
