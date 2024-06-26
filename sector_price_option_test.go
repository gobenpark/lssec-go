package lssec_go

import (
	"context"
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

func Test_MarketSectorData(t *testing.T) {
	cli := ClientHelper(t)

	t.Run("IndustryTrendsOverTime", func(t *testing.T) {
		result, err := cli.Execute(context.TODO(), IndustryTrendsOverTimeOption{
			SectorCode: "001",
			Date:       "1",
			CtsDate:    "20230803",
			Count:      1,
			Rate:       1,
		}, "")
		require.NoError(t, err)
		require.Greater(t, len(result), 0)
	})

	t.Run("TotalIndustry", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), TotalIndustryOption{}, "")
		require.NoError(t, err)
	})

	t.Run("ExpectedStockIndex", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), ExpectedStockIndexOption{
			SectorCode: "001",
			Gubun:      "2",
		}, "")
		require.NoError(t, err)
	})

	t.Run("CurrentPriceOfIndustryOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), CurrentPriceOfIndustryOption{
			SectorCode: "101",
		}, "")
		require.NoError(t, err)
	})

	t.Run("PriceOfIndustryOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), PriceOfIndustryOption{
			SectorCode: "100",
			Gubun:      "1",
			Ticker:     "005930",
		}, "")
		require.NoError(t, err)
	})
}
