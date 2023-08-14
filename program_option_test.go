package ebest_go

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ProgramOption(t *testing.T) {
	cli := ClientHelper(t)

	t.Run("ProgramTradeSummaryOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), ProgramTradeSummaryOption{
			Market:    "1",
			DateType:  "1",
			StartDate: "",
			EndDate:   "",
		})
		require.NoError(t, err)
	})

	t.Run("ProgramTradeTimeSeriesOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), ProgramTradeTimeSeriesOption{
			Market:        "1",
			PriceOrAmount: "1",
		})
		require.NoError(t, err)
	})

	t.Run("ProgramTradeTrendOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), ProgramTradeTrendOption{
			Market:        "1",
			PriceOrAmount: "0",
			Sort:          "0",
			Ticker:        "005930",
			Index:         0,
		})
		require.NoError(t, err)
	})

	t.Run("ProgramTradeTimeSeriesByPeriodOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), ProgramTradeTimeSeriesByPeriodOption{
			Market:       "1",
			PriceType:    "0",
			Accumulation: "0",
			Period:       "1",
			StartDate:    "20230801",
			EndDate:      "20230810",
			ComparePre:   "0",
			Date:         "",
		})
		require.NoError(t, err)
	})

	t.Run("ProgramTradeTrendByStockOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), ProgramTradeTrendByStockOption{
			PriceOrAmount: "0",
			TimeOrDate:    "0",
			Ticker:        "005930",
			Date:          "20230814",
			Time:          "",
			ChartIndex:    9999,
		})
		require.NoError(t, err)
	})

}
