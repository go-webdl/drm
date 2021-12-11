package local

import (
	"context"
	"sync"

	"github.com/go-webdl/drm"
)

var SystemsRegistry = map[drm.SystemName]func(device *drm.DeviceInfo) (cdm drm.CDM, err error){}

type localRegistry struct {
	sync.RWMutex
	devices map[string]*drm.DeviceInfo
}

func New() drm.Registry {
	return &localRegistry{devices: make(map[string]*drm.DeviceInfo)}
}

func (reg *localRegistry) Get(ctx context.Context, name string) (cdm drm.CDM, err error) {
	var ok bool
	var device *drm.DeviceInfo
	if device, ok = reg.devices[name]; device == nil || !ok {
		err = drm.ErrDeviceNotFound
		return
	}
	if factory, ok := SystemsRegistry[device.System]; ok {
		return factory(device)
	} else {
		err = drm.ErrUnknownDRMSystem
		return
	}
}

func (reg *localRegistry) List(ctx context.Context, names []string) (devices []*drm.DeviceInfo, err error) {
	if len(names) == 0 {
		for _, device := range reg.devices {
			devices = append(devices, device)
		}
	} else {
		for _, device := range reg.devices {
			for _, name := range names {
				if device.Device == name {
					devices = append(devices, device)
				}
			}
		}
	}
	return
}

func (reg *localRegistry) Register(ctx context.Context, devices []*drm.DeviceInfo) (err error) {
	for _, info := range devices {
		if _, ok := reg.devices[info.Device]; ok {
			err = drm.ErrDeviceExists
			return
		}
		reg.devices[info.Device] = info
	}
	return
}

func (reg *localRegistry) Remove(ctx context.Context, devices []string) (err error) {
	return
}
