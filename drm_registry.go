package drm

import (
	"context"
)

type Registry interface {
	Get(ctx context.Context, device string) (CDM, error)
	List(ctx context.Context, devices []string) ([]*DeviceInfo, error)
	Register(ctx context.Context, devices []*DeviceInfo) error
	Remove(ctx context.Context, devices []string) error
}
