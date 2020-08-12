// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/dreamvo/gilfoyle/ent/predicate"
	"github.com/dreamvo/gilfoyle/ent/video"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// VideoDelete is the builder for deleting a Video entity.
type VideoDelete struct {
	config
	hooks      []Hook
	mutation   *VideoMutation
	predicates []predicate.Video
}

// Where adds a new predicate to the delete builder.
func (vd *VideoDelete) Where(ps ...predicate.Video) *VideoDelete {
	vd.predicates = append(vd.predicates, ps...)
	return vd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (vd *VideoDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(vd.hooks) == 0 {
		affected, err = vd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VideoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			vd.mutation = mutation
			affected, err = vd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(vd.hooks) - 1; i >= 0; i-- {
			mut = vd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (vd *VideoDelete) ExecX(ctx context.Context) int {
	n, err := vd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (vd *VideoDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: video.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: video.FieldID,
			},
		},
	}
	if ps := vd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, vd.driver, _spec)
}

// VideoDeleteOne is the builder for deleting a single Video entity.
type VideoDeleteOne struct {
	vd *VideoDelete
}

// Exec executes the deletion query.
func (vdo *VideoDeleteOne) Exec(ctx context.Context) error {
	n, err := vdo.vd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{video.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (vdo *VideoDeleteOne) ExecX(ctx context.Context) {
	vdo.vd.ExecX(ctx)
}
