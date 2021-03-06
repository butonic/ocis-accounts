package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/owncloud/ocis-accounts/pkg/account"
	"github.com/owncloud/ocis-accounts/pkg/config"
	"github.com/owncloud/ocis-accounts/pkg/proto/v0"
	olog "github.com/owncloud/ocis-pkg/v2/log"
)

// New returns a new instance of Service
func New(cfg *config.Config) Service {
	s := Service{
		Config: cfg,
	}

	if newReg, ok := account.Registry[cfg.Manager]; ok {
		s.Manager = newReg(cfg)
	} else {
		l := olog.NewLogger(olog.Name("ocis-accounts"))
		l.Fatal().Msgf("unknown manager: %v", cfg.Manager)
	}

	return s
}

// Service implements the SettingsServiceHandler interface
type Service struct {
	Config  *config.Config
	Manager account.Manager
}

// Set implements the SettingsServiceHandler interface
// This implementation replaces the existent data with the requested. It does not calculate diff
func (s Service) Set(c context.Context, req *proto.Record, res *proto.Record) error {
	r, err := s.Manager.Write(req)
	if err != nil {
		return err
	}

	res.Payload = r.GetPayload()
	return nil
}

// Get implements the SettingsServiceHandler interface
func (s Service) Get(c context.Context, req *proto.Query, res *proto.Record) error {
	r, err := s.Manager.Read(req.GetKey())
	if err != nil {
		return err
	}

	res.Payload = r.GetPayload()
	return nil
}

// List implements the SettingsServiceHandler interface
func (s Service) List(ctx context.Context, in *empty.Empty, res *proto.Records) error {
	r, err := s.Manager.List()
	if err != nil {
		return err
	}

	res.Records = r
	return nil
}
