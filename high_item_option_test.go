package lssec_go

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_HighItemOption(t *testing.T) {
	cli := ClientHelper(t)

	t.Run("UpDownRateTopOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), UpDownRateTopOption{
			Market:         "1",
			UpDown:         "1",
			DayOption:      "1",
			ExcludeOption:  0x00400000,
			StartPrice:     0,
			EndPrice:       0,
			Volume:         0,
			Index:          0,
			ExcludeOption2: 0,
		}, "")
		require.NoError(t, err)
	})

	t.Run("MarketCapTopOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), MarketCapTopOption{
			SectorCode: "001",
			Index:      0,
		}, "")
		require.NoError(t, err)
	})

	t.Run("VolumeTopOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), VolumeTopOption{
			Market:          "1",
			DayOption:       "1",
			StartUpDownRate: 0,
			EndUpDownRate:   0,
			ExcludeOption:   0,
			StartPrice:      0,
			EndPrice:        0,
			Volume:          0,
			Index:           0,
		}, "")
		require.NoError(t, err)
	})

	t.Run("AmountTopOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), AmountTopOption{
			Market:         "1",
			DayOption:      "1",
			ExcludeOption:  0,
			StartPrice:     0,
			EndPrice:       0,
			Volume:         0,
			Index:          0,
			ExcludeOption2: 0,
		}, "")
		require.NoError(t, err)
	})

	t.Run("VolumeIncreaseOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), VolumeIncreaseOption{
			Market:           "1",
			PreDayVolume:     "1",
			TradeRateTopRate: "0",
			ExcludeOption:    0,
			StartPrice:       0,
			EndPrice:         0,
			Volume:           0,
			Index:            0,
			ExcludeOption2:   0,
		}, "")
		require.NoError(t, err)
	})

	t.Run("AfterMarketPriceTopOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), AfterMarketPriceTopOption{
			Market:      "1",
			UpDownRate:  "1",
			TickerCheck: "1",
			Volume:      "1",
			Index:       0,
		}, "")
		require.NoError(t, err)
	})

	t.Run("AfterMarketVolumeTopOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), AfterMarketVolumeTopOption{
			Market: "1",
			Volume: "1",
			Index:  0,
		}, "")
		require.NoError(t, err)
	})

	t.Run("ExpectVolumeTopOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), ExpectVolumeTopOption{
			Market:           "1",
			OnOffTime:        "1",
			TickerCheck:      "1",
			Index:            0,
			ExpectStartPrice: 0,
			ExpectEndPrice:   0,
			ExpectVolume:     0,
		}, "")
		require.NoError(t, err)
	})

	t.Run("ExpectUpDownRateTopOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), ExpectUpDownRateTopOption{
			Market:      "1",
			UpDownRate:  "1",
			TickerCheck: "1",
			Volume:      "1",
			Index:       0,
		}, "")
		require.NoError(t, err)
	})

}
