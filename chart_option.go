package ebest_go

import (
	"fmt"
)

type TickSectorChartOption struct {
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

func (TickSectorChartOption) String() string {
	return "t8417"
}

func (t TickSectorChartOption) MarshalJSON() ([]byte, error) {
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

func (t TickSectorChartOption) Path() string {
	return "/indtp/chart"
}

// 기간별투자자매매추이(챠트)
type InvestorTrendChartOption struct {
	// 1:코스피 2:KP200 3:코스닥 4:선물 5:풋옵션 6:콜옵션
	Market string
	// 업종코드
	SectorCode
	// 수치구분
	// 1: 수치, 2: 누적
	AmountType string
	// 1:일,2:주,3:월
	UnitType  string
	StartDate string
	EndDate   string
}

func (InvestorTrendChartOption) String() string {
	return "t1665"
}

func (InvestorTrendChartOption) Path() string {
	return "/stock/chart"
}

func (t InvestorTrendChartOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"market": "%s",
	"upcode": "%s",
	"gubun2": "%s",
	"gubun3": "%s",
	"from_date": "%s",
	"to_date": "%s"
  }
}`, t.String(),
		t.Market,
		t.SectorCode,
		t.AmountType,
		t.UnitType,
		t.StartDate,
		t.EndDate,
	)), nil
}

// API전용주식챠트(일주월년)
type StockChartOption struct {
	Ticker string
	// 2:일 3:주 4:월 5:년
	Period string
	// 최대압축 2000 비압축 500
	Count     int
	StartDate string
	EndDate   string
	// 압축여부
	Compress bool
	// 수정주가 여부
	Edited bool
}

func (StockChartOption) String() string {
	return "t8410"
}

func (StockChartOption) Path() string {
	return "/stock/chart"
}

func (t StockChartOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"shcode": "%s",
	"gubun": "%s",
	"qrycnt": %d,
	"sdate": "%s",
	"edate": "%s",
	"comp_yn": "%s",
	"sudung": "%s"
  }
}`, t.String(),
		t.Ticker,
		t.Period,
		t.Count,
		t.StartDate,
		t.EndDate,
		func() string {
			if t.Compress {
				return "Y"
			}
			return "N"
		}(),
		func() string {
			if t.Edited {
				return "Y"
			}
			return "N"
		}(),
	)), nil
}

type TickChartOption struct {
	Ticker string
	NTick  int
	Count  int
	// 조회영업일수(0:미사용1>=사용)
	NDay          int
	StartDate     string
	EndDate       string
	ContinuesDate string
	ContinuesTime string
	Compress      bool
}

func (TickChartOption) String() string {
	return "t8411"
}

func (TickChartOption) Path() string {
	return "/stock/chart"
}

func (t TickChartOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"shcode": "%s",
	"ncnt": %d,
	"qrycnt": %d,
	"nday": "%d",
	"sdate": "%s",
	"edate": "%s",
	"cts_date": "%s",
	"cts_time": "%s",
	"comp_yn": "%s"
  }
}`, t.String(),
		t.Ticker,
		t.NTick,
		t.Count,
		t.NDay,
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

type MinChartOption struct {
	Ticker string
	// n분
	MinType int
	// 최대압축 2000 비압축 500
	Count int
	// 	조회영업일수(0:미사용1>=사용)
	NDay          string
	StartDate     string
	EndDate       string
	ContinuesDate string
	ContinuesTime string
	Compress      bool
}

func (MinChartOption) String() string {
	return "t8412"
}

func (MinChartOption) Path() string {
	return "/stock/chart"
}

func (t MinChartOption) MarshalJSON() ([]byte, error) {
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
		t.Ticker,
		t.MinType,
		t.Count,
		t.NDay,
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
