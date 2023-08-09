package ebest_go

import (
	"context"
	"testing"
)

func Test_ProgramOption(t *testing.T) {
	cli := ClientHelper(t)

	t.Run("ProgramTradeSummaryOption", func(t *testing.T) {
		cli.Execute(context.TODO(), ProgramTradeSummaryOption{
			Market:    "1",
			DateType:  "1",
			StartDate: "",
			EndDate:   "",
		})
	})

	t.Run("ProgramTradeTimeSeriesOption", func(t *testing.T) {
		cli.Execute(context.TODO(), ProgramTradeTimeSeriesOption{
			Market:        "1",
			PriceOrAmount: "1",
		})
	})

	t.Run("ProgramTradeTrendOption", func(t *testing.T) {
		cli.Execute(context.TODO(), ProgramTradeTrendOption{
			Market:        "1",
			PriceOrAmount: "0",
			Sort:          "0",
			Ticker:        "005930",
			Index:         0,
		})
	})

}
