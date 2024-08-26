package helpers

import (
	"math/big"
	"strings"

	"github.com/spf13/viper"
)

const base62 = 62

func CreateNewMiniURL() string {
	newURLExpectedLen := 6

	counterValue := viper.GetInt64("COUNTER_VALUE")
	counterMaxValue := viper.GetInt64("COUNTER_MAX")

	if !IsRangeConfigValid(counterValue, counterMaxValue) {
		err := LoadIDRanges()
		if err != nil {
			panic(err)
		}
	}

	newURL := big.NewInt(counterValue).Text(base62)
	newURLLen := len(newURL)

	if newURLLen < newURLExpectedLen {
		paddingLength := newURLExpectedLen - newURLLen

		newURL = strings.Repeat("0", paddingLength) + newURL
	}

	viper.Set("COUNTER_VALUE", counterValue+1)

	return newURL
}
