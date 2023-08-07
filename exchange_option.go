package ebest_go

import "fmt"

// 거래원 정보

// 종목별 상위 거래원
type TopMemberCompanyOption struct {
	Ticker    string
	StartDate string
	EndDate   string
	// 전체:0 외국계만: 1
	Foreigner bool
}

func (t TopMemberCompanyOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
    "shcode": "%s",
    "traddate1": "%s",
    "traddate2": "%s",
	"fwgubun1": "%s"
  }
}`, t.String(),
		t.Ticker,
		t.StartDate,
		t.EndDate,
		func() string {
			if t.Foreigner {
				return "1"
			}
			return "0"
		}(),
	)), nil
}

func (TopMemberCompanyOption) String() string {
	return "t1752"
}

func (TopMemberCompanyOption) Path() string {
	return "/stock/exchange"
}

// 종목별 회원사 추이
type MemberShipTrandsByStockOption struct {
	Ticker       string
	MemberNumber string `json:"tradno"`
	StartDate    string
	EndDate      string
	// 0: 시간별 1: 일별
	TimeOrDate string
	Count      int
}

func (t MemberShipTrandsByStockOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
    "shcode": "%s",
	"tradno": "%s",
    "traddate1": "%s",
    "traddate2": "%s",
	"gubun1": "%s",
	"cnt": %d
  }
}`, t.String(),
		t.Ticker,
		t.MemberNumber,
		t.StartDate,
		t.EndDate,
		t.TimeOrDate,
		t.Count,
	)), nil
}

func (MemberShipTrandsByStockOption) String() string {
	return "t1771"
}

func (MemberShipTrandsByStockOption) Path() string {
	return "/stock/exchange"
}
