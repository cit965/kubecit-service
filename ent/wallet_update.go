// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kubecit-service/ent/predicate"
	"kubecit-service/ent/wallet"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WalletUpdate is the builder for updating Wallet entities.
type WalletUpdate struct {
	config
	hooks    []Hook
	mutation *WalletMutation
}

// Where appends a list predicates to the WalletUpdate builder.
func (wu *WalletUpdate) Where(ps ...predicate.Wallet) *WalletUpdate {
	wu.mutation.Where(ps...)
	return wu
}

// SetGoldLeaf sets the "gold_leaf" field.
func (wu *WalletUpdate) SetGoldLeaf(i int32) *WalletUpdate {
	wu.mutation.ResetGoldLeaf()
	wu.mutation.SetGoldLeaf(i)
	return wu
}

// SetNillableGoldLeaf sets the "gold_leaf" field if the given value is not nil.
func (wu *WalletUpdate) SetNillableGoldLeaf(i *int32) *WalletUpdate {
	if i != nil {
		wu.SetGoldLeaf(*i)
	}
	return wu
}

// AddGoldLeaf adds i to the "gold_leaf" field.
func (wu *WalletUpdate) AddGoldLeaf(i int32) *WalletUpdate {
	wu.mutation.AddGoldLeaf(i)
	return wu
}

// ClearGoldLeaf clears the value of the "gold_leaf" field.
func (wu *WalletUpdate) ClearGoldLeaf() *WalletUpdate {
	wu.mutation.ClearGoldLeaf()
	return wu
}

// SetSilverLeaf sets the "silver_leaf" field.
func (wu *WalletUpdate) SetSilverLeaf(i int32) *WalletUpdate {
	wu.mutation.ResetSilverLeaf()
	wu.mutation.SetSilverLeaf(i)
	return wu
}

// SetNillableSilverLeaf sets the "silver_leaf" field if the given value is not nil.
func (wu *WalletUpdate) SetNillableSilverLeaf(i *int32) *WalletUpdate {
	if i != nil {
		wu.SetSilverLeaf(*i)
	}
	return wu
}

// AddSilverLeaf adds i to the "silver_leaf" field.
func (wu *WalletUpdate) AddSilverLeaf(i int32) *WalletUpdate {
	wu.mutation.AddSilverLeaf(i)
	return wu
}

// ClearSilverLeaf clears the value of the "silver_leaf" field.
func (wu *WalletUpdate) ClearSilverLeaf() *WalletUpdate {
	wu.mutation.ClearSilverLeaf()
	return wu
}

// SetFrozenGoldLeaf sets the "frozen_gold_leaf" field.
func (wu *WalletUpdate) SetFrozenGoldLeaf(i int32) *WalletUpdate {
	wu.mutation.ResetFrozenGoldLeaf()
	wu.mutation.SetFrozenGoldLeaf(i)
	return wu
}

// SetNillableFrozenGoldLeaf sets the "frozen_gold_leaf" field if the given value is not nil.
func (wu *WalletUpdate) SetNillableFrozenGoldLeaf(i *int32) *WalletUpdate {
	if i != nil {
		wu.SetFrozenGoldLeaf(*i)
	}
	return wu
}

// AddFrozenGoldLeaf adds i to the "frozen_gold_leaf" field.
func (wu *WalletUpdate) AddFrozenGoldLeaf(i int32) *WalletUpdate {
	wu.mutation.AddFrozenGoldLeaf(i)
	return wu
}

// ClearFrozenGoldLeaf clears the value of the "frozen_gold_leaf" field.
func (wu *WalletUpdate) ClearFrozenGoldLeaf() *WalletUpdate {
	wu.mutation.ClearFrozenGoldLeaf()
	return wu
}

// SetFrozenSilverLeaf sets the "frozen_silver_leaf" field.
func (wu *WalletUpdate) SetFrozenSilverLeaf(i int32) *WalletUpdate {
	wu.mutation.ResetFrozenSilverLeaf()
	wu.mutation.SetFrozenSilverLeaf(i)
	return wu
}

// SetNillableFrozenSilverLeaf sets the "frozen_silver_leaf" field if the given value is not nil.
func (wu *WalletUpdate) SetNillableFrozenSilverLeaf(i *int32) *WalletUpdate {
	if i != nil {
		wu.SetFrozenSilverLeaf(*i)
	}
	return wu
}

// AddFrozenSilverLeaf adds i to the "frozen_silver_leaf" field.
func (wu *WalletUpdate) AddFrozenSilverLeaf(i int32) *WalletUpdate {
	wu.mutation.AddFrozenSilverLeaf(i)
	return wu
}

