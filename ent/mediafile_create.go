// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dreamvo/gilfoyle/ent/mediafile"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// MediaFileCreate is the builder for creating a MediaFile entity.
type MediaFileCreate struct {
	config
	mutation *MediaFileMutation
	hooks    []Hook
}

// SetVideoBitrate sets the video_bitrate field.
func (mfc *MediaFileCreate) SetVideoBitrate(i int64) *MediaFileCreate {
	mfc.mutation.SetVideoBitrate(i)
	return mfc
}

// SetScaledWidth sets the scaled_width field.
func (mfc *MediaFileCreate) SetScaledWidth(i int16) *MediaFileCreate {
	mfc.mutation.SetScaledWidth(i)
	return mfc
}

// SetEncoderPreset sets the encoder_preset field.
func (mfc *MediaFileCreate) SetEncoderPreset(mp mediafile.EncoderPreset) *MediaFileCreate {
	mfc.mutation.SetEncoderPreset(mp)
	return mfc
}

// SetFramerate sets the framerate field.
func (mfc *MediaFileCreate) SetFramerate(i int8) *MediaFileCreate {
	mfc.mutation.SetFramerate(i)
	return mfc
}

// SetDurationSeconds sets the duration_seconds field.
func (mfc *MediaFileCreate) SetDurationSeconds(f float64) *MediaFileCreate {
	mfc.mutation.SetDurationSeconds(f)
	return mfc
}

// SetMediaType sets the media_type field.
func (mfc *MediaFileCreate) SetMediaType(mt mediafile.MediaType) *MediaFileCreate {
	mfc.mutation.SetMediaType(mt)
	return mfc
}

