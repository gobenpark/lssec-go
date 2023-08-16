package ebest_go

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_StockInfoOption(t *testing.T) {
	cli := ClientHelper(t)

	t.Run("news", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), NewsContentOption{
			NewsNumber: "2023051510383935PL7HQ87D",
		}, "")
		require.NoError(t, err)
	})

	t.Run("stock schedule", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), StockMarketScheduleOption{
			Date:   "",
			Ticker: "005930",
		}, "")
		require.NoError(t, err)
	})

	t.Run("FNG Summary", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), FNGSummaryOption{
			Ticker: "005930",
		}, "")
		require.NoError(t, err)
	})

	t.Run("FinancialRankingOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), FinancialRankingOption{
			Market:  "0",
			Ranking: "1",
			Index:   0,
		}, "")
		require.NoError(t, err)
	})

	t.Run("InvestmentOpinionOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), InvestmentOpinionOption{
			Ticker: "005930",
		}, "")
		require.NoError(t, err)
	})

	t.Run("OverseasRealtimeIndexOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), OverseasRealtimeIndexOption{
			Kind:        "S",
			Symbol:      "NAS@IXIC",
			Count:       20,
			SearchClass: "4",
			MinNum:      0,
			Date:        " ",
			Time:        " ",
		}, "")
		require.NoError(t, err)
	})

	t.Run("OverseasIndexOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), OverseasIndexOption{
			Kind:   "S",
			Symbol: "NAS@IXIC",
		}, "")
		require.NoError(t, err)
	})

	t.Run("MarketAroundMoneyOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), MarketAroundMoneyOption{
			StartDate: "",
			EndDate:   "",
			Type:      "",
			Date:      "",
			Market:    "001",
			Count:     1,
		}, "")
		require.NoError(t, err)
	})
}
