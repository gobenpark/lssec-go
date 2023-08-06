package ebest_go

import (
	"context"
	"ebest-go/test"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_Charts(t *testing.T) {
	key, secret := test.Secret(t)
	cli := NewClient(
		WithAuth(key, secret), WithAutomaticTokenCache(true))

	t.Run("TotalChartOption", func(t *testing.T) {
		_, err := cli.Charts(context.TODO(), "", TotalChartOption{
			Code:       "",
			ChartType:  "",
			TickCount:  0,
			QueryCount: 0,
			Tdgb:       "",
			StartDate:  time.Time{},
			EndDate:    time.Time{},
			CtsDate:    time.Time{},
			CtsTime:    time.Time{},
			CtsDaygb:   "",
		})
		require.NoError(t, err)
	})

}
