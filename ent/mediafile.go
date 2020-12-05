// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/dreamvo/gilfoyle/ent/mediafile"
	"github.com/facebook/ent/dialect/sql"
	"github.com/google/uuid"
)

// MediaFile is the model entity for the MediaFile schema.
type MediaFile struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// VideoBitrate holds the value of the "video_bitrate" field.
	VideoBitrate int64 `json:"video_bitrate,omitempty"`
	// ScaledWidth holds the value of the "scaled_width" field.
	ScaledWidth int16 `json:"scaled_width,omitempty"`
	// EncoderPreset holds the value of the "encoder_preset" field.
	EncoderPreset mediafile.EncoderPreset `json:"encoder_preset,omitempty"`
	// Framerate holds the value of the "framerate" field.
	Framerate int8 `json:"framerate,omitempty"`
	// DurationSeconds holds the value of the "duration_seconds" field.
	DurationSeconds float64 `json:"duration_seconds,omitempty"`
	// MediaType holds the value of the "media_type" field.
	MediaType mediafile.MediaType `json:"media_type,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MediaFile) scanValues() []interface{} {
	return []interface{}{
		&uuid.UUID{},       // id
		&sql.NullInt64{},   // video_bitrate
		&sql.NullInt64{},   // scaled_width
		&sql.NullString{},  // encoder_preset
		&sql.NullInt64{},   // framerate
		&sql.NullFloat64{}, // duration_seconds
		&sql.NullString{},  // media_type
		&sql.NullTime{},    // created_at
		&sql.NullTime{},    // updated_at
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MediaFile fields.
func (mf *MediaFile) assignValues(values ...interface{}) error {
	if m, n := len(values), len(mediafile.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	if value, ok := values[0].(*uuid.UUID); !ok {
		return fmt.Errorf("unexpected type %T for field id", values[0])
	} else if value != nil {
		mf.ID = *value
	}
	values = values[1:]
	if value, ok := values[0].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field video_bitrate", values[0])
	} else if value.Valid {
		mf.VideoBitrate = value.Int64
	}
	if value, ok := values[1].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field scaled_width", values[1])
	} else if value.Valid {
		mf.ScaledWidth = int16(value.Int64)
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field encoder_preset", values[2])
	} else if value.Valid {
		mf.EncoderPreset = mediafile.EncoderPreset(value.String)
	}
	if value, ok := values[3].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field framerate", values[3])
	} else if value.Valid {
		mf.Framerate = int8(value.Int64)
	}
	if value, ok := values[4].(*sql.NullFloat64); !ok {
		return fmt.Errorf("unexpected type %T for field duration_seconds", values[4])
	} else if value.Valid {
		mf.DurationSeconds = value.Float64
	}
	if value, ok := values[5].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field media_type", values[5])
	} else if value.Valid {
		mf.MediaType = mediafile.MediaType(value.String)
	}
	if value, ok := values[6].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field created_at", values[6])
	} else if value.Valid {
		mf.CreatedAt = value.Time
	}
	if value, ok := values[7].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field updated_at", values[7])
	} else if value.Valid {
		mf.UpdatedAt = value.Time
	}
	return nil
}

// Update returns a builder for updating this MediaFile.
// Note that, you need to call MediaFile.Unwrap() before calling this method, if this MediaFile
// was returned from a transaction, and the transaction was committed or rolled back.
func (mf *MediaFile) Update() *MediaFileUpdateOne {
	return (&MediaFileClient{config: mf.config}).UpdateOne(mf)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (mf *MediaFile) Unwrap() *MediaFile {
	tx, ok := mf.config.driver.(*txDriver)
	if !ok {
		panic("ent: MediaFile is not a transactional entity")
	}
	mf.config.driver = tx.drv
	return mf
}

// String implements the fmt.Stringer.
func (mf *MediaFile) String() string {
	var builder strings.Builder
	builder.WriteString("MediaFile(")
	builder.WriteString(fmt.Sprintf("id=%v", mf.ID))
	builder.WriteString(", video_bitrate=")
	builder.WriteString(fmt.Sprintf("%v", mf.VideoBitrate))
	builder.WriteString(", scaled_width=")
	builder.WriteString(fmt.Sprintf("%v", mf.ScaledWidth))
	builder.WriteString(", encoder_preset=")
	builder.WriteString(fmt.Sprintf("%v", mf.EncoderPreset))
	builder.WriteString(", framerate=")
	builder.WriteString(fmt.Sprintf("%v", mf.Framerate))
	builder.WriteString(", duration_seconds=")
	builder.WriteString(fmt.Sprintf("%v", mf.DurationSeconds))
	builder.WriteString(", media_type=")
	builder.WriteString(fmt.Sprintf("%v", mf.MediaType))
	builder.WriteString(", created_at=")
	builder.WriteString(mf.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(mf.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// MediaFiles is a parsable slice of MediaFile.
type MediaFiles []*MediaFile

func (mf MediaFiles) config(cfg config) {
	for _i := range mf {
		mf[_i].config = cfg
	}
}
