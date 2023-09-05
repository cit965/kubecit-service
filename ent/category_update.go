// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kubecit-service/ent/category"
	"kubecit-service/ent/course"
	"kubecit-service/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CategoryUpdate is the builder for updating Category entities.
type CategoryUpdate struct {
	config
	hooks    []Hook
	mutation *CategoryMutation
}

// Where appends a list predicates to the CategoryUpdate builder.
func (cu *CategoryUpdate) Where(ps ...predicate.Category) *CategoryUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetName sets the "name" field.
func (cu *CategoryUpdate) SetName(s string) *CategoryUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cu *CategoryUpdate) SetNillableName(s *string) *CategoryUpdate {
	if s != nil {
		cu.SetName(*s)
	}
	return cu
}

// SetParentId sets the "parentId" field.
func (cu *CategoryUpdate) SetParentId(s string) *CategoryUpdate {
	cu.mutation.SetParentId(s)
	return cu
}

// SetLevel sets the "level" field.
func (cu *CategoryUpdate) SetLevel(s string) *CategoryUpdate {
	cu.mutation.SetLevel(s)
	return cu
}

// SetStatus sets the "status" field.
func (cu *CategoryUpdate) SetStatus(s string) *CategoryUpdate {
	cu.mutation.SetStatus(s)
	return cu
}

// AddCourseIDs adds the "course" edge to the Course entity by IDs.
func (cu *CategoryUpdate) AddCourseIDs(ids ...string) *CategoryUpdate {
	cu.mutation.AddCourseIDs(ids...)
	return cu
}

// AddCourse adds the "course" edges to the Course entity.
func (cu *CategoryUpdate) AddCourse(c ...*Course) *CategoryUpdate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.AddCourseIDs(ids...)
}

// Mutation returns the CategoryMutation object of the builder.
func (cu *CategoryUpdate) Mutation() *CategoryMutation {
	return cu.mutation
}

// ClearCourse clears all "course" edges to the Course entity.
func (cu *CategoryUpdate) ClearCourse() *CategoryUpdate {
	cu.mutation.ClearCourse()
	return cu
}

// RemoveCourseIDs removes the "course" edge to Course entities by IDs.
func (cu *CategoryUpdate) RemoveCourseIDs(ids ...string) *CategoryUpdate {
	cu.mutation.RemoveCourseIDs(ids...)
	return cu
}

// RemoveCourse removes "course" edges to Course entities.
func (cu *CategoryUpdate) RemoveCourse(c ...*Course) *CategoryUpdate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.RemoveCourseIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CategoryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CategoryUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CategoryUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CategoryUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CategoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(category.Table, category.Columns, sqlgraph.NewFieldSpec(category.FieldID, field.TypeString))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(category.FieldName, field.TypeString, value)
	}
	if value, ok := cu.mutation.ParentId(); ok {
		_spec.SetField(category.FieldParentId, field.TypeString, value)
	}
	if value, ok := cu.mutation.Level(); ok {
		_spec.SetField(category.FieldLevel, field.TypeString, value)
	}
	if value, ok := cu.mutation.Status(); ok {
		_spec.SetField(category.FieldStatus, field.TypeString, value)
	}
	if cu.mutation.CourseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   category.CourseTable,
			Columns: category.CoursePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(course.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedCourseIDs(); len(nodes) > 0 && !cu.mutation.CourseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   category.CourseTable,
			Columns: category.CoursePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(course.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.CourseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   category.CourseTable,
			Columns: category.CoursePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(course.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{category.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CategoryUpdateOne is the builder for updating a single Category entity.
type CategoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CategoryMutation
}

// SetName sets the "name" field.
func (cuo *CategoryUpdateOne) SetName(s string) *CategoryUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cuo *CategoryUpdateOne) SetNillableName(s *string) *CategoryUpdateOne {
	if s != nil {
		cuo.SetName(*s)
	}
	return cuo
}

// SetParentId sets the "parentId" field.
func (cuo *CategoryUpdateOne) SetParentId(s string) *CategoryUpdateOne {
	cuo.mutation.SetParentId(s)
	return cuo
}

// SetLevel sets the "level" field.
func (cuo *CategoryUpdateOne) SetLevel(s string) *CategoryUpdateOne {
	cuo.mutation.SetLevel(s)
	return cuo
}

// SetStatus sets the "status" field.
func (cuo *CategoryUpdateOne) SetStatus(s string) *CategoryUpdateOne {
	cuo.mutation.SetStatus(s)
	return cuo
}

// AddCourseIDs adds the "course" edge to the Course entity by IDs.
func (cuo *CategoryUpdateOne) AddCourseIDs(ids ...string) *CategoryUpdateOne {
	cuo.mutation.AddCourseIDs(ids...)
	return cuo
}

// AddCourse adds the "course" edges to the Course entity.
func (cuo *CategoryUpdateOne) AddCourse(c ...*Course) *CategoryUpdateOne {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.AddCourseIDs(ids...)
}

// Mutation returns the CategoryMutation object of the builder.
func (cuo *CategoryUpdateOne) Mutation() *CategoryMutation {
	return cuo.mutation
}

// ClearCourse clears all "course" edges to the Course entity.
func (cuo *CategoryUpdateOne) ClearCourse() *CategoryUpdateOne {
	cuo.mutation.ClearCourse()
	return cuo
}

// RemoveCourseIDs removes the "course" edge to Course entities by IDs.
func (cuo *CategoryUpdateOne) RemoveCourseIDs(ids ...string) *CategoryUpdateOne {
	cuo.mutation.RemoveCourseIDs(ids...)
	return cuo
}

// RemoveCourse removes "course" edges to Course entities.
func (cuo *CategoryUpdateOne) RemoveCourse(c ...*Course) *CategoryUpdateOne {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.RemoveCourseIDs(ids...)
}

// Where appends a list predicates to the CategoryUpdate builder.
func (cuo *CategoryUpdateOne) Where(ps ...predicate.Category) *CategoryUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CategoryUpdateOne) Select(field string, fields ...string) *CategoryUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Category entity.
func (cuo *CategoryUpdateOne) Save(ctx context.Context) (*Category, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CategoryUpdateOne) SaveX(ctx context.Context) *Category {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CategoryUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CategoryUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CategoryUpdateOne) sqlSave(ctx context.Context) (_node *Category, err error) {
	_spec := sqlgraph.NewUpdateSpec(category.Table, category.Columns, sqlgraph.NewFieldSpec(category.FieldID, field.TypeString))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Category.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, category.FieldID)
		for _, f := range fields {
			if !category.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != category.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(category.FieldName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.ParentId(); ok {
		_spec.SetField(category.FieldParentId, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Level(); ok {
		_spec.SetField(category.FieldLevel, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Status(); ok {
		_spec.SetField(category.FieldStatus, field.TypeString, value)
	}
	if cuo.mutation.CourseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   category.CourseTable,
			Columns: category.CoursePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(course.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedCourseIDs(); len(nodes) > 0 && !cuo.mutation.CourseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   category.CourseTable,
			Columns: category.CoursePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(course.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.CourseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   category.CourseTable,
			Columns: category.CoursePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(course.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Category{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{category.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
