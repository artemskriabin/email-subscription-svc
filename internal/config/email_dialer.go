package config

import (
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gopkg.in/gomail.v2"
)

type Emailer interface {
	Dialer() *gomail.Dialer
	Sender() string
}

type emailer struct {
	getter kv.Getter
	once   comfig.Once
}

func NewEmailer(getter kv.Getter) Emailer {
	return &emailer{
		getter: getter,
	}
}

type emialerCfg struct {
	Host     string `fig:"host"`
	Sender   string `fig:"sender,required"`
	Password string `fig:"password,required"`
	Port     int    `fig:"port"`
}

func (d *emailer) Dialer() *gomail.Dialer {
	return d.once.Do(func() interface{} {
		cfg := d.readConfig()

		return gomail.NewDialer(cfg.Host, cfg.Port, cfg.Sender, cfg.Password)
	}).(*gomail.Dialer)
}

func (d *emailer) readConfig() *emialerCfg {
	cfg := emialerCfg{
		Host: "smtp.gmail.com",
		Port: 587,
	}
	err := figure.Out(&cfg).
		From(kv.MustGetStringMap(d.getter, "email")).
		Please()
	if err != nil {
		panic(errors.Wrap(err, "failed to figure out"))
	}
	return &cfg
}

func (d *emailer) Sender() string {
	return d.readConfig().Sender
}
