package remote

import (
	"context"
	"net/http"
	"net/url"

	"github.com/go-webdl/drm"
	"github.com/go-webdl/jsonrpc"
)

type RemoteOptions struct {
	IDStore jsonrpc.IDStore
	Base    http.RoundTripper
	URL     *url.URL
}

type registryT struct {
	jsonrpc.Client
	URL *url.URL
}

var _ drm.Registry = (*registryT)(nil)

func New(options *RemoteOptions) drm.Registry {
	r := new(registryT)
	r.IDStore, r.Base, r.URL = options.IDStore, options.Base, options.URL
	return r
}

func (c *registryT) Get(ctx context.Context, device string) (drm.CDM, error) {
	return &deviceClient{c, device}, nil
}

// Get devices CDM info for provided devices list. If len(devices) == 0, list all devices.
func (c *registryT) List(ctx context.Context, devices []string) (infos []*drm.DeviceInfo, err error) {
	infos = []*drm.DeviceInfo{}
	if err = c.Call(ctx, c.URL.String(), "list_device", devices, &infos); err != nil {
		return
	}
	if len(devices) != 0 && len(infos) != len(devices) {
		err = ErrInvalidResponse
		return
	}
	return
}

func (c *registryT) Register(ctx context.Context, devices []*drm.DeviceInfo) (err error) {
	var reply interface{}
	return c.Call(ctx, c.URL.String(), "register_device", devices, &reply)
}

func (c *registryT) Remove(ctx context.Context, devices []string) (err error) {
	var reply interface{}
	return c.Call(ctx, c.URL.String(), "remove_device", devices, &reply)
}
