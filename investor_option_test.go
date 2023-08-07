package ebest_go

import (
	"context"
	"ebest-go/test"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_InvestorOption(t *testing.T) {
	key, secret := test.Secret(t)
	cli := NewClient(
		WithAuth(key, secret), WithAutomaticTokenCache(true))
	t.Run("TotalInvestorOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), TotalInvestorOption{
			StockAmountOrPrice:   "1",
			OptionAmountOrPrice:  "1",
			FuturesAmountOrPrice: "1",
		})
		require.NoError(t, err)
	})
}
