package core_test

import (
	"ioc/core"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/assert"
)

type HTTPClientMock struct {
	mock.Mock
}

func (m *HTTPClientMock) Get() string {
	args := m.Called()
	return args.Get(0).(string)
}

func TestPingService_Ping(t *testing.T) {
	httpClient := new(HTTPClientMock)
	httpClient.On("Get").Return("pong")

	pingService := core.NewPingService(httpClient)
	actual := pingService.Ping()

	assert.Equal(t, "pong", actual)
}
