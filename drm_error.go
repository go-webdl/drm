package drm

import "errors"

var (
	ErrPermission         = errors.New("permission denied")
	ErrDeviceNotFound     = errors.New("device not found")
	ErrDeviceExists       = errors.New("device already exists")
	ErrInvalidSessionData = errors.New("invalid session data")
	ErrUnknownDRMSystem   = errors.New("unknown DRM system")
	ErrInvalidParams      = errors.New("invalid parameters")
	ErrInvalidLicense     = errors.New("invalid license")
	ErrInvalidInitData    = errors.New("invalid init data")
	ErrInternalFailure    = errors.New("internal failure")
)
