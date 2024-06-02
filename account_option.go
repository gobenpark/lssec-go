package lssec_go

import "fmt"

// 계좌 거래내역
type AccountTradeHistoryOption struct {
	//0:전체, 1:입출금, 2:입출고, 3:매매, 4:환전, 9:기타
	QueryType   string
	StartDate   string
	EndDate     string
	StartNumber int
	// 00:전체, 01:주식, 02:채권, 04:펀드, 03:선물, 05:해외주식, 06:해외파생
	StockType string
	Ticker    string
}

func (AccountTradeHistoryOption) String() string {
	return "CDPCQ04700"
}

func (t AccountTradeHistoryOption) Path() string {
	return "/stock/accno"
}

func (t AccountTradeHistoryOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock1": {
	"RecCnt": 1,
	"QryTp": "%s",
	"QrySrtDt": "%s",
	"QryEndDt": "%s",
	"SrtNo": %d,
	"PdptnCode": "01",
	"IsuLgclssCode": "%s",
	"IsuNo": "%s"
  }
}`, t.String(),
		t.QueryType,
		t.StartDate,
		t.EndDate,
		t.StartNumber,
		t.StockType,
		t.Ticker,
	)), nil
}

// 현물계좌 주문체결내역 조회(API)
type AccountOrderHistoryOption struct {
	//00:전체 10:거래소 20:코스닥 30:프리보드
	Market string
	// 	매매구분
	// 0:전체 1: 매도 2: 매수
	BuySell string
	//주식 : A+종목코드
	//ELW : J+종목코드
	Ticker string
	// 체결여부 0:전체 1:체결 3:미체결
	ContractExec string
	OrderDate    string
	// 역순구분이 순 : 000000000 역순구분이 역순 : 999999999
	StartOrderNumber int
	// 역순 구분 0:역순 1:정순
	Sort string
	//00:전체
	//98:매도전체
	//99:매수전체
	//01:현금매도
	//02:현금매수
	//05:저축매도
	//06:저축매수
	//09:상품매도
	//10:상품매수
	//03:융자매도
	//04:융자매수
	//07:대주매도
	//08:대주매수
	//11:선물대용매도
	//13:현금매도(프)
	//14:현금매수(프)
	//17:대출
	//18:대출상환
	OrderTypeCode string
}

func (AccountOrderHistoryOption) String() string {
	return "CSPAQ13700"
}

func (t AccountOrderHistoryOption) Path() string {
	return "/stock/accno"
}

func (t AccountOrderHistoryOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock1": {
	"OrdMktCode": "%s",
	"BnsTpCode": "%s",
	"IsuNo": "%s",
	"ExecYn": "%s",
	"OrdDt": "%s",
	"SrtOrdNo2": %d,
	"BkseqTpCode": "%s",
	"OrdPtnCode": "%s" 
  }
}`, t.String(),
		t.Market,
		t.BuySell,
		t.Ticker,
		t.ContractExec,
		t.OrderDate,
		t.StartOrderNumber,
		t.Sort,
		t.OrderTypeCode,
	)), nil
}

// 현물계좌증거금률별주문가능수량조회
type AccountOrderAvailableOption struct {
	// 1:매도, 2:매수
	BuySell    string
	Ticker     string
	OrderPrice float32
}

func (AccountOrderAvailableOption) String() string {
	return "CSPBQ00200"
}

func (t AccountOrderAvailableOption) Path() string {
	return "/stock/accno"
}

func (t AccountOrderAvailableOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock1": {
	"BnsTpCode": "%s",
	"IsuNo": "%s",
	"OrdPrc": %.2f
  }
}`, t.String(),
		t.BuySell,
		t.Ticker,
		t.OrderPrice,
	)), nil
}

// 주식계좌 기간별수익률 상세
type AccountProfitDetailOption struct {
	StartDate string
	EndDate   string
	// 1:일별, 2:주별, 3:월별
	Period string
}

func (AccountProfitDetailOption) String() string {
	return "FOCCQ33600"
}

func (t AccountProfitDetailOption) Path() string {
	return "/stock/accno"
}

func (t AccountProfitDetailOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock1": {
	"QrySrtDt": "%s",
	"QryEndDt": "%s",
	"TermTp": "%s"
  }
}`, t.String(),
		t.StartDate,
		t.EndDate,
		t.Period,
	)), nil
}

// 주식당일매매일지/수수료
type AccountDailyTradeOption struct {
	// 	0:전체 1:매도 2:매수
	BuySell string
	Ticker  string
	Price   string
	Where   string
}

func (AccountDailyTradeOption) String() string {
	return "t0150"
}

func (t AccountDailyTradeOption) Path() string {
	return "/stock/accno"
}

func (t AccountDailyTradeOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"cts_medosu": "%s",
	"cts_expcode": "%s",
	"cts_price": "%s",
	"cts_middiv": "%s"
  }
}`, t.String(),
		t.BuySell,
		t.Ticker,
		t.Price,
		t.Where,
	)), nil
}

// 주식잔고2
type AccountBalanceOption struct {
	PriceType       string
	ContractType    string
	SinglePriceType string
	Charge          string
	Ticker          string
}

func (AccountBalanceOption) String() string {
	return "t0424"
}

func (t AccountBalanceOption) Path() string {
	return "/stock/accno"
}

func (t AccountBalanceOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"prcgb": "%s",
	"chegb": "%s",
	"dangb": "%s",
	"charge": "%s",
	"cts_expcode": "%s"
  }
}`, t.String(),
		t.PriceType,
		t.ContractType,
		t.SinglePriceType,
		t.Charge,
		t.Ticker,
	)), nil
}

// 주식체결/미체결
type AccountContractOption struct {
	Ticker string
	// 	0:전체 1:체결 2:미체결
	Contract string
	// 	0:전체 1:매도 2:매수
	BuySell string
	// 1:주문번호 역순 2:주문번호 순
	Sort        string
	OrderNumber string
}

func (AccountContractOption) String() string {
	return "t0425"
}

func (t AccountContractOption) Path() string {
	return "/stock/accno"
}

func (t AccountContractOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"expcode": "%s",
	"chegb": "%s",
	"medosu": "%s",
	"sortgb": "%s",
	"cts_ordno": "%s"
  }
}`, t.String(),
		t.Ticker,
		t.Contract,
		t.BuySell,
		t.Sort,
		t.OrderNumber,
	)), nil
}

// 현물계좌예수금 주문가능금액 총평가 조회
type AccountOrderAvailableTotalOption struct {
}

func (AccountOrderAvailableTotalOption) String() string {
	return "CSPAQ12200"
}

func (t AccountOrderAvailableTotalOption) Path() string {
	return "/stock/accno"
}

func (t AccountOrderAvailableTotalOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"RecCnt": 1,
	"MgmtBrnNo": "1",
	"BalCreTp": "1"
  }
}`, t.String(),
	)), nil
}
