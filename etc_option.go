package ebest_go

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
