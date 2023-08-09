package ebest_go

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_MarketData(t *testing.T) {
	cli := ClientHelper(t)

	t.Run("CurrentAskingPriceOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), CurrentAskingPriceOption{Ticker: "005930"})
		require.NoError(t, err)
	})

	t.Run("CurrentPriceOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), CurrentPriceOption{Ticker: "005930"})
		require.NoError(t, err)
	})

	t.Run("CurrentPriceMemoOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), CurrentPriceMemoOption{Ticker: "005930"})
		require.NoError(t, err)
	})

	t.Run("SearchPivotDemarkOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), SearchPivotDemarkOption{
			Ticker: "005930",
		})
		require.NoError(t, err)
	})

	t.Run("OvertimeTransactionCountOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), OvertimeTransactionCountOption{
			Ticker:     "005930",
			DanChetime: time.Now(),
			Index:      0,
		})
		require.NoError(t, err)
	})

	t.Run("TimeOfDayTransaction", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), TimeOfDayTransactionOption{
			Ticker:    "005930",
			Volume:    0,
			StartTime: time.Now().Add(-2 * time.Hour),
			EndTime:   time.Now(),
		})
		require.NoError(t, err)
	})

	t.Run("MinuteOfDayPriceOption", func(t *testing.T) {
		loc, _ := time.LoadLocation("Asia/Seoul")
		_, err := cli.Execute(context.TODO(), MinuteOfDayPriceOption{
			Ticker: "001200",
			Gubun:  "0",
			Time:   time.Date(0, 0, 0, 11, 0, 0, 0, loc),
			Count:  0,
		})
		require.NoError(t, err)
	})

}
