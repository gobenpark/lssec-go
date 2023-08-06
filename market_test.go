package ebest_go

import (
	"context"
	"ebest-go/test"
	"testing"

	"github.com/stretchr/testify/require"
)

//
//{
//"t1514InBlock": {
//"upcode": "001",
//"gubun1": "",
//"gubun2": "1",
//"cts_date": "20230804",
//"cnt": 1,
//"rate_gbn": "1"
//}
//}

func Test_MarketData(t *testing.T) {
	key, secret := test.Secret(t)
	cli := NewClient(
		WithAuth(key, secret), WithAutomaticTokenCache(true))

	t.Run("IndustryTrendsOverTime", func(t *testing.T) {
		_, err := cli.MarketData(context.TODO(), "", IndustryTrendsOverTimeOption{
			Code:    "001",
			Date:    "1",
			CtsDate: "20230803",
			Count:   1,
			Rate:    1,
		})
		require.NoError(t, err)
	})

	t.Run("TotalIndustry", func(t *testing.T) {
		_, err := cli.MarketData(context.TODO(), "", TotalIndustryOption{})
		require.NoError(t, err)
	})

	t.Run("ExpectedStockIndex", func(t *testing.T) {
		_, err := cli.MarketData(context.TODO(), "", ExpectedStockIndexOption{
			Code:  "001",
			Gubun: "2",
		})
		require.NoError(t, err)
	})

	t.Run("CurrentPriceOfIndustryOption", func(t *testing.T) {
		_, err := cli.MarketData(context.TODO(), "", CurrentPriceOfIndustryOption{
			Code: "101",
		})
		require.NoError(t, err)
	})

	t.Run("PriceOfIndustryOption", func(t *testing.T) {
		_, err := cli.MarketData(context.TODO(), "005930", PriceOfIndustryOption{
			Code:   "100",
			Gubun:  "1",
			Ticker: "005930",
		})
		require.NoError(t, err)
	})
}
