package core

type IPingService interface {
	Ping() string
}

type PingService struct {
	HTTPClient IHTTPClient
}

func NewPingService(httpClient IHTTPClient) *PingService {
	return &PingService{
		HTTPClient: httpClient,
	}
}

func (p PingService) Ping() string {
	return p.HTTPClient.GetMessage()
}
