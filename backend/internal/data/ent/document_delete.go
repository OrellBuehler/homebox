// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/thechosenlan/homebox/backend/internal/data/ent/document"
	"github.com/thechosenlan/homebox/backend/internal/data/ent/predicate"
)

// DocumentDelete is the builder for deleting a Document entity.
type DocumentDelete struct {
	config
	hooks    []Hook
	mutation *DocumentMutation
}

// Where appends a list predicates to the DocumentDelete builder.
func (dd *DocumentDelete) Where(ps ...predicate.Document) *DocumentDelete {
	dd.mutation.Where(ps...)
	return dd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dd *DocumentDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, DocumentMutation](ctx, dd.sqlExec, dd.mutation, dd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (dd *DocumentDelete) ExecX(ctx context.Context) int {
	n, err := dd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dd *DocumentDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: document.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: document.FieldID,
			},
		},
	}
	if ps := dd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, dd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	dd.mutation.done = true
	return affected, err
}

// DocumentDeleteOne is the builder for deleting a single Document entity.
type DocumentDeleteOne struct {
	dd *DocumentDelete
}

// Exec executes the deletion query.
func (ddo *DocumentDeleteOne) Exec(ctx context.Context) error {
	n, err := ddo.dd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{document.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ddo *DocumentDeleteOne) ExecX(ctx context.Context) {
	ddo.dd.ExecX(ctx)
}
