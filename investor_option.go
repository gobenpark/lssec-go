package ebest_go

import "fmt"

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
