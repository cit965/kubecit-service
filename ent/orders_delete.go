// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"kubecit-service/ent/orders"
	"kubecit-service/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OrdersDelete is the builder for deleting a Orders entity.
type OrdersDelete struct {
	config
	hooks    []Hook
	mutation *OrdersMutation
}

// Where appends a list predicates to the OrdersDelete builder.
func (od *OrdersDelete) Where(ps ...predicate.Orders) *OrdersDelete {
	od.mutation.Where(ps...)
	return od
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (od *OrdersDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, od.sqlExec, od.mutation, od.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (od *OrdersDelete) ExecX(ctx context.Context) int {
	n, err := od.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (od *OrdersDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(orders.Table, sqlgraph.NewFieldSpec(orders.FieldID, field.TypeInt))
	if ps := od.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, od.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	od.mutation.done = true
	return affected, err
}

// OrdersDeleteOne is the builder for deleting a single Orders entity.
type OrdersDeleteOne struct {
	od *OrdersDelete
}

// Where appends a list predicates to the OrdersDelete builder.
func (odo *OrdersDeleteOne) Where(ps ...predicate.Orders) *OrdersDeleteOne {
	odo.od.mutation.Where(ps...)
	return odo
}

// Exec executes the deletion query.
func (odo *OrdersDeleteOne) Exec(ctx context.Context) error {
	n, err := odo.od.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{orders.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (odo *OrdersDeleteOne) ExecX(ctx context.Context) {
	if err := odo.Exec(ctx); err != nil {
		panic(err)
	}
}