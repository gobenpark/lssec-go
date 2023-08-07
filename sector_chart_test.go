package ebest_go

import (
	"context"
	"ebest-go/test"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_SectorChart(t *testing.T) {
	key, secret := test.Secret(t)
	cli := NewClient(
		WithAuth(key, secret), WithAutomaticTokenCache(true))

	result, err := cli.Execute(context.TODO(), TotalChartOption{
		Code:       "001",
		ChartType:  DayChart,
		TickCount:  0,
		QueryCount: 500,
		Tdgb:       "",
		StartDate:  "20230405",
		EndDate:    "20230807",
		CtsDate:    "",
		CtsTime:    "",
		CtsDaygb:   "",
	})
	require.NoError(t, err)
	fmt.Println(string(result))
}
