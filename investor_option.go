package ebest_go

import "fmt"

// 투자자별종합
type TotalInvestorOption struct {
	// 1: 수량, 2:금액
	StockAmountOrPrice string
	// 1: 수량, 2:금액
	OptionAmountOrPrice string
	// 1: 수량, 2:금액
	FuturesAmountOrPrice string
}

func (TotalInvestorOption) String() string {
	return "t1601"
}

func (TotalInvestorOption) Path() string {
	return "/stock/investor"
}

func (t TotalInvestorOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"gubun1": "%s",
	"gubun2": "%s",
	"gubun4": "%s"
  }
}`, t.String(),
		t.StockAmountOrPrice,
		t.OptionAmountOrPrice,
		t.FuturesAmountOrPrice,
	)), nil
}

// 시간대별투자자매매추이
type InvestorTradeTimeSeriesOption struct {
	// 	1:코스피 2:KP200 3:코스닥 4:선물 5:콜옵션 6:풋옵션 7:ELW 8:ETF
	Market     string
	SectorCode SectorCode
	// 1:수량 2:금액
	PriceOrAmount string
	// 	0:금일 1:전일
	PreDayDivision string
	ContinueTime   string
	Count          int
}

func (InvestorTradeTimeSeriesOption) String() string {
	return "t1602"
}

func (InvestorTradeTimeSeriesOption) Path() string {
	return "/stock/investor"
}

func (t InvestorTradeTimeSeriesOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"market": "%s",
	"upcode": "%s",
	"gubun1": "%s",
	"gubun2": "%s",
	"cts_time": "%s",
	"cnt": %d,
	"gubun3": "C"
  }
}`, t.String(),
		t.Market,
		t.SectorCode,
		t.PriceOrAmount,
		t.PreDayDivision,
		t.ContinueTime,
		t.Count,
	)), nil
}

// 시간대별투자자매매추이상세
type InvestorTradeTimeSeriesDetailOption struct {
	// 1:코스피 2:코스닥 3:선물 4:콜옵션 5:풋옵션 6:ELW 7:ETF
	Market     string
	SectorCode SectorCode
	// 1:개인 2:외인 3:기관계 4:증권 5:투신 6:은행 7:보험 8:종금 9:기금 A:국가 B:기타 C:사모펀드
	InvestorType string
	// 0:당일 1:전일
	PreDayDivision string
	ContinuesTime  string
	Index          int
	Count          int
}

func (InvestorTradeTimeSeriesDetailOption) String() string {
	return "t1603"
}

func (InvestorTradeTimeSeriesDetailOption) Path() string {
	return "/stock/investor"
}

func (t InvestorTradeTimeSeriesDetailOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"market": "%s",
	"upcode": "%s",
	"gubun1": "%s",
	"gubun2": "%s",
	"cts_time": "%s",
	"cts_idx": %d,
	"cnt": %d
  }
}`, t.String(),
		t.Market,
		t.SectorCode,
		t.InvestorType,
		t.PreDayDivision,
		t.ContinuesTime,
		t.Index,
		t.Count,
	)), nil
}

// 투자자매매종합1
type InvestorTradeSummaryOption1 struct {
	// 1:수량 2:금액
	StockPriceOrAmount string
	// 1:수량 2:금액
	OptionPriceOrAmount string
}

func (InvestorTradeSummaryOption1) String() string {
	return "t1615"
}

func (InvestorTradeSummaryOption1) Path() string {
	return "/stock/investor"
}

func (t InvestorTradeSummaryOption1) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"gubun1": "%s",
	"gubun2": "%s"
  }
}`, t.String(),
		t.StockPriceOrAmount,
		t.OptionPriceOrAmount,
	)), nil
}

// 투자자매매종합2
type InvestorTradeSummaryOption2 struct {
	// 1:코스피 2:코스닥 3:선물 4:콜옵션 5:풋옵션 6:주식선물 7:변동성 8:M선물 9:M콜옵션 0:M풋옵션
	Market string
	// 1:수량2:금액
	PriceOrAmount string
	// 1:시간대별2:일별
	DayDivision   string
	ContinuesTime string
	ContinuesDate string
}

func (InvestorTradeSummaryOption2) String() string {
	return "t1617"
}

func (InvestorTradeSummaryOption2) Path() string {
	return "/stock/investor"
}

func (t InvestorTradeSummaryOption2) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"gubun1": "%s",
	"gubun2": "%s",
	"gubun3": "%s",
	"cts_time": "%s",
	"cts_date": "%s"
  }
}`, t.String(),
		t.Market,
		t.PriceOrAmount,
		t.DayDivision,
		t.ContinuesTime,
		t.ContinuesDate,
	)), nil
}
