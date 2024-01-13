package core

type IHTTPClient interface {
	GetMessage() string
}

type HTTPClient struct {
	config *Config
}

func NewHTTPClient(config *Config) *HTTPClient {
	return &HTTPClient{
		config: config,
	}
}

func (h *HTTPClient) GetMessage() string {
	return h.config.Message
}
