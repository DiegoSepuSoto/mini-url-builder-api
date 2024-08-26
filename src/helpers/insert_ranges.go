package helpers

import (
	"net/http"
	"regexp"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/DiegoSepuSoto/mini-url-builder-api/src/infrastructure/http/repositories"
	"github.com/DiegoSepuSoto/mini-url-builder-api/src/infrastructure/http/repositories/sync"
)

func IsRangeConfigValid(counterInitialValue, counterMaxValue int64) bool {
	if counterInitialValue < 0 || counterMaxValue < 0 {
		return false
	}

	if counterInitialValue > counterMaxValue {
		return false
	}

	return true
}

func LoadIDRanges() error {
	httpClient := &http.Client{}

	syncRepository := sync.NewSyncRepository(httpClient)

	return loadIDRanges(syncRepository)
}

func loadIDRanges(syncRepository repositories.SyncRepository) error {
	rageIDs, err := syncRepository.GetIDRanges()

	if err != nil {
		return err
	}

	re := regexp.MustCompile(`(\d+)-(\d+)`)

	matches := re.FindStringSubmatch(rageIDs)

	counterInitialValue, _ := strconv.Atoi(matches[1])
	counterMaxValue, _ := strconv.Atoi(matches[2])

	viper.Set("COUNTER_VALUE", int64(counterInitialValue))
	viper.Set("COUNTER_MAX", int64(counterMaxValue))

	log.Infof("new counters set: %d to %d", counterInitialValue, counterMaxValue)

	return nil
}