// ClearFrozenSilverLeaf clears the value of the "frozen_silver_leaf" field.
func (wu *WalletUpdate) ClearFrozenSilverLeaf() *WalletUpdate {
	wu.mutation.ClearFrozenSilverLeaf()
	return wu
}

// SetUserID sets the "user_id" field.
func (wu *WalletUpdate) SetUserID(i int32) *WalletUpdate {
	wu.mutation.ResetUserID()
	wu.mutation.SetUserID(i)
	return wu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (wu *WalletUpdate) SetNillableUserID(i *int32) *WalletUpdate {
	if i != nil {
		wu.SetUserID(*i)
	}
	return wu
}

// AddUserID adds i to the "user_id" field.
func (wu *WalletUpdate) AddUserID(i int32) *WalletUpdate {
	wu.mutation.AddUserID(i)
	return wu
}

// ClearUserID clears the value of the "user_id" field.
func (wu *WalletUpdate) ClearUserID() *WalletUpdate {
	wu.mutation.ClearUserID()
	return wu
}

// SetUsername sets the "username" field.
func (wu *WalletUpdate) SetUsername(s string) *WalletUpdate {
	wu.mutation.SetUsername(s)
	return wu
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (wu *WalletUpdate) SetNillableUsername(s *string) *WalletUpdate {
	if s != nil {
		wu.SetUsername(*s)
	}
	return wu
}

// ClearUsername clears the value of the "username" field.
func (wu *WalletUpdate) ClearUsername() *WalletUpdate {
	wu.mutation.ClearUsername()
	return wu
}

// SetCreatedAt sets the "created_at" field.
func (wu *WalletUpdate) SetCreatedAt(t time.Time) *WalletUpdate {
	wu.mutation.SetCreatedAt(t)
	return wu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (wu *WalletUpdate) SetNillableCreatedAt(t *time.Time) *WalletUpdate {
	if t != nil {
		wu.SetCreatedAt(*t)
	}
	return wu
}

// SetUpdatedAt sets the "updated_at" field.
func (wu *WalletUpdate) SetUpdatedAt(t time.Time) *WalletUpdate {
	wu.mutation.SetUpdatedAt(t)
	return wu
}

// Mutation returns the WalletMutation object of the builder.
func (wu *WalletUpdate) Mutation() *WalletMutation {
	return wu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wu *WalletUpdate) Save(ctx context.Context) (int, error) {
	wu.defaults()
	return withHooks(ctx, wu.sqlSave, wu.mutation, wu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wu *WalletUpdate) SaveX(ctx context.Context) int {
	affected, err := wu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wu *WalletUpdate) Exec(ctx context.Context) error {
	_, err := wu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wu *WalletUpdate) ExecX(ctx context.Context) {
	if err := wu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wu *WalletUpdate) defaults() {
	if _, ok := wu.mutation.UpdatedAt(); !ok {
		v := wallet.UpdateDefaultUpdatedAt()
		wu.mutation.SetUpdatedAt(v)
	}
}

func (wu *WalletUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(wallet.Table, wallet.Columns, sqlgraph.NewFieldSpec(wallet.FieldID, field.TypeInt))
	if ps := wu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wu.mutation.GoldLeaf(); ok {
		_spec.SetField(wallet.FieldGoldLeaf, field.TypeInt32, value)
	}
	if value, ok := wu.mutation.AddedGoldLeaf(); ok {
		_spec.AddField(wallet.FieldGoldLeaf, field.TypeInt32, value)
	}
	if wu.mutation.GoldLeafCleared() {
		_spec.ClearField(wallet.FieldGoldLeaf, field.TypeInt32)
	}
	if value, ok := wu.mutation.SilverLeaf(); ok {
		_spec.SetField(wallet.FieldSilverLeaf, field.TypeInt32, value)
	}
	if value, ok := wu.mutation.AddedSilverLeaf(); ok {
		_spec.AddField(wallet.FieldSilverLeaf, field.TypeInt32, value)
	}
	if wu.mutation.SilverLeafCleared() {
		_spec.ClearField(wallet.FieldSilverLeaf, field.TypeInt32)
	}
	if value, ok := wu.mutation.FrozenGoldLeaf(); ok {
		_spec.SetField(wallet.FieldFrozenGoldLeaf, field.TypeInt32, value)
	}
	if value, ok := wu.mutation.AddedFrozenGoldLeaf(); ok {
		_spec.AddField(wallet.FieldFrozenGoldLeaf, field.TypeInt32, value)
	}
	if wu.mutation.FrozenGoldLeafCleared() {
		_spec.ClearField(wallet.FieldFrozenGoldLeaf, field.TypeInt32)
	}
	if value, ok := wu.mutation.FrozenSilverLeaf(); ok {
		_spec.SetField(wallet.FieldFrozenSilverLeaf, field.TypeInt32, value)
	}
	if value, ok := wu.mutation.AddedFrozenSilverLeaf(); ok {
		_spec.AddField(wallet.FieldFrozenSilverLeaf, field.TypeInt32, value)
	}
	if wu.mutation.FrozenSilverLeafCleared() {
		_spec.ClearField(wallet.FieldFrozenSilverLeaf, field.TypeInt32)
	}
	if value, ok := wu.mutation.UserID(); ok {
		_spec.SetField(wallet.FieldUserID, field.TypeInt32, value)
	}
	if value, ok := wu.mutation.AddedUserID(); ok {
		_spec.AddField(wallet.FieldUserID, field.TypeInt32, value)
	}
	if wu.mutation.UserIDCleared() {
		_spec.ClearField(wallet.FieldUserID, field.TypeInt32)
	}
	if value, ok := wu.mutation.Username(); ok {
		_spec.SetField(wallet.FieldUsername, field.TypeString, value)
	}
	if wu.mutation.UsernameCleared() {
		_spec.ClearField(wallet.FieldUsername, field.TypeString)
	}
	if value, ok := wu.mutation.CreatedAt(); ok {
		_spec.SetField(wallet.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := wu.mutation.UpdatedAt(); ok {
		_spec.SetField(wallet.FieldUpdatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{wallet.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	wu.mutation.done = true
	return n, nil
}

// WalletUpdateOne is the builder for updating a single Wallet entity.
type WalletUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WalletMutation
}

// SetGoldLeaf sets the "gold_leaf" field.
func (wuo *WalletUpdateOne) SetGoldLeaf(i int32) *WalletUpdateOne {
	wuo.mutation.ResetGoldLeaf()
	wuo.mutation.SetGoldLeaf(i)
	return wuo
}

// SetNillableGoldLeaf sets the "gold_leaf" field if the given value is not nil.
func (wuo *WalletUpdateOne) SetNillableGoldLeaf(i *int32) *WalletUpdateOne {
	if i != nil {
		wuo.SetGoldLeaf(*i)
	}
	return wuo
}

// AddGoldLeaf adds i to the "gold_leaf" field.
func (wuo *WalletUpdateOne) AddGoldLeaf(i int32) *WalletUpdateOne {
	wuo.mutation.AddGoldLeaf(i)
	return wuo
}

// ClearGoldLeaf clears the value of the "gold_leaf" field.
func (wuo *WalletUpdateOne) ClearGoldLeaf() *WalletUpdateOne {
	wuo.mutation.ClearGoldLeaf()
	return wuo
}

// SetSilverLeaf sets the "silver_leaf" field.
func (wuo *WalletUpdateOne) SetSilverLeaf(i int32) *WalletUpdateOne {
	wuo.mutation.ResetSilverLeaf()
	wuo.mutation.SetSilverLeaf(i)
	return wuo
}

// SetNillableSilverLeaf sets the "silver_leaf" field if the given value is not nil.
func (wuo *WalletUpdateOne) SetNillableSilverLeaf(i *int32) *WalletUpdateOne {
	if i != nil {
		wuo.SetSilverLeaf(*i)
	}
	return wuo
}

// AddSilverLeaf adds i to the "silver_leaf" field.
func (wuo *WalletUpdateOne) AddSilverLeaf(i int32) *WalletUpdateOne {
	wuo.mutation.AddSilverLeaf(i)
	return wuo
}

// ClearSilverLeaf clears the value of the "silver_leaf" field.
func (wuo *WalletUpdateOne) ClearSilverLeaf() *WalletUpdateOne {
	wuo.mutation.ClearSilverLeaf()
	return wuo
}

// SetFrozenGoldLeaf sets the "frozen_gold_leaf" field.
func (wuo *WalletUpdateOne) SetFrozenGoldLeaf(i int32) *WalletUpdateOne {
	wuo.mutation.ResetFrozenGoldLeaf()
	wuo.mutation.SetFrozenGoldLeaf(i)
	return wuo
}

// SetNillableFrozenGoldLeaf sets the "frozen_gold_leaf" field if the given value is not nil.
func (wuo *WalletUpdateOne) SetNillableFrozenGoldLeaf(i *int32) *WalletUpdateOne {
	if i != nil {
		wuo.SetFrozenGoldLeaf(*i)
	}
	return wuo
}

// AddFrozenGoldLeaf adds i to the "frozen_gold_leaf" field.
func (wuo *WalletUpdateOne) AddFrozenGoldLeaf(i int32) *WalletUpdateOne {
	wuo.mutation.AddFrozenGoldLeaf(i)
	return wuo
}

// ClearFrozenGoldLeaf clears the value of the "frozen_gold_leaf" field.
func (wuo *WalletUpdateOne) ClearFrozenGoldLeaf() *WalletUpdateOne {
	wuo.mutation.ClearFrozenGoldLeaf()
	return wuo
}

// SetFrozenSilverLeaf sets the "frozen_silver_leaf" field.
func (wuo *WalletUpdateOne) SetFrozenSilverLeaf(i int32) *WalletUpdateOne {
	wuo.mutation.ResetFrozenSilverLeaf()
	wuo.mutation.SetFrozenSilverLeaf(i)
	return wuo
}

// SetNillableFrozenSilverLeaf sets the "frozen_silver_leaf" field if the given value is not nil.
func (wuo *WalletUpdateOne) SetNillableFrozenSilverLeaf(i *int32) *WalletUpdateOne {
	if i != nil {
		wuo.SetFrozenSilverLeaf(*i)
	}
	return wuo
}

// AddFrozenSilverLeaf adds i to the "frozen_silver_leaf" field.
func (wuo *WalletUpdateOne) AddFrozenSilverLeaf(i int32) *WalletUpdateOne {
	wuo.mutation.AddFrozenSilverLeaf(i)
	return wuo
}

// ClearFrozenSilverLeaf clears the value of the "frozen_silver_leaf" field.
func (wuo *WalletUpdateOne) ClearFrozenSilverLeaf() *WalletUpdateOne {
	wuo.mutation.ClearFrozenSilverLeaf()
	return wuo
}

// SetUserID sets the "user_id" field.
func (wuo *WalletUpdateOne) SetUserID(i int32) *WalletUpdateOne {
	wuo.mutation.ResetUserID()
	wuo.mutation.SetUserID(i)
	return wuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (wuo *WalletUpdateOne) SetNillableUserID(i *int32) *WalletUpdateOne {
	if i != nil {
		wuo.SetUserID(*i)
	}
	return wuo
}

// AddUserID adds i to the "user_id" field.
func (wuo *WalletUpdateOne) AddUserID(i int32) *WalletUpdateOne {
	wuo.mutation.AddUserID(i)
	return wuo
}

// ClearUserID clears the value of the "user_id" field.
func (wuo *WalletUpdateOne) ClearUserID() *WalletUpdateOne {
	wuo.mutation.ClearUserID()
	return wuo
}

// SetUsername sets the "username" field.
func (wuo *WalletUpdateOne) SetUsername(s string) *WalletUpdateOne {
	wuo.mutation.SetUsername(s)
	return wuo
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (wuo *WalletUpdateOne) SetNillableUsername(s *string) *WalletUpdateOne {
	if s != nil {
		wuo.SetUsername(*s)
	}
	return wuo
}

// ClearUsername clears the value of the "username" field.
func (wuo *WalletUpdateOne) ClearUsername() *WalletUpdateOne {
	wuo.mutation.ClearUsername()
	return wuo
}

// SetCreatedAt sets the "created_at" field.
func (wuo *WalletUpdateOne) SetCreatedAt(t time.Time) *WalletUpdateOne {
	wuo.mutation.SetCreatedAt(t)
	return wuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (wuo *WalletUpdateOne) SetNillableCreatedAt(t *time.Time) *WalletUpdateOne {
	if t != nil {
		wuo.SetCreatedAt(*t)
	}
	return wuo
}

// SetUpdatedAt sets the "updated_at" field.
func (wuo *WalletUpdateOne) SetUpdatedAt(t time.Time) *WalletUpdateOne {
	wuo.mutation.SetUpdatedAt(t)
	return wuo
}

// Mutation returns the WalletMutation object of the builder.
func (wuo *WalletUpdateOne) Mutation() *WalletMutation {
	return wuo.mutation
}

// Where appends a list predicates to the WalletUpdate builder.
func (wuo *WalletUpdateOne) Where(ps ...predicate.Wallet) *WalletUpdateOne {
	wuo.mutation.Where(ps...)
	return wuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wuo *WalletUpdateOne) Select(field string, fields ...string) *WalletUpdateOne {
	wuo.fields = append([]string{field}, fields...)
	return wuo
}

// Save executes the query and returns the updated Wallet entity.
func (wuo *WalletUpdateOne) Save(ctx context.Context) (*Wallet, error) {
	wuo.defaults()
	return withHooks(ctx, wuo.sqlSave, wuo.mutation, wuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wuo *WalletUpdateOne) SaveX(ctx context.Context) *Wallet {
	node, err := wuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wuo *WalletUpdateOne) Exec(ctx context.Context) error {
	_, err := wuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wuo *WalletUpdateOne) ExecX(ctx context.Context) {
	if err := wuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wuo *WalletUpdateOne) defaults() {
	if _, ok := wuo.mutation.UpdatedAt(); !ok {
		v := wallet.UpdateDefaultUpdatedAt()
		wuo.mutation.SetUpdatedAt(v)
	}
}

func (wuo *WalletUpdateOne) sqlSave(ctx context.Context) (_node *Wallet, err error) {
	_spec := sqlgraph.NewUpdateSpec(wallet.Table, wallet.Columns, sqlgraph.NewFieldSpec(wallet.FieldID, field.TypeInt))
	id, ok := wuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Wallet.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, wallet.FieldID)
		for _, f := range fields {
			if !wallet.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != wallet.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wuo.mutation.GoldLeaf(); ok {
		_spec.SetField(wallet.FieldGoldLeaf, field.TypeInt32, value)
	}
	if value, ok := wuo.mutation.AddedGoldLeaf(); ok {
		_spec.AddField(wallet.FieldGoldLeaf, field.TypeInt32, value)
	}
	if wuo.mutation.GoldLeafCleared() {
		_spec.ClearField(wallet.FieldGoldLeaf, field.TypeInt32)
	}
	if value, ok := wuo.mutation.SilverLeaf(); ok {
		_spec.SetField(wallet.FieldSilverLeaf, field.TypeInt32, value)
	}
	if value, ok := wuo.mutation.AddedSilverLeaf(); ok {
		_spec.AddField(wallet.FieldSilverLeaf, field.TypeInt32, value)
	}
	if wuo.mutation.SilverLeafCleared() {
		_spec.ClearField(wallet.FieldSilverLeaf, field.TypeInt32)
	}
	if value, ok := wuo.mutation.FrozenGoldLeaf(); ok {
		_spec.SetField(wallet.FieldFrozenGoldLeaf, field.TypeInt32, value)
	}
	if value, ok := wuo.mutation.AddedFrozenGoldLeaf(); ok {
		_spec.AddField(wallet.FieldFrozenGoldLeaf, field.TypeInt32, value)
	}
	if wuo.mutation.FrozenGoldLeafCleared() {
		_spec.ClearField(wallet.FieldFrozenGoldLeaf, field.TypeInt32)
	}
	if value, ok := wuo.mutation.FrozenSilverLeaf(); ok {
		_spec.SetField(wallet.FieldFrozenSilverLeaf, field.TypeInt32, value)
	}
	if value, ok := wuo.mutation.AddedFrozenSilverLeaf(); ok {
		_spec.AddField(wallet.FieldFrozenSilverLeaf, field.TypeInt32, value)
	}
	if wuo.mutation.FrozenSilverLeafCleared() {
		_spec.ClearField(wallet.FieldFrozenSilverLeaf, field.TypeInt32)
	}
	if value, ok := wuo.mutation.UserID(); ok {
		_spec.SetField(wallet.FieldUserID, field.TypeInt32, value)
	}
	if value, ok := wuo.mutation.AddedUserID(); ok {
		_spec.AddField(wallet.FieldUserID, field.TypeInt32, value)
	}
	if wuo.mutation.UserIDCleared() {
		_spec.ClearField(wallet.FieldUserID, field.TypeInt32)
	}
	if value, ok := wuo.mutation.Username(); ok {
		_spec.SetField(wallet.FieldUsername, field.TypeString, value)
	}
	if wuo.mutation.UsernameCleared() {
		_spec.ClearField(wallet.FieldUsername, field.TypeString)
	}
	if value, ok := wuo.mutation.CreatedAt(); ok {
		_spec.SetField(wallet.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := wuo.mutation.UpdatedAt(); ok {
		_spec.SetField(wallet.FieldUpdatedAt, field.TypeTime, value)
	}
	_node = &Wallet{config: wuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{wallet.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	wuo.mutation.done = true
	return _node, nil
}
