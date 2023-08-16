package ebest_go

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_InvestorOption(t *testing.T) {
	cli := ClientHelper(t)
	t.Run("TotalInvestorOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), TotalInvestorOption{
			StockAmountOrPrice:   "1",
			OptionAmountOrPrice:  "1",
			FuturesAmountOrPrice: "1",
		}, "")
		require.NoError(t, err)
	})

	t.Run("InvestorTradeTimeSeriesOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), InvestorTradeTimeSeriesOption{
			Market:         "1",
			SectorCode:     KOSPI_CODE,
			PriceOrAmount:  "1",
			PreDayDivision: "0",
			ContinueTime:   " ",
			Count:          1,
		}, "")
		require.NoError(t, err)
	})

	t.Run("InvestorTradeTimeSeriesDetailOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), InvestorTradeTimeSeriesDetailOption{
			Market:         "1",
			SectorCode:     KOSPI_CODE,
			InvestorType:   "1",
			PreDayDivision: "0",
			ContinuesTime:  " ",
			Index:          0,
			Count:          20,
		}, "")
		require.NoError(t, err)
	})

	t.Run("InvestorTradeSummaryOption1", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), InvestorTradeSummaryOption1{
			StockPriceOrAmount:  "1",
			OptionPriceOrAmount: "1",
		}, "")
		require.NoError(t, err)
	})

	t.Run("InvestorTradeSummaryOption2", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), InvestorTradeSummaryOption2{
			Market:        "1",
			PriceOrAmount: "1",
			DayDivision:   "1",
		}, "")
		require.NoError(t, err)
	})

	t.Run("InvestorTradeSectorTimeSeriesOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), InvestorTradeSectorTimeSeriesOption{
			SectorCode: KOSPI_CODE,
			NMin:       5,
			Count:      1,
			Day:        "1",
		}, "")
		require.NoError(t, err)
	})

	t.Run("InvestorTradeSummaryChartOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), InvestorTradeSummaryChartOption{
			Market:        "1",
			PriceOrAmount: "1",
			DayDivision:   "1",
			Count:         10,
		}, "")
		require.NoError(t, err)
	})
}
