package ebest_go

import "fmt"

// 외인기관종목별동향
type ForeignerAgencyOption1 struct {
	Ticker        string
	EndDate       string
	PriceOrAmount string
	// 0:순매수1:매수2:매도
	TradeType string
	//누적구분(0:일간1:누적)
	AccType        string
	CountinuesDate string
	Index          int
}

func (ForeignerAgencyOption1) String() string {
	return "t1702"
}

func (t ForeignerAgencyOption1) Path() string {
	return "/stock/frgr-itt"
}

func (t ForeignerAgencyOption1) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"shcode": "%s",
	"todt": "%s",
	"volvalgb": "%s",
	"msmdgb": "%s",
	"cumulgb": "%s",
	"cts_date": "%s",
	"ctx_idx": %d
  }
}`, t.String(),
		t.Ticker,
		t.EndDate,
		t.PriceOrAmount,
		t.TradeType,
		t.AccType,
		t.CountinuesDate,
		t.Index,
	)), nil
}
