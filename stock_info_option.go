package lssec_go

import "fmt"

// 뉴스본문
type NewsContentOption struct {
	NewsNumber string
}

func (NewsContentOption) String() string {
	return "t3102"
}

func (t NewsContentOption) Path() string {
	return "/stock/investinfo"
}

func (t NewsContentOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"sNewsno": "%s"
  }
}`, t.String(),
		t.NewsNumber,
	)), nil
}

// 종목별증시일정
type StockMarketScheduleOption struct {
	Ticker string
	Date   string
}

func (StockMarketScheduleOption) String() string {
	return "t3202"
}

func (t StockMarketScheduleOption) Path() string {
	return "/stock/investinfo"
}

func (t StockMarketScheduleOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"shcode": "%s",
	"date": "%s"
  }
}`, t.String(),
		t.Ticker,
		t.Date,
	)), nil
}

// FNG_요약
type FNGSummaryOption struct {
	Ticker string
}

func (FNGSummaryOption) String() string {
	return "t3320"
}

func (t FNGSummaryOption) Path() string {
	return "/stock/investinfo"
}

func (t FNGSummaryOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"gicode": "%s"
  }
}`, t.String(),
		t.Ticker,
	)), nil
}

// 재무순위종합
type FinancialRankingOption struct {
	// 0: 전체, 1: 코스피, 2: 코스닥
	Market string
	// 1:매출액증가율 2:영업이익증가율 3:세전계속이익증가율 4:부채비율 5:유보율 6:EPS 7:BPS 8:ROE 9:PER오름차순 a:PBR오름차순 b:PEG오름차순
	Ranking string
	// idx 첫조회시 space 연속조회시 Outblock의 idx 값 세팅
	Index int
}

func (FinancialRankingOption) String() string {
	return "t3341"
}

func (t FinancialRankingOption) Path() string {
	return "/stock/investinfo"
}

func (t FinancialRankingOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"idx": %d,
	"gubun": "%s",
	"gubun1": "%s",
	"gubun2": "1"
  }
}`, t.String(),
		t.Index,
		t.Market,
		t.Ranking,
	)), nil
}

// 투자의견
type InvestmentOpinionOption struct {
	Ticker    string
	IndexDate string
}

func (InvestmentOpinionOption) String() string {
	return "t3401"
}

func (t InvestmentOpinionOption) Path() string {
	return "/stock/investinfo"
}

func (t InvestmentOpinionOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"shcode": "%s",
	"cts_date": "%s",
	"gubun1": "",
	"tradno": ""
  }
}`, t.String(),
		t.Ticker,
		t.IndexDate,
	)), nil
}

// 해외실시간지수
type OverseasRealtimeIndexOption struct {
	// S:해외지수 F:해외선물 R:환율/금리
	Kind string
	// 다우산업 NAS@IXIC : 나스닥 종합 SPI@SPX : S&P 500 USI@SOXX : 필라델피아 반도체 NII@NI225 : 니케이 225 TWS@TI01 : 대만 가권 SHS@000002 : 상해 A SHS@000003 : 상해 B SGI@STI : 싱가폴 STI HSI@HSI : 항셍 PAS@CAC40 : 프랑스 CAC 40 LNS@FTSE100 : 영국 FTSE 100 XTR@DAX30 : 독일 DAX 30
	Symbol string
	Count  int
	// 조회구분 0:일 1:주 2:월 3:분 4:틱
	SearchClass string
	// SearchClass가 3인 경우에 n분
	MinNum int
	//다음 조회시 OutBlock의 cts_date 입력 처음 조회시 스페이스
	Date string
	// 다음 조회시 OutBlock의 cts_time 입력 처음 조회시 스페이스
	Time string
}

func (OverseasRealtimeIndexOption) String() string {
	return "t3518"
}

func (t OverseasRealtimeIndexOption) Path() string {
	return "/stock/investinfo"
}

func (t OverseasRealtimeIndexOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"kind": "%s",
	"symbol": "%s",
	"cnt": %d,
	"cts_date": "%s",
	"cts_time": "%s",
	"jgbn": "%s",
	"nmin": %d
  }
}`, t.String(),
		t.Kind,
		t.Symbol,
		t.Count,
		t.Date,
		t.Time,
		t.SearchClass,
		t.MinNum,
	)), nil
}

// 해외지수조회(API용)
type OverseasIndexOption struct {
	// S:해외지수 F:해외선물 R:환율/금리
	Kind string
	// 다우산업 NAS@IXIC : 나스닥 종합 SPI@SPX : S&P 500 USI@SOXX : 필라델피아 반도체 NII@NI225 : 니케이 225 TWS@TI01 : 대만 가권 SHS@000002 : 상해 A SHS@000003 : 상해 B SGI@STI : 싱가폴 STI HSI@HSI : 항셍 PAS@CAC40 : 프랑스 CAC 40 LNS@FTSE100 : 영국 FTSE 100 XTR@DAX30 : 독일 DAX 30
	Symbol string
}

func (OverseasIndexOption) String() string {
	return "t3521"
}

func (t OverseasIndexOption) Path() string {
	return "/stock/investinfo"
}

func (t OverseasIndexOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"kind": "%s",
	"symbol": "%s"
  }
}`, t.String(),
		t.Kind,
		t.Symbol,
	)), nil
}

// 증시주변자금추이
type MarketAroundMoneyOption struct {
	StartDate string
	EndDate   string
	//	1:예탁금 2:수익증권
	Type string
	// 	다음 조회시 사용함. 다음 조회시 OutBlock의 date 필드값 입력. 처음 조회시 Space
	Date string
	// 	001:코스피 301:코스닥
	Market string
	Count  int
}

func (MarketAroundMoneyOption) String() string {
	return "t8428"
}

func (t MarketAroundMoneyOption) Path() string {
	return "/stock/investinfo"
}

func (t MarketAroundMoneyOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"fdate": "%s",
	"tdate": "%s",
	"gubun": "%s",
	"key_date": "%s",
	"upcode": "%s",
	"cnt": %d
  }
}`, t.String(),
		t.StartDate,
		t.EndDate,
		t.Type,
		t.Date,
		t.Market,
		t.Count,
	)), nil
}
