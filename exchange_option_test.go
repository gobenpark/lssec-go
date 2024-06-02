package lssec_go

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Exchange(t *testing.T) {
	cli := ClientHelper(t)
	t.Run("TopMemberCompanyOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), TopMemberCompanyOption{
			Ticker:    "005930",
			StartDate: "20230806",
			EndDate:   "20230807",
			Foreigner: true,
		}, "")
		require.NoError(t, err)
	})

	t.Run("MemberShipTrandsByStockOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), MemberShipTrandsByStockOption{
			Ticker:       "005930",
			MemberNumber: "033",
			StartDate:    "20230806",
			EndDate:      "20230807",
			TimeOrDate:   "0",
			Count:        500,
		}, "")
		require.NoError(t, err)
	})

	t.Run("MemberShipListOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), MemberShipListOption{
			Ticker: "005930",
		}, "")
		require.NoError(t, err)
	})
}
