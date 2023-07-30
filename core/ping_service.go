package core

type IPingService interface {
	Ping() string
}

type PingService struct {
	httpClient IHTTPClient
}

func NewPingService(httpClient IHTTPClient) IPingService {
	return &PingService{
		httpClient: httpClient,
	}
}

func (p PingService) Ping() string {
	return p.httpClient.Get()
}
