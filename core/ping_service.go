package core

import "go.uber.org/fx"

type IPingService interface {
	Ping() string
}

type PingService struct {
	fx.In
	HTTPClient HTTPClient
}

func (p PingService) Ping() string {
	return p.HTTPClient.Get()
}
