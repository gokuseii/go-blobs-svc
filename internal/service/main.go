package service

import (
	"net"
	"net/http"

	"blobs-svc/internal/config"

	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type service struct {
	cfg      config.Config
	log      *logan.Entry
	copus    types.Copus
	listener net.Listener
}

func (s *service) run() error {
	s.log.Info("Service started")
	r := s.router(s.cfg)

	if err := s.copus.RegisterChi(r); err != nil {
		return errors.Wrap(err, "cop failed")
	}

	return http.Serve(s.listener, r)
}

func newService(cfg config.Config) *service {
	return &service{
		cfg:      cfg,
		log:      cfg.Log(),
		copus:    cfg.Copus(),
		listener: cfg.Listener(),
	}
}

func Run(cfg config.Config) {
	if err := newService(cfg).run(); err != nil {
		panic(err)
	}
}
