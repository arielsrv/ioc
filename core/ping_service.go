package core

import "go.uber.org/fx"

type PingService struct {
	fx.In
	HTTPClient HTTPClient
}

func (p PingService) Ping() string {
	return p.HTTPClient.Get()
}
