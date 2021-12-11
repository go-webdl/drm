package remote

import "github.com/go-webdl/drm"

type GenerateChallengeParam struct {
	*drm.ChallengOptions
	Device string `json:"device"`
}

type GenerateChallengeResult struct {
	*drm.Challenge
	EncryptedSessionData []byte `json:"session_data"`
}

type ProvideLicenseParam struct {
	License              []byte `json:"license"`
	EncryptedSessionData []byte `json:"session_data"`
}
