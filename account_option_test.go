package ebest_go

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAccountTradeHistoryOption(t *testing.T) {
	cli := ClientHelper(t)
	t.Run("AccountTradeHistoryOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), AccountTradeHistoryOption{
			QueryType:   "1",
			StartDate:   "20230801",
			EndDate:     "20230810",
			StartNumber: 0,
			StockType:   "01",
			Ticker:      "KR7000020008",
		})
		require.NoError(t, err)
	})

	t.Run("AccountOrderHistoryOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), AccountOrderHistoryOption{
			Market:           "00",
			BuySell:          "0",
			Ticker:           "A080160",
			ContractExec:     "0",
			OrderDate:        "20230810",
			StartOrderNumber: 0,
			Sort:             "0",
			OrderTypeCode:    "00",
		})
		require.NoError(t, err)
	})

	t.Run("AccountOrderAvailableOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), AccountOrderAvailableOption{
			BuySell:    "2",
			Ticker:     "A005930",
			OrderPrice: 0.00,
		})
		require.NoError(t, err)
	})

	t.Run("AccountProfitDetailOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), AccountProfitDetailOption{
			StartDate: "20230801",
			EndDate:   "20230810",
			Period:    "1",
		})
		require.NoError(t, err)
	})

	t.Run("AccountDailyTradeOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), AccountDailyTradeOption{
			BuySell: "0",
			Ticker:  "1",
			Price:   "1",
			Where:   "1",
		})
		require.NoError(t, err)
	})

	t.Run("AccountBalanceOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), AccountBalanceOption{
			PriceType:       "",
			ContractType:    "",
			SinglePriceType: "",
			Charge:          "",
			Ticker:          "",
		})
		require.NoError(t, err)
	})

	t.Run("AccountContractOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), AccountContractOption{
			Ticker:      "080160",
			Contract:    "0",
			BuySell:     "0",
			Sort:        "2",
			OrderNumber: " ",
		})
		require.NoError(t, err)
	})
}
