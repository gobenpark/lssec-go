package ebest_go

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ETC(t *testing.T) {
	cli := ClientHelper(t)
	t.Run("NewTickerOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), NewTickerOption{
			Market:     "0",
			StartMonth: "202301",
			EndMonth:   "202308",
			Index:      0,
		})
		require.NoError(t, err)
	})

	t.Run("CollateralLoanableStockOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), CollateralLoanableStockOption{
			QueryType:        "0",
			Ticker:           "A005930",
			StockType:        "0",
			LoanInterestCode: "00",
			LoanType:         "1",
		})
		require.NoError(t, err)
	})

	t.Run("MarginRateByTickerOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), MarginRateByTickerOption{
			Market:     "0",
			CreditType: "1",
			MarginType: "1",
			Ticker:     "005930",
			Index:      0,
		})
		require.NoError(t, err)
	})

	t.Run("StockRemainderOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), StockRemainderOption{
			Market: "1",
			Ticker: "005930",
			Order:  "1",
		})
		require.NoError(t, err)
	})

	t.Run("CreditTradeTrendOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), CreditTradeTrendOption{
			Ticker:     "005930",
			CreditType: "1",
			Date:       "",
		})
		require.NoError(t, err)
	})

}
