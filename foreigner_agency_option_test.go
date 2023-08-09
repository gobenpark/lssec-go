package ebest_go

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ForeignerAgencyOption(t *testing.T) {
	cli := ClientHelper(t)
	t.Run("ForeignerAgencyOption1", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), ForeignerAgencyOption1{
			Ticker:         "001200",
			EndDate:        "",
			PriceOrAmount:  "0",
			TradeType:      "0",
			AccType:        "0",
			CountinuesDate: "",
			Index:          0,
		})
		require.NoError(t, err)
	})
}
