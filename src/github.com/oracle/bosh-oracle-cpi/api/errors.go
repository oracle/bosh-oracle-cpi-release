package api

import (
	"fmt"
)

// CloudError represents an error thrown by
// CPI upon encountering failures when interacting
// with BMC
type CloudError interface {
	error

	// Type returns the error type string that can be interpreted
	// by BOSH (whihc invokes the CPI)
	Type() string
}

// RetryableError interface implemented by errors to communicate
// whether the operation resulting in the error can be retried
// or not
type RetryableError interface {
	error

	CanRetry() bool
}

// NotSupportedError is raised if a method is not supported
type NotSupportedError struct{}

// Type "Bosh::Clouds::NotSupported"
func (e NotSupportedError) Type() string  { return "Bosh::Clouds::NotSupported" }
func (e NotSupportedError) Error() string { return "Not supported" }

// VMNotFoundError is raised if a  can't be found in BMC
type VMNotFoundError struct {
	vmID string
}

// NewVMNotFoundError creates a VMNotFoundError instance with the ID
// of the vm that can't be found
func NewVMNotFoundError(vmID string) VMNotFoundError {
	return VMNotFoundError{vmID: vmID}
}

// Type "Bosh::Clouds::VMNotFound"
func (e VMNotFoundError) Type() string  { return "Bosh::Clouds::VMNotFound" }
func (e VMNotFoundError) Error() string { return fmt.Sprintf("VM '%s' not found", e.vmID) }

// VMCreationFailedError is raised if an instance can't be created in BMC
type VMCreationFailedError struct {
	reason   string
	canRetry bool
}

//NewVMCreationFailedError creates a VMCreationFailedError instance with the reason for the failure
func NewVMCreationFailedError(reason string, canRetry bool) VMCreationFailedError {
	return VMCreationFailedError{reason: reason, canRetry: canRetry}
}

// Type "Bosh::Clouds::VMCreationFailed"
func (e VMCreationFailedError) Type() string  { return "Bosh::Clouds::VMCreationFailed" }
func (e VMCreationFailedError) Error() string { return fmt.Sprintf("VM failed to create: %v", e.reason) }

// CanRetry returns if the operation can be retried
func (e VMCreationFailedError) CanRetry() bool { return e.canRetry }

// NoDiskSpaceError represents the  "Bosh::Clouds::NoDiskSpace" error
// in Go
type NoDiskSpaceError struct {
	diskID   string
	canRetry bool
}

// NewNoDiskSpaceError creates a NewNoDiskSpaceError instance
func NewNoDiskSpaceError(diskID string, canRetry bool) NoDiskSpaceError {
	return NoDiskSpaceError{diskID: diskID, canRetry: canRetry}
}

// Type "Bosh::Clouds::NoDiskSpace"
func (e NoDiskSpaceError) Type() string  { return "Bosh::Clouds::NoDiskSpace" }
func (e NoDiskSpaceError) Error() string { return fmt.Sprintf("Disk '%s' has no space", e.diskID) }

// CanRetry returns if the operation can be retried
func (e NoDiskSpaceError) CanRetry() bool { return e.canRetry }

// DiskNotAttachedError represents "Bosh::Clouds::DiskNotAttached" error
// in Go
type DiskNotAttachedError struct {
	vmID     string
	diskID   string
	canRetry bool
}

// NewDiskNotAttachedError creates a DiskNotAttachedError instance
func NewDiskNotAttachedError(vmID string, diskID string, canRetry bool) DiskNotAttachedError {
	return DiskNotAttachedError{vmID: vmID, diskID: diskID, canRetry: canRetry}
}

// Type "Bosh::Clouds::DiskNotAttached"
func (e DiskNotAttachedError) Type() string { return "Bosh::Clouds::DiskNotAttached" }
func (e DiskNotAttachedError) Error() string {
	return fmt.Sprintf("Disk '%s' not attached to VM '%s'", e.diskID, e.vmID)
}

// CanRetry returns if the operation can be retried
func (e DiskNotAttachedError) CanRetry() bool { return e.canRetry }

// DiskNotFoundError represents "Bosh::Clouds::DiskNotFound" error
// in Go
type DiskNotFoundError struct {
	diskID   string
	canRetry bool
}

// NewDiskNotFoundError creates a DiskNotFoundError instance
func NewDiskNotFoundError(diskID string, canRetry bool) DiskNotFoundError {
	return DiskNotFoundError{diskID: diskID, canRetry: canRetry}
}

// Type "Bosh::Clouds::DiskNotFound"
func (e DiskNotFoundError) Type() string  { return "Bosh::Clouds::DiskNotFound" }
func (e DiskNotFoundError) Error() string { return fmt.Sprintf("Disk '%s' not found", e.diskID) }

// CanRetry returns if the operation can be retried
func (e DiskNotFoundError) CanRetry() bool { return e.canRetry }
