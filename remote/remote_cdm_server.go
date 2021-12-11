package remote

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-webdl/crypto/nacl"
	"github.com/go-webdl/drm"
	"github.com/go-webdl/jsonrpc"

	"google.golang.org/protobuf/proto"
)

var (
	ErrInternal      = errors.New("internal error")
	ErrInvalidParams = errors.New("invalid request params")
)

type ServerOptions struct {
	Registry   drm.Registry
	EncryptKey *nacl.Key // optional
}

type serverT struct {
	*jsonrpc.Server
	*ServerOptions
}

func NewServer(options *ServerOptions) (handler http.Handler, err error) {
	server := &serverT{Server: new(jsonrpc.Server), ServerOptions: options}
	if server.EncryptKey == nil {
		server.EncryptKey = nacl.New(nil)
	}
	if err = server.Register("generate_challenge", server.generateChallenge); err != nil {
		return
	}
	if err = server.Register("provide_license", server.provideLicense); err != nil {
		return
	}
	if err = server.Register("list_device", server.listDevice); err != nil {
		return
	}
	if err = server.Register("register_device", server.registerDevice); err != nil {
		return
	}
	if err = server.Register("remove_device", server.removeDevice); err != nil {
		return
	}
	handler = server
	return
}

func (s *serverT) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.Server.ServeHTTP(w, req)
}

func (s *serverT) generateChallenge(r *http.Request, params *[]*GenerateChallengeParam, reply *[]*GenerateChallengeResult) (err error) {
	ctx := r.Context()
	if params == nil {
		err = ErrInvalidParams
		return
	}
	for _, param := range *params {
		var cdm drm.CDM
		if cdm, err = s.Registry.Get(ctx, param.Device); err != nil {
			return
		}

		var session drm.Session
		var challenge *drm.Challenge
		if challenge, session, err = cdm.GenerateChallenge(context.Background(), param.ChallengOptions); err != nil {
			return
		}

		var ok bool
		var serializableSession drm.StatelessSession
		if serializableSession, ok = session.(drm.StatelessSession); !ok {
			return ErrInternal
		}

		var deviceSessionData []byte
		if deviceSessionData, err = serializableSession.Serialize(); err != nil {
			return
		}

		sessionData := &SessionData{
			Device:            param.Device,
			DeviceSessionData: deviceSessionData,
		}

		var sessionDataBytes []byte
		if sessionDataBytes, err = proto.Marshal(sessionData); err != nil {
			return
		}

		var encryptedSessionData []byte
		if encryptedSessionData, err = s.EncryptKey.Encrypt(nil, sessionDataBytes, nil); err != nil {
			return
		}

		*reply = append(*reply, &GenerateChallengeResult{challenge, encryptedSessionData})
	}
	return
}

func (s *serverT) provideLicense(r *http.Request, params *[]*ProvideLicenseParam, reply *[][]*drm.Key) (err error) {
	ctx := r.Context()
	if params == nil {
		err = ErrInvalidParams
		return
	}
	for _, param := range *params {
		var sessionDataBytes []byte
		if sessionDataBytes, err = s.EncryptKey.Decrypt(nil, param.EncryptedSessionData); err != nil {
			return
		}

		sessionData := new(SessionData)
		if err = proto.Unmarshal(sessionDataBytes, sessionData); err != nil {
			return
		}

		var cdm drm.CDM
		if cdm, err = s.Registry.Get(ctx, sessionData.Device); err != nil {
			return
		}

		var ok bool
		var recoverableCDM drm.StatelessCDM
		if recoverableCDM, ok = cdm.(drm.StatelessCDM); !ok {
			return ErrInternal
		}

		var session drm.Session
		if session, err = recoverableCDM.RestoreSession(sessionData.DeviceSessionData); err != nil {
			return
		}

		var keys []*drm.Key
		if keys, err = session.ProvideLicense(context.Background(), param.License); err != nil {
			return
		}

		*reply = append(*reply, keys)
	}
	return
}

func (s *serverT) listDevice(r *http.Request, devices *[]string, reply *[]*drm.DeviceInfo) (err error) {
	if devices == nil {
		err = ErrInvalidParams
		return
	}
	*reply, err = s.Registry.List(r.Context(), *devices)
	return
}

func (s *serverT) registerDevice(r *http.Request, devices *[]*drm.DeviceInfo, reply *interface{}) (err error) {
	if devices == nil {
		err = ErrInvalidParams
		return
	}
	return s.Registry.Register(r.Context(), *devices)
}

func (s *serverT) removeDevice(r *http.Request, devices *[]string, reply *interface{}) (err error) {
	if devices == nil {
		err = ErrInvalidParams
		return
	}
	return s.Registry.Remove(r.Context(), *devices)
}
