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
		_, err := cli.Execute(context.TODO(), CurrentAskingPriceOption{Ticker: "005930"}, "")
		require.NoError(t, err)
	})

	t.Run("CurrentPriceOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), CurrentPriceOption{Ticker: "005930"}, "")
		require.NoError(t, err)
	})

	t.Run("CurrentPriceMemoOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), CurrentPriceMemoOption{Ticker: "005930"}, "")
		require.NoError(t, err)
	})

	t.Run("SearchPivotDemarkOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), SearchPivotDemarkOption{
			Ticker: "005930",
		}, "")
		require.NoError(t, err)
	})

	t.Run("OvertimeTransactionCountOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), OvertimeTransactionCountOption{
			Ticker:     "005930",
			DanChetime: time.Now(),
			Index:      0,
		}, "")
		require.NoError(t, err)
	})

	t.Run("TimeOfDayTransaction", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), TimeOfDayTransactionOption{
			Ticker:    "005930",
			Volume:    0,
			StartTime: time.Now().Add(-2 * time.Hour),
			EndTime:   time.Now(),
		}, "")
		require.NoError(t, err)
	})

	t.Run("MinuteOfDayPriceOption", func(t *testing.T) {
		loc, _ := time.LoadLocation("Asia/Seoul")
		_, err := cli.Execute(context.TODO(), MinuteOfDayPriceOption{
			Ticker: "001200",
			Gubun:  "0",
			Time:   time.Date(0, 0, 0, 11, 0, 0, 0, loc),
			Count:  0,
		}, "")
		require.NoError(t, err)
	})

	t.Run("PeriodPriceOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), PeriodPriceOption{
			Ticker:   "001200",
			DateType: 1,
			Date:     "20230801",
			Count:    1,
		}, "")
		require.NoError(t, err)
	})

	t.Run("TimeOfDayTransactionChartOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), TimeOfDayTransactionChartOption{
			Ticker:    "005930",
			StartTime: "090000",
			EndTime:   "153000",
		}, "")
		require.NoError(t, err)
	})

	t.Run("DayMinuteTickOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), DayMinuteTickOption{
			Ticker: "005930",
		}, "")
		require.NoError(t, err)
	})

	t.Run("ManagementUnfaithfulInvestmentCautionOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), ManagementUnfaithfulInvestmentCautionOption{
			Market:          "0",
			Check:           "1",
			ContinuesTicker: "",
		}, "")
		require.NoError(t, err)
	})

	t.Run("InvestmentWarningStopClearanceOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), InvestmentWarningStopClearanceOption{
			Market:          "0",
			Check:           "1",
			ContinuesTicker: "",
		}, "")
		require.NoError(t, err)
	})

	t.Run("LowLiquidityOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), LowLiquidityOption{
			Market:          "0",
			ContinuesTicker: "",
		}, "")
		require.NoError(t, err)
	})

	t.Run("UpperLowerLimitOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), UpperLowerLimitOption{
			QueryType:  "1",
			Market:     "0",
			PredayType: "0",
			UpDown:     "1",
			Exclude:    8,
			StartPrice: 0,
			EndPrice:   0,
			Volume:     0,
			Index:      0,
		}, "")
		require.NoError(t, err)
	})

	t.Run("UpperLowerLimitBeforeOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), UpperLowerLimitBeforeOption{
			QueryType:     "1",
			Market:        "0",
			UpDown:        "1",
			Exclude:       0,
			ChangeRate:    0,
			StartPrice:    0,
			EndPrice:      0,
			Volume:        0,
			Index:         0,
			PredayExclude: "c",
		}, "")
		require.NoError(t, err)
	})

	t.Run("NewHighLowOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), NewHighLowOption{
			Market:     "0",
			UpDown:     "0",
			Exclude:    0,
			StartPrice: 0,
			EndPrice:   0,
			Volume:     0,
			Index:      0,
		}, "")
		require.NoError(t, err)
	})

	t.Run("PriceRangeTransactionRatioOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), PriceRangeTransactionRatioOption{
			Ticker:   "005930",
			DateType: "1",
		}, "")
		require.NoError(t, err)
	})

	t.Run("TimeOfDayAskingPriceOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), TimeOfDayAskingPriceOption{
			Ticker: "005930",
			Minute: "00",
			Time:   "",
			Count:  "010",
		}, "")
		require.NoError(t, err)
	})

	t.Run("TransactionIntensityOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), TransactionIntensityOption{
			Ticker:    "005930",
			UpDown:    "0",
			DataCount: 0,
			Date:      0,
			Time:      0,
			QueryType: "",
		}, "")
		require.NoError(t, err)
	})

	t.Run("TimeOfDayExpectedTransactionPriceOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), TimeOfDayExpectedTransactionPriceOption{
			Ticker:        "005930",
			ContinuesTime: "",
			Count:         20,
		}, "")
		require.NoError(t, err)
	})

	t.Run("ExpectedTransactionPriceChangeRateTopOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), ExpectedTransactionPriceChangeRateTopOption{
			Market:            "0",
			UpDown:            "1",
			MarketType:        "1",
			TickerCheck:       "0x00000080",
			Index:             0,
			Volume:            "0",
			StartPrice:        0,
			EndPrice:          0,
			TransactionVolume: 0,
		}, "")
		require.NoError(t, err)
	})

	t.Run("MultiCurrentPriceOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), MultiCurrentPriceOption{
			Tickers: "005930,005380",
		}, "")
		require.NoError(t, err)
	})

	t.Run("StockMasterOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), StockMasterOption{
			Type: "1",
		}, "")
		require.NoError(t, err)
	})
}
