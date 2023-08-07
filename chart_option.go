package ebest_go

import (
	"fmt"
)

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

func (t TickChartOption) Path() string {
	return "/indtp/chart"
}
