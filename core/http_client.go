package core

type IHTTPClient interface {
	Get() string
}

type HTTPClient struct {
}

func NewHTTPClient() IHTTPClient {
	return &HTTPClient{}
}

func (h HTTPClient) Get() string {
	return "pong"
}
