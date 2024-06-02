package lssec_go

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ThemeStockOption(t *testing.T) {
	cli := ClientHelper(t)

	t.Run("ThemeStockOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), ThemeStockOption{Code: "0090", Name: ""}, "")
		require.NoError(t, err)
	})

	t.Run("StockThemeOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), StockThemeOption{
			Ticker: "314130",
		}, "")
		require.NoError(t, err)
	})

	t.Run("SpecialThemeOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), SpecialThemeOption{
			Ranking:      "1",
			ContrastDate: 0,
		}, "")
		require.NoError(t, err)
	})

	t.Run("ThemeStockQuoteOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), ThemeStockQuoteOption{
			Code: "0090",
		}, "")
		require.NoError(t, err)
	})

	t.Run("AllThemeOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), AllThemeOption{}, "")
		require.NoError(t, err)
	})
}
