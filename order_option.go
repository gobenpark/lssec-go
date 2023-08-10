package ebest_go

import "fmt"

type OrderOption struct {
	// A+종목코드
	Ticker string
	Amount int
	Price  int
	//1:매도 2:매수
	BuySell string
	//00:지정가 03:시장가 05:조건부지정가 06:최유리지정가 07:최우선지정가 61:장개시전시간외종가 81:시간외종가 82:시간외단일가
	OrderType string
	// 000:보통 003:유통/자기융자신규 005:유통대주신규 007:자기대주신규 101:유통융자상환 103:자기융자상환 105:유통대주상환 107:자기대주상환 180:예탁담보대출상환(신용)
	CreditCode string
	// 대출일
	LoanDate string
	// 주문조건 구분
	//0:없음,1:IOC,2:FOK
	OrderCondition string
}

func (OrderOption) String() string {
	return "CSPAT00601"
}

func (t OrderOption) Path() string {
	return "/stock/order"
}

func (t OrderOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock1": {
	"IsuNo": "%s",
	"OrdQty": %d,
	"OrdPrc": %d,
	"BnsTpCode": "%s",
	"OrdprcPtnCode": "%s",
	"MgntrnCode": "%s",
	"LoanDt": "%s",
	"OrdCndiTpCode": "%s"
  }
}`, t.String(),
		t.Ticker,
		t.Amount,
		t.Price,
		t.BuySell,
		t.OrderType,
		t.CreditCode,
		t.LoanDate,
		t.OrderCondition,
	)), nil
}

type OrderEditOption struct {
	// 주문번호
	OrderNumber int
	// A+종목코드
	Ticker string
	Amount int
	Price  int
	//00:지정가 03:시장가 05:조건부지정가 06:최유리지정가 07:최우선지정가 61:장개시전시간외종가 81:시간외종가 82:시간외단일가
	OrderType string
	//0:없음,1:IOC,2:FOK
	OrderCondition string
}

func (OrderEditOption) String() string {
	return "CSPAT00701"
}

func (t OrderEditOption) Path() string {
	return "/stock/order"
}

func (t OrderEditOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock1": {
	"OrgOrdNo": %d,
	"IsuNo": "%s",
	"OrdQty": %d,
	"OrdPrc": %d,
	"OrdprcPtnCode": "%s",
	"OrdCndiTpCode": "%s"
  }
}`, t.String(),
		t.OrderNumber,
		t.Ticker,
		t.Amount,
		t.Price,
		t.OrderType,
		t.OrderCondition,
	)), nil
}

type OrderCancelOption struct {
	// 주문번호
	OrderNumber int
	// 	주식 : 종목코드 or A+종목코드(모의투자는 A+종목코드) ELW : J+종목코드 ETN : Q+종목코드
	Ticker string
	Amount int
}

func (OrderCancelOption) String() string {
	return "CSPAT00801"
}

func (t OrderCancelOption) Path() string {
	return "/stock/order"
}

func (t OrderCancelOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock1": {
	"OrgOrdNo": %d,
	"IsuNo": "%s",
	"OrdQty": %d
  }
}`, t.String(),
		t.OrderNumber,
		t.Ticker,
		t.Amount,
	)), nil
}
