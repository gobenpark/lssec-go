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

type ForeignerAgencyOption2 struct {
	Ticker    string
	StartDate string
	EndDate   string
}

func (ForeignerAgencyOption2) String() string {
	return "t1716"
}

func (t ForeignerAgencyOption2) Path() string {
	return "/stock/frgr-itt"
}

func (t ForeignerAgencyOption2) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"shcode": "%s",
	"fromdt": "%s",
	"gubun": "0",
	"todt": "%s",
	"prapp": 0,
	"progubun": "0",
	"orggubun": "0",
	"frggubun": "0"
  }
}`, t.String(),
		t.Ticker,
		t.StartDate,
		t.EndDate,
	)), nil
}

type ForeignerAgencyOption3 struct {
	Ticker    string
	StartDate string
	EndDate   string
}

func (ForeignerAgencyOption3) String() string {
	return "t1717"
}

func (t ForeignerAgencyOption3) Path() string {
	return "/stock/frgr-itt"
}

func (t ForeignerAgencyOption3) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"shcode": "%s",
	"fromdt": "%s",
	"gubun": "0",
	"todt": "%s"
  }
}`, t.String(),
		t.Ticker,
		t.StartDate,
		t.EndDate,
	)), nil
}
