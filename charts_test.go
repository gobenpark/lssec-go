package ebest_go

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_Charts(t *testing.T) {
	cli := ClientHelper(t)

	t.Run("TotalChartOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), TotalChartOption{
			Code:       KOSDAQ_CODE,
			ChartType:  "1",
			TickCount:  0,
			QueryCount: 500,
			Tdgb:       "0",
			StartDate:  "20230404",
			EndDate:    time.Now().Format("20060102"),
		})
		require.NoError(t, err)
	})

}
