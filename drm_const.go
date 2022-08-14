package drm

type SystemName string

const (
	WIDEVINE  SystemName = "WIDEVINE"
	FAIRPLAY  SystemName = "FAIRPLAY"
	PLAYREADY SystemName = "PLAYREADY"
)

var SystemID = map[SystemName][]byte{
	WIDEVINE:  {0xED, 0xEF, 0x8B, 0xA9, 0x79, 0xD6, 0x4A, 0xCE, 0xA3, 0xC8, 0x27, 0xDC, 0xD5, 0x1D, 0x21, 0xED},
	FAIRPLAY:  {0x29, 0x70, 0x1F, 0xE4, 0x3C, 0xC7, 0x4A, 0x34, 0x8C, 0x5B, 0xAE, 0x90, 0xC7, 0x43, 0x9A, 0x47},
	PLAYREADY: {0x9A, 0x04, 0xF0, 0x79, 0x98, 0x40, 0x42, 0x86, 0xAB, 0x92, 0xE6, 0x5B, 0xE0, 0x88, 0x5F, 0x95},
}

type SecurityLevel string

const (
	// Widevine Security Levels
	L1 SecurityLevel = "L1"
	L2 SecurityLevel = "L2"
	L3 SecurityLevel = "L3"

	// PlayReady Security Levels
	SL150  SecurityLevel = "SL150"
	SL2000 SecurityLevel = "SL2000"
	SL3000 SecurityLevel = "SL3000"

	// FairPlay Streaminig Security Levels
	AppleBaseline SecurityLevel = "AppleBaseline"
	AppleMain     SecurityLevel = "AppleMain"
)

type LicenseType string

const (
	STREAMING LicenseType = "STREAMING"
	OFFLINE   LicenseType = "OFFLINE"
)

type KeyType string

const (
	SIGNING          KeyType = "SIGNING"
	CONTENT          KeyType = "CONTENT"
	KEY_CONTROL      KeyType = "KEY_CONTROL"
	OPERATOR_SESSION KeyType = "OPERATOR_SESSION"
	ENTITLEMENT      KeyType = "ENTITLEMENT"
)

type Permissions string

const (
	ALLOW_ENCRYPT          Permissions = "ALLOW_ENCRYPT"
	ALLOW_DECRYPT          Permissions = "ALLOW_DECRYPT"
	ALLOW_SIGN             Permissions = "ALLOW_SIGN"
	ALLOW_SIGNATURE_VERIFY Permissions = "ALLOW_SIGNATURE_VERIFY"
)
