package lssec_go

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

// 업종챠트(N분)
type MinSectorChartOption struct {
	SectorCode SectorCode
	// 단위(N분)
	NMin int
	// 조회건수
	Count int
	// 조회영업일수(0:미사용,1>= 사용)
	Day           string
	StartDate     string
	EndDate       string
	ContinuesDate string
	ContinuesTime string
	Compress      bool
}

func (MinSectorChartOption) String() string {
	return "t8418"
}

func (MinSectorChartOption) Path() string {
	return "/indtp/chart"
}

func (t MinSectorChartOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"shcode": "%s",
	"ncnt": %d,
	"qrycnt": %d,
	"nday": "%s",
	"sdate": "%s",
	"edate": "%s",
	"cts_date": "%s",
	"cts_time": "%s",
	"comp_yn": "%s"
  }
}`, t.String(),
		t.SectorCode,
		t.NMin,
		t.Count,
		t.Day,
		t.StartDate,
		t.EndDate,
		t.ContinuesDate,
		t.ContinuesTime,
		func() string {
			if t.Compress {
				return "Y"
			}
			return "N"
		}(),
	)), nil
}
