package ebest_go

import "fmt"

// 거래원 정보

// 종목별 상위 거래원
type TopMemberCompany struct {
	Ticker    string
	StartDate string
	EndDate   string
	// 전체:0 외국계만: 1
	Foreigner bool
}

func (t TopMemberCompany) MarshalJSON() ([]byte, error) {
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

func (TopMemberCompany) String() string {
	return "t1752"
}

func (TopMemberCompany) Path() string {
	return "/stock/exchange"
}

// 종목별 회원사 추이
type MemberShipTrandsByStock struct {
	Ticker       string
	MemberNumber string `json:"tradno"`
	StartDate    string
	EndDate      string
	// 0: 시간별 1: 일별
	Gubun string
	Count int
}

func (t MemberShipTrandsByStock) MarshalJSON() ([]byte, error) {
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
		t.Gubun,
		t.Count,
	)), nil
}

func (MemberShipTrandsByStock) String() string {
	return "t1771"
}

func (MemberShipTrandsByStock) Path() string {
	return "/stock/exchange"
}
