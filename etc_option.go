package lssec_go

import "fmt"

// 신규상장종목조회
type NewTickerOption struct {
	// 0: 전체, 1:코스피, 2:코스닥
	Market string
	// 200601
	StartMonth string
	// 200601
	EndMonth string
	Index    int
}

func (NewTickerOption) String() string {
	return "t1403"
}

func (NewTickerOption) Path() string {
	return "/stock/etc"
}

func (t NewTickerOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"gubun": "%s",
	"fromdt": "%s",
	"todt": "%s",
	"idx": %d
  }
}`, t.String(),
		t.Market,
		t.StartMonth,
		t.EndMonth,
		t.Index,
	)), nil
}

// 예탁담보융자가능종목현황조회
type CollateralLoanableStockOption struct {
	// 0: 전체, 1:가능, 2:불가능
	QueryType string
	// A+종목번호
	Ticker string
	//유가증권구분
	StockType string
	// 대출이자등급코드
	LoanInterestCode string
	// 대출구분
	LoanType string
}

func (CollateralLoanableStockOption) String() string {
	return "CLNAQ00100"
}

func (CollateralLoanableStockOption) Path() string {
	return "/stock/etc"
}

func (t CollateralLoanableStockOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock1": {
	"QryTp": "%s",
	"IsuNo": "%s",
	"SecTpCode": "%s",
	"LoanIntrstGrdCode": "%s",
	"LoanTp": "%s"
  }
}`, t.String(),
		t.QueryType,
		t.Ticker,
		t.StockType,
		t.LoanInterestCode,
		t.LoanType,
	)), nil
}

// 증거금율별종목조회
type MarginRateByTickerOption struct {
	// 0: 전체, 1:코스피, 2:코스닥
	Market string
	// 위탁신용구분
	CreditType string
	// 증거금구분
	MarginType string
	Ticker     string
	Index      int
}

func (MarginRateByTickerOption) String() string {
	return "t1411"
}

func (MarginRateByTickerOption) Path() string {
	return "/stock/etc"
}

func (t MarginRateByTickerOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"gubun": "%s",
	"jongchk": "%s",
	"jkrate": "%s",
	"shcode": "%s",
	"idx": %d
  }
}`, t.String(),
		t.Market,
		t.CreditType,
		t.MarginType,
		t.Ticker,
		t.Index,
	)), nil
}

// 종목별잔량/사전공시
type StockRemainderOption struct {
	// 1:코스피, 2:코스닥
	Market string
	Ticker string
	// 1:시가총액비중 2:순매수잔량상위 3: 순매수잔량하위 4:매수잔량 5:매수공시수량 6:매도잔량 7:매도공시수량
	Order string
}

func (StockRemainderOption) String() string {
	return "t1638"
}

func (StockRemainderOption) Path() string {
	return "/stock/etc"
}

func (t StockRemainderOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"gubun1": "%s",
	"shcode": "%s",
	"gubun2": "%s"
  }
}`, t.String(),
		t.Market,
		t.Ticker,
		t.Order,
	)), nil
}

// 신용거래동향
type CreditTradeTrendOption struct {
	Ticker string
	// 융자대주구분 1:융자 2:대주
	CreditType string
	Date       string
}

func (CreditTradeTrendOption) String() string {
	return "t1921"
}

func (CreditTradeTrendOption) Path() string {
	return "/stock/etc"
}

func (t CreditTradeTrendOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"shcode": "%s",
	"gubun": "%s",
	"date": "%s"
  }
}`, t.String(),
		t.Ticker,
		t.CreditType,
		t.Date,
	)), nil
}

// 종목별신용정보
type StockCreditInfoOption struct {
	Ticker string
}

func (StockCreditInfoOption) String() string {
	return "t1926"
}

func (StockCreditInfoOption) Path() string {
	return "/stock/etc"
}

func (t StockCreditInfoOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"shcode": "%s"
  }
}`, t.String(),
		t.Ticker,
	)), nil
}

// 공매도일별추이
type ShortStockDailyOption struct {
	Ticker string
	// 다음 조회시 사용. OutBlock의 date 필드 값을 입력함.
	Date      string
	StartDate string
	EndDate   string
}

func (ShortStockDailyOption) String() string {
	return "t1927"
}

func (ShortStockDailyOption) Path() string {
	return "/stock/etc"
}

func (t ShortStockDailyOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"shcode": "%s",
	"date": "%s",
	"sdate": "%s",
	"edate": "%s"
  }
}`, t.String(),
		t.Ticker,
		t.Date,
		t.StartDate,
		t.EndDate,
	)), nil
}

// 종목별대차거래일간추이
type StockMarginTradeTrendOption struct {
	Ticker    string
	StartDate string
	EndDate   string
}

func (StockMarginTradeTrendOption) String() string {
	return "t1941"
}

func (StockMarginTradeTrendOption) Path() string {
	return "/stock/etc"
}

func (t StockMarginTradeTrendOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"shcode": "%s",
	"sdate": "%s",
	"edate": "%s"
  }
}`, t.String(),
		t.Ticker,
		t.StartDate,
		t.EndDate,
	)), nil
}

// 주식종목조회
type StockOption struct {
	// 0:전체, 1:코스피, 2:코스닥
	Market string
}

func (StockOption) String() string {
	return "t8430"
}

func (StockOption) Path() string {
	return "/stock/etc"
}

func (t StockOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"gubun": "%s"
  }
}`, t.String(),
		t.Market,
	)), nil
}

// 주식종목조회 API용
type StockOptionAPI struct {
	// 0:전체, 1:코스피, 2:코스닥
	Market string
}

func (StockOptionAPI) String() string {
	return "t8436"
}

func (StockOptionAPI) Path() string {
	return "/stock/etc"
}

func (t StockOptionAPI) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"gubun": "%s"
  }
}`, t.String(),
		t.Market,
	)), nil
}
