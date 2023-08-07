package ebest_go

import "fmt"

type Chart string

const (
	TickChart  Chart = "0"
	MinChart   Chart = "1"
	DayChart   Chart = "2"
	WeekChart  Chart = "3"
	MonthChart Chart = "4"
)

type TotalChartOption struct {
	Code       SectorCode
	ChartType  Chart
	TickCount  int
	QueryCount int
	// 0: 전체, 1: 당일만
	Tdgb      string
	StartDate string
	EndDate   string
	CtsDate   string
	CtsTime   string
	CtsDaygb  string
}

func (TotalChartOption) String() string {
	return "t4203"
}

func (t TotalChartOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
    "shcode": "%s",
    "gubun": "%s",
    "ncnt": %d,
    "qrycnt": %d,
    "tdgb": "%s",
    "sdate": "%s",
    "edate": "%s",
    "cts_date": "%s",
    "cts_time": "%s",
    "cts_daygb": "%s"
  }
}`, t.String(),
		t.Code,
		t.ChartType,
		t.TickCount,
		t.QueryCount,
		t.Tdgb,
		t.StartDate,
		t.EndDate,
		t.CtsDate,
		t.CtsTime,
		t.CtsDaygb,
	)), nil
}

func (t TotalChartOption) Path() string {
	return "/indtp/chart"
}

type DaySectorChartOption struct {
	SectorCode  SectorCode
	ChartType   Chart
	Count       int
	StartDate   string
	EndDate     string
	Compression bool
}

func (DaySectorChartOption) String() string {
	return "t8419"
}

func (t DaySectorChartOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
    "shcode": "%s",
    "gubun": "%s",
    "qrycnt": %d,
    "sdate": "%s",
    "edate": "%s",
	"comp_yn": "%s"
  }
}`, t.String(),
		t.SectorCode,
		t.ChartType,
		t.Count,
		t.StartDate,
		t.EndDate,
		func() string {
			if t.Compression {
				return "Y"
			}
			return "N"
		}(),
	)), nil
}

func (DaySectorChartOption) Path() string {
	return "/indtp/chart"
}
