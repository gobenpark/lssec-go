package ebest_go

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_Charts(t *testing.T) {
	cli := ClientHelper(t)

	t.Run("TotalChartOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), TotalChartOption{
			Code:       KOSDAQ_CODE,
			ChartType:  "1",
			TickCount:  0,
			QueryCount: 500,
			Tdgb:       "0",
			StartDate:  "20230404",
			EndDate:    time.Now().Format("20060102"),
		})
		require.NoError(t, err)
	})

	t.Run("InvestorTrendChartOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), InvestorTrendChartOption{
			Market:     "1",
			SectorCode: "001",
			AmountType: "1",
			UnitType:   "1",
			StartDate:  "20230604",
			EndDate:    time.Now().Format("20060102"),
		})
		require.NoError(t, err)
	})

	t.Run("StockChartOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), StockChartOption{
			Ticker:    "005930",
			Period:    "2",
			Count:     0,
			StartDate: "20230701",
			EndDate:   time.Now().Format("20060102"),
			Compress:  false,
			Edited:    true,
		})
		require.NoError(t, err)
	})

	t.Run("TickChartOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), TickChartOption{
			Ticker:        "005930",
			NTick:         1,
			Count:         0,
			NDay:          0,
			StartDate:     " ",
			EndDate:       " ",
			ContinuesDate: "",
			ContinuesTime: "",
			Compress:      false,
		})
		require.NoError(t, err)
	})

	t.Run("MinChartOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), MinChartOption{
			Ticker:        "005930",
			MinType:       3,
			Count:         0,
			NDay:          "0",
			StartDate:     "20230701",
			EndDate:       time.Now().Format("20060102"),
			ContinuesDate: "",
			ContinuesTime: "",
			Compress:      false,
		})
		require.NoError(t, err)
	})

}
