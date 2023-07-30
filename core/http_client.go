package core

type IHTTPClient interface {
	GetMessage() string
}

type HTTPClient struct {
}

func NewHTTPClient() IHTTPClient {
	return &HTTPClient{}
}

func (h HTTPClient) GetMessage() string {
	return "pong"
}
