package ebest_go

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_SectorChart(t *testing.T) {
	cli := ClientHelper(t)

	t.Run("TotalChartOption", func(t *testing.T) {

		_, err := cli.Execute(context.TODO(), TotalChartOption{
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
		}, "")
		require.NoError(t, err)
	})

	t.Run("DaySectorChartOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), DaySectorChartOption{
			SectorCode:  "001",
			ChartType:   DayChart,
			Count:       2000,
			StartDate:   "20230701",
			EndDate:     "20230807",
			Compression: false,
		}, "")
		require.NoError(t, err)
	})

	t.Run("MinSectorChartOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), MinSectorChartOption{
			SectorCode: "001",
			Count:      2000,
			StartDate:  "20230701",
			EndDate:    "20230807",
		}, "")
		require.NoError(t, err)
	})

}
