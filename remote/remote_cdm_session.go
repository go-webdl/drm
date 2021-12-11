package remote

import (
	"context"

	"github.com/go-webdl/drm"
)

type clientSession struct {
	*registryT
	SessionData []byte `json:"session_data"`
}

var _ drm.Session = (*clientSession)(nil)

func (c *clientSession) ProvideLicense(ctx context.Context, license []byte) (keys []*drm.Key, err error) {
	keysList := [][]*drm.Key{}
	if err = c.Call(ctx, c.URL.String(), "provide_license", []*ProvideLicenseParam{{license, c.SessionData}}, &keysList); err != nil {
		return
	}
	if len(keysList) != 1 {
		err = ErrInvalidResponse
		return
	}
	keys = keysList[0]
	return
}
