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

// 기간별프로그램매매추이
type ProgramTradeTimeSeriesByPeriodOption struct {
	// 0 거래소, 1코스닥
	Market string
	// 0: 금액, 1: 수량
	PriceType string
	// 수치누적구분 0:수치, 1:누적
	Accumulation string
	// 일주월구분 1:일 2:주 3:월
	Period    string
	StartDate string
	EndDate   string
	//직전대비증감구분 0:default 1:직전대비 증감
	ComparePre string
	// 처음 조회시는 Space 연속 조회시에 이전 조회한 OutBlock의 date 값으로 설정
	Date string
}

func (ProgramTradeTimeSeriesByPeriodOption) String() string {
	return "t1633"
}

func (t ProgramTradeTimeSeriesByPeriodOption) Path() string {
	return "/stock/program"
}

func (t ProgramTradeTimeSeriesByPeriodOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"gubun": "%s",
	"gubun1": "%s",
	"gubun2": "%s",
	"gubun3": "%s",
	"fdate": "%s",
	"tdate": "%s",
	"gubun4": "%s",
	"date": "%s"
  }
}`, t.String(),
		t.Market,
		t.PriceType,
		t.Accumulation,
		t.Period,
		t.StartDate,
		t.EndDate,
		t.ComparePre,
		t.Date,
	)), nil
}

// 종목별프로그램매매추이
type ProgramTradeTrendByStockOption struct {
	// 수량금액구분(0:수량,1:금액)
	PriceOrAmount string
	// 시간일별구분(0:시간,1:일자)
	TimeOrDate string
	Ticker     string
	Date       string
	Time       string
	// 차트조회시에만 9999입력
	ChartIndex int
}

func (ProgramTradeTrendByStockOption) String() string {
	return "t1637"
}

func (t ProgramTradeTrendByStockOption) Path() string {
	return "/stock/program"
}

func (t ProgramTradeTrendByStockOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"gubun": "%s",
	"gubun1": "%s",
	"shcode": "%s",
	"date": "%s",
	"time": "%s",
	"cts_idx": %d
  }
}`, t.String(),
		t.PriceOrAmount,
		t.TimeOrDate,
		t.Ticker,
		t.Date,
		t.Time,
		t.ChartIndex,
	)), nil
}
