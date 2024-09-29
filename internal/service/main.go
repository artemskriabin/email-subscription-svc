package service

import (
	"github.com/artemskriabin/email-subscription-svc/internal/config"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gopkg.in/gomail.v2"
	"net"
	"net/http"
)

type service struct {
	log      *logan.Entry
	copus    types.Copus
	listener net.Listener
	dialer   *gomail.Dialer
	Sender   string
}

func (s *service) run() error {
	s.log.Info("Service started")

	m := gomail.NewMessage()
	m.SetHeader("From", s.Sender)
	m.SetHeader("To", "Alex@gmail.com")
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Normal from")
	m.SetBody("text/html", "Hello <b>Alice</b> and <i>Kora</i>!")
	//m.Attach("/home/Alex/lolcat.jpg")

	// Send the email to Bob, Cora and Dan.
	if err := s.dialer.DialAndSend(m); err != nil {
		panic(err)
	}

	r := s.router()

	if err := s.copus.RegisterChi(r); err != nil {
		return errors.Wrap(err, "cop failed")
	}

	return http.Serve(s.listener, r)
}

func newService(cfg config.Config) *service {
	return &service{
		log:      cfg.Log(),
		copus:    cfg.Copus(),
		listener: cfg.Listener(),
		dialer:   cfg.Dialer(),
		Sender:   cfg.Sender(),
	}
}

func Run(cfg config.Config) {
	if err := newService(cfg).run(); err != nil {
		panic(err)
	}
}
