package types

import "errors"

var (
	ErrUploadFailed        = errors.New("All upload failed")
	ErrUploadFailedPartial = errors.New("Some items failed to be uploaded")
	ErrRollbackFailed      = errors.New("Rollback failed. You have orphaned images.")
)
