// Code generated by entc, DO NOT EDIT.

package media

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the media type in the database.
	Label = "media"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldOriginalFilename holds the string denoting the original_filename field in the database.
	FieldOriginalFilename = "original_filename"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"

	// EdgeMediaFiles holds the string denoting the media_files edge name in mutations.
	EdgeMediaFiles = "media_files"

	// Table holds the table name of the media in the database.
	Table = "media"
	// MediaFilesTable is the table the holds the media_files relation/edge.
	MediaFilesTable = "media_file"
	// MediaFilesInverseTable is the table name for the MediaFile entity.
	// It exists in this package in order to avoid circular dependency with the "mediafile" package.
	MediaFilesInverseTable = "media_file"
	// MediaFilesColumn is the table column denoting the media_files relation/edge.
	MediaFilesColumn = "media"
)

// Columns holds all SQL columns for media fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldOriginalFilename,
	FieldStatus,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// DefaultOriginalFilename holds the default value on creation for the original_filename field.
	DefaultOriginalFilename string
	// OriginalFilenameValidator is a validator for the "original_filename" field. It is called by the builders before save.
	OriginalFilenameValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the created_at field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the updated_at field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the id field.
	DefaultID func() uuid.UUID
)

// Status defines the type for the status enum field.
type Status string

// Status values.
const (
	StatusAwaitingUpload Status = "AwaitingUpload"
	StatusProcessing     Status = "Processing"
	StatusReady          Status = "Ready"
	StatusErrored        Status = "Errored"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusAwaitingUpload, StatusProcessing, StatusReady, StatusErrored:
		return nil
	default:
		return fmt.Errorf("media: invalid enum value for status field: %q", s)
	}
}