// SetCreatedAt sets the created_at field.
func (mfc *MediaFileCreate) SetCreatedAt(t time.Time) *MediaFileCreate {
	mfc.mutation.SetCreatedAt(t)
	return mfc
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (mfc *MediaFileCreate) SetNillableCreatedAt(t *time.Time) *MediaFileCreate {
	if t != nil {
		mfc.SetCreatedAt(*t)
	}
	return mfc
}

// SetUpdatedAt sets the updated_at field.
func (mfc *MediaFileCreate) SetUpdatedAt(t time.Time) *MediaFileCreate {
	mfc.mutation.SetUpdatedAt(t)
	return mfc
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (mfc *MediaFileCreate) SetNillableUpdatedAt(t *time.Time) *MediaFileCreate {
	if t != nil {
		mfc.SetUpdatedAt(*t)
	}
	return mfc
}

// SetID sets the id field.
func (mfc *MediaFileCreate) SetID(u uuid.UUID) *MediaFileCreate {
	mfc.mutation.SetID(u)
	return mfc
}

// Mutation returns the MediaFileMutation object of the builder.
func (mfc *MediaFileCreate) Mutation() *MediaFileMutation {
	return mfc.mutation
}

// Save creates the MediaFile in the database.
func (mfc *MediaFileCreate) Save(ctx context.Context) (*MediaFile, error) {
	var (
		err  error
		node *MediaFile
	)
	mfc.defaults()
	if len(mfc.hooks) == 0 {
		if err = mfc.check(); err != nil {
			return nil, err
		}
		node, err = mfc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MediaFileMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mfc.check(); err != nil {
				return nil, err
			}
			mfc.mutation = mutation
			node, err = mfc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(mfc.hooks) - 1; i >= 0; i-- {
			mut = mfc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mfc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mfc *MediaFileCreate) SaveX(ctx context.Context) *MediaFile {
	v, err := mfc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (mfc *MediaFileCreate) defaults() {
	if _, ok := mfc.mutation.CreatedAt(); !ok {
		v := mediafile.DefaultCreatedAt()
		mfc.mutation.SetCreatedAt(v)
	}
	if _, ok := mfc.mutation.UpdatedAt(); !ok {
		v := mediafile.DefaultUpdatedAt()
		mfc.mutation.SetUpdatedAt(v)
	}
	if _, ok := mfc.mutation.ID(); !ok {
		v := mediafile.DefaultID()
		mfc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mfc *MediaFileCreate) check() error {
	if _, ok := mfc.mutation.VideoBitrate(); !ok {
		return &ValidationError{Name: "video_bitrate", err: errors.New("ent: missing required field \"video_bitrate\"")}
	}
	if v, ok := mfc.mutation.VideoBitrate(); ok {
		if err := mediafile.VideoBitrateValidator(v); err != nil {
			return &ValidationError{Name: "video_bitrate", err: fmt.Errorf("ent: validator failed for field \"video_bitrate\": %w", err)}
		}
	}
	if _, ok := mfc.mutation.ScaledWidth(); !ok {
		return &ValidationError{Name: "scaled_width", err: errors.New("ent: missing required field \"scaled_width\"")}
	}
	if v, ok := mfc.mutation.ScaledWidth(); ok {
		if err := mediafile.ScaledWidthValidator(v); err != nil {
			return &ValidationError{Name: "scaled_width", err: fmt.Errorf("ent: validator failed for field \"scaled_width\": %w", err)}
		}
	}
	if _, ok := mfc.mutation.EncoderPreset(); !ok {
		return &ValidationError{Name: "encoder_preset", err: errors.New("ent: missing required field \"encoder_preset\"")}
	}
	if v, ok := mfc.mutation.EncoderPreset(); ok {
		if err := mediafile.EncoderPresetValidator(v); err != nil {
			return &ValidationError{Name: "encoder_preset", err: fmt.Errorf("ent: validator failed for field \"encoder_preset\": %w", err)}
		}
	}
	if _, ok := mfc.mutation.Framerate(); !ok {
		return &ValidationError{Name: "framerate", err: errors.New("ent: missing required field \"framerate\"")}
	}
	if v, ok := mfc.mutation.Framerate(); ok {
		if err := mediafile.FramerateValidator(v); err != nil {
			return &ValidationError{Name: "framerate", err: fmt.Errorf("ent: validator failed for field \"framerate\": %w", err)}
		}
	}
	if _, ok := mfc.mutation.DurationSeconds(); !ok {
		return &ValidationError{Name: "duration_seconds", err: errors.New("ent: missing required field \"duration_seconds\"")}
	}
	if v, ok := mfc.mutation.DurationSeconds(); ok {
		if err := mediafile.DurationSecondsValidator(v); err != nil {
			return &ValidationError{Name: "duration_seconds", err: fmt.Errorf("ent: validator failed for field \"duration_seconds\": %w", err)}
		}
	}
	if _, ok := mfc.mutation.MediaType(); !ok {
		return &ValidationError{Name: "media_type", err: errors.New("ent: missing required field \"media_type\"")}
	}
	if v, ok := mfc.mutation.MediaType(); ok {
		if err := mediafile.MediaTypeValidator(v); err != nil {
			return &ValidationError{Name: "media_type", err: fmt.Errorf("ent: validator failed for field \"media_type\": %w", err)}
		}
	}
	if _, ok := mfc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := mfc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New("ent: missing required field \"updated_at\"")}
	}
	return nil
}

func (mfc *MediaFileCreate) sqlSave(ctx context.Context) (*MediaFile, error) {
	_node, _spec := mfc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mfc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (mfc *MediaFileCreate) createSpec() (*MediaFile, *sqlgraph.CreateSpec) {
	var (
		_node = &MediaFile{config: mfc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: mediafile.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: mediafile.FieldID,
			},
		}
	)
	if id, ok := mfc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mfc.mutation.VideoBitrate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: mediafile.FieldVideoBitrate,
		})
		_node.VideoBitrate = value
	}
	if value, ok := mfc.mutation.ScaledWidth(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt16,
			Value:  value,
			Column: mediafile.FieldScaledWidth,
		})
		_node.ScaledWidth = value
	}
	if value, ok := mfc.mutation.EncoderPreset(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: mediafile.FieldEncoderPreset,
		})
		_node.EncoderPreset = value
	}
	if value, ok := mfc.mutation.Framerate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: mediafile.FieldFramerate,
		})
		_node.Framerate = value
	}
	if value, ok := mfc.mutation.DurationSeconds(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: mediafile.FieldDurationSeconds,
		})
		_node.DurationSeconds = value
	}
	if value, ok := mfc.mutation.MediaType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: mediafile.FieldMediaType,
		})
		_node.MediaType = value
	}
	if value, ok := mfc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: mediafile.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := mfc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: mediafile.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// MediaFileCreateBulk is the builder for creating a bulk of MediaFile entities.
type MediaFileCreateBulk struct {
	config
	builders []*MediaFileCreate
}

// Save creates the MediaFile entities in the database.
func (mfcb *MediaFileCreateBulk) Save(ctx context.Context) ([]*MediaFile, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mfcb.builders))
	nodes := make([]*MediaFile, len(mfcb.builders))
	mutators := make([]Mutator, len(mfcb.builders))
	for i := range mfcb.builders {
		func(i int, root context.Context) {
			builder := mfcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MediaFileMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mfcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mfcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, mfcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (mfcb *MediaFileCreateBulk) SaveX(ctx context.Context) []*MediaFile {
	v, err := mfcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
