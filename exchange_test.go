package ebest_go

import (
	"context"
	"ebest-go/test"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Exchange(t *testing.T) {
	key, secret := test.Secret(t)
	cli := NewClient(
		WithAuth(key, secret), WithAutomaticTokenCache(true))
	t.Run("TopMemberCompany", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), TopMemberCompany{
			Ticker:    "005930",
			StartDate: "20230806",
			EndDate:   "20230807",
			Foreigner: true,
		})
		require.NoError(t, err)
	})

	t.Run("MemberShipTrandsByStock", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), MemberShipTrandsByStock{
			Ticker:       "005930",
			MemberNumber: "033",
			StartDate:    "20230806",
			EndDate:      "20230807",
			Gubun:        "0",
			Count:        500,
		})
		require.NoError(t, err)
	})
}
