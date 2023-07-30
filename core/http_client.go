package core

import "go.uber.org/fx"

type IHTTPClient interface {
	Get() string
}

type HTTPClient struct {
	fx.In
}

func (h HTTPClient) Get() string {
	return "pong"
}
