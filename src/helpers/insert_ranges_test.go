package helpers

import (
	"errors"
	"testing"

	"github.com/spf13/viper"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type syncRepositoryMock struct {
	mock.Mock
}

func (m *syncRepositoryMock) GetIDRanges() (string, error) {
	args := m.Called()

	return args.Get(0).(string), args.Error(1)
}

func TestInsertRanges(t *testing.T) {
	t.Run("when is range config valid receives invalid values then return false", func(t *testing.T) {
		isValid := IsRangeConfigValid(140001, 140000)

		assert.False(t, isValid)
	})

	t.Run("when is range config valid receives valid values then return true", func(t *testing.T) {
		isValid := IsRangeConfigValid(10, 20)

		assert.True(t, isValid)
	})

	t.Run("when load id ranges is executed as expected then the values are available on viper", func(t *testing.T) {
		mockSyncRepository := new(syncRepositoryMock)

		mockSyncRepository.On("GetIDRanges").Return("10-20", nil)

		err := loadIDRanges(mockSyncRepository)

		assert.Nil(t, err)
		assert.Equal(t, int64(10), viper.GetInt64("COUNTER_VALUE"))
		assert.Equal(t, int64(20), viper.GetInt64("COUNTER_MAX"))
	})

	t.Run("when load id ranges fails because of sync repository", func(t *testing.T) {
		mockSyncRepository := new(syncRepositoryMock)

		mockSyncRepository.On("GetIDRanges").Return("", errors.New("repository error"))

		err := loadIDRanges(mockSyncRepository)

		assert.NotNil(t, err)
		assert.Equal(t, "repository error", err.Error())
	})
}
