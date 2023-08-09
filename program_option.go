package ebest_go

import "fmt"

// 프로그램매매종합조회
type ProgramTradeSummaryOption struct {
	// 1 거래소, 2코스닥
	Market string
	// 1: 당일,2: 기간
	DateType  string
	StartDate string
	EndDate   string
}

func (ProgramTradeSummaryOption) String() string {
	return "t1631"
}

func (t ProgramTradeSummaryOption) Path() string {
	return "/stock/program"
}

func (t ProgramTradeSummaryOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"gubun": "%s",
	"dgubun": "%s",
	"sdate": "%s",
	"edate": "%s"
  }
}`, t.String(),
		t.Market,
		t.DateType,
		t.StartDate,
		t.EndDate,
	)), nil
}

// 시간대별프로그램매매추이
type ProgramTradeTimeSeriesOption struct {
	// 1 거래소, 2코스닥
	Market string
	// 0: 금액, 1: 수량
	PriceOrAmount string
	//처음 조회시는 Space 연속 조회시에 이전 조회한 OutBlock의 date 값으로 설정
	Date string
	// 처음 조회시는 Space 연속 조회시에 이전 조회한 OutBlock의 time 값으로 설정
	Time string
}

func (ProgramTradeTimeSeriesOption) String() string {
	return "t1632"
}

func (t ProgramTradeTimeSeriesOption) Path() string {
	return "/stock/program"
}

func (t ProgramTradeTimeSeriesOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"gubun": "%s",
	"gubun1": "%s",
	"gubun2": "1",
	"gubun3": "1",
	"date": "%s",
	"time": "%s"
  }
}`, t.String(),
		t.Market,
		t.PriceOrAmount,
		func() string {
			if t.Date == "" {
				return " "
			}
			return t.Date
		}(),
		func() string {
			if t.Time == "" {
				return " "
			}
			return t.Time
		}(),
	)), nil
}

// 종목별프로그램매매동향
type ProgramTradeTrendOption struct {
	// 0: 코스피, 1: 코스닥
	Market string
	// 0:수량 1:금액
	PriceOrAmount string
	// 0:시가총액비중 1:순매수상위 2:순매도상위 3:매도상위 4:매수상위
	Sort   string
	Ticker string
	Index  int
}

func (ProgramTradeTrendOption) String() string {
	return "t1636"
}

func (t ProgramTradeTrendOption) Path() string {
	return "/stock/program"
}

func (t ProgramTradeTrendOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"gubun": "%s",
	"gubun1": "%s",
	"gubun2": "%s",
	"shcode": "%s",
	"cts_idx": %d
  }
}`, t.String(),
		t.Market,
		t.PriceOrAmount,
		t.Sort,
		t.Ticker,
		t.Index,
	)), nil
}
