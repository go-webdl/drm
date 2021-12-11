// Package drm provides unified interfaces for all concrete CDM implementations.
//
// By unifying the CDM interface we abstract away the complexity in dealing with different underlying CDM implementations.
// For example, site clients that need access to CDM functions can accept a Session object instead of concrete CDM
// implementations as its parameter. So the consumer of the client code can choose from available concrete implementations.
package drm

import (
	"context"
)

type CDM interface {
	// Generates a license request challenge for the given options.
	GenerateChallenge(ctx context.Context, options *ChallengOptions) (challenge *Challenge, session Session, err error)

	// Get the underlying CDM implementation information.
	Info(ctx context.Context) (info *DeviceInfo, err error)
}

type StatelessCDM interface {
	CDM
	// Making the CDM server stateless by allow restoring a session from its internal states.
	RestoreSession(sessionData []byte) (session Session, err error)
}

type Session interface {
	// Updates the session with the provided license.
	ProvideLicense(ctx context.Context, license []byte) (keys []*Key, err error)
}

type StatelessSession interface {
	Session
	// Generate recoverable session data for its internal states.
	Serialize() (sessionData []byte, err error)
}

type Challenge struct {
	Challenge   []byte `json:"challenge"`
	ChallengeId []byte `json:"challenge_id"`
}

type ChallengOptions struct {
	// Common values are "cenc" and "webm" but implementation may add more.
	InitDataType string `json:"init_data_type,omitempty"`

	// An array of bytes. This is because license server may accept more than one init data at once, to request
	// licenses for multiple titles in a single license request.
	InitData [][]byte `json:"init_data"`

	LicenseType LicenseType `json:"license_type,omitempty"`

	// Optional. When it's nil, LicenseRequest.ClientId is set. If it's a valid Widevine proxy license server
	// certificate, then LicenseRequest.EncryptedClientId is set to the encrypted version of ClientIdentification.
	// Valid Widevine proxy license server certificate can be either a response wrapped in a SignedMessage structure,
	// usually responded from the license proxy server when issuing the request "CAQ=", which decodes to
	// SERVICE_CERTIFICATE_REQUEST. OR it can be the actual data in SignedDrmCertificate structure, which is usually
	// used in EME function call MediaKeys.setServerCertificate.
	ServerCert []byte `json:"server_cert,omitempty"`

	// If the implementation has VMP support, VMP data is included by default unless this option is set to true.
	DisableVmp bool `json:"disable_vmp"`

	// If ServerCert is set, LicenseRequest.EncryptedClientId is used unless this option is set to true, which will
	// force the implementation to use LicenseRequest.ClientId instead.
	DisableClientIDEncryption bool `json:"disable_client_id_encryption"`

	// If the implementation supports generating device ID, this paramter can be used to deterministically seed the
	// generator to produce fixed device ID. If not set, empty seed is used.
	DeviceIDSeed []byte `json:"device_id_seed,omitempty"`
}

type Key struct {
	Type        KeyType       `json:"type"`
	KID         []byte        `json:"kid"`
	Key         []byte        `json:"key"`
	Permissions []Permissions `json:"permissions"`
}

type DeviceInfo struct {
	System                SystemName    `json:"system"`
	Device                string        `json:"device,omitempty"`
	Version               string        `json:"version,omitempty"`
	SecurityLevel         SecurityLevel `json:"security_level,omitempty"`
	SupportRandomDeviceID bool          `json:"support_random_device_id"`
	DeviceData            []byte        `json:"device_data,omitempty"`
}
