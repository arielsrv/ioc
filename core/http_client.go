package core

import (
	"go.uber.org/fx"
)

type IHTTPClient interface {
	GetMessage() string
}

type HTTPClient struct {
	fx.In

	config *Config
}

func NewHTTPClient(config *Config) IHTTPClient {
	return &HTTPClient{
		config: config,
	}
}

func (h HTTPClient) GetMessage() string {
	return h.config.Message
}
