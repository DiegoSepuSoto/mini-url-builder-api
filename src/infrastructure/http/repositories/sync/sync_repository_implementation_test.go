package sync

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

type HTTPClientMock struct {
	mock.Mock
}

func (m *HTTPClientMock) Get(url string) (*http.Response, error) {
	args := m.Called(url)

	if args.Get(1) == nil {
		return args.Get(0).(*http.Response), nil
	}

	return nil, args.Get(1).(error)
}
