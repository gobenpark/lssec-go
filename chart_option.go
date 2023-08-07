package ebest_go

import (
	"fmt"
)

type TotalChartOption struct {
	Code SectorCode
	// 0: 틱,1: 분, 2:일, 3:주, 4:월
	ChartType  string
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

type TickChartOption struct {
	Code SectorCode
	// 요청건수 (최대-압축: 2000 비압축:500)
	QueryCount int
	// 단위(n틱)
	Nday      string
	StartDate string
	EndDate   string
	TickCount int
	CtsDate   string
	CtsTime   string
	Compress  bool
}

func (TickChartOption) String() string {
	return "t8417"
}

func (t TickChartOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
    "shcode": "%s",
    "ncnt": %d,
    "qrycnt": %d,
    "tdgb": "%s",
    "sdate": "%s",
    "edate": "%s",
    "cts_date": "%s",
  }
}`, t.String(),
		t.Code,
		t.TickCount,
		t.QueryCount,
		t.StartDate,
		t.EndDate,
		t.CtsDate,
		t.CtsTime,
	)), nil
}
