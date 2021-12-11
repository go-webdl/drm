package remote

import (
	"context"

	"github.com/go-webdl/drm"
)

type deviceClient struct {
	*registryT
	device string
}

var _ drm.CDM = (*deviceClient)(nil)

func (c *registryT) GenerateChallenge(ctx context.Context, device string, options *drm.ChallengOptions) (challenge *drm.Challenge, session drm.Session, err error) {
	challenges := []*GenerateChallengeResult{}
	if err = c.Call(ctx, c.URL.String(), "generate_challenge", []*GenerateChallengeParam{{options, device}}, &challenges); err != nil {
		return
	}
	if len(challenges) != 1 {
		err = ErrInvalidResponse
		return
	}
	challenge = challenges[0].Challenge
	session = &clientSession{c, challenges[0].EncryptedSessionData}
	return
}

func (c *deviceClient) GenerateChallenge(ctx context.Context, options *drm.ChallengOptions) (challenge *drm.Challenge, session drm.Session, err error) {
	return c.registryT.GenerateChallenge(ctx, c.device, options)
}

func (c *deviceClient) Info(ctx context.Context) (info *drm.DeviceInfo, err error) {
	infos, err := c.List(ctx, []string{c.device})
	if err != nil {
		return
	}
	info = infos[0]
	return
}
