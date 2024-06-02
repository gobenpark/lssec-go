package lssec_go

import "fmt"

// 테마별종목
type ThemeStockOption struct {
	Name string
	Code string
}

func (ThemeStockOption) String() string {
	return "t1531"
}

func (t ThemeStockOption) Path() string {
	return "/stock/sector"
}

func (t ThemeStockOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"tmname": "%s",
	"tmcode": "%s"
  }
}`, t.String(),
		t.Name,
		t.Code,
	)), nil
}

// 종목별테마
type StockThemeOption struct {
	Ticker string
}

func (StockThemeOption) String() string {
	return "t1532"
}

func (t StockThemeOption) Path() string {
	return "/stock/sector"
}

func (t StockThemeOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"shcode": "%s"
  }
}`, t.String(),
		t.Ticker,
	)), nil
}

// 특이테마
type SpecialThemeOption struct {
	// 1:상승율 상위 2:하락율 상위 3:거래증가율 상위 4:거래증가율 하위 5:상승종목비율 상위 6:상승종목비율 하위 7:기준대비 상승율 상위 8:기준대비 하락율 상위
	Ranking      string
	ContrastDate int
}

func (SpecialThemeOption) String() string {
	return "t1533"
}

func (t SpecialThemeOption) Path() string {
	return "/stock/sector"
}

func (t SpecialThemeOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"gubun": "%s",
	"chgdate": %d
  }
}`, t.String(),
		t.Ranking,
		t.ContrastDate,
	)), nil
}

// 테마종목별시세조회
type ThemeStockQuoteOption struct {
	Code string
}

func (ThemeStockQuoteOption) String() string {
	return "t1537"
}

func (t ThemeStockQuoteOption) Path() string {
	return "/stock/sector"
}

func (t ThemeStockQuoteOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"tmcode": "%s"
  }
}`, t.String(),
		t.Code,
	)), nil
}

// 전체테마
type AllThemeOption struct {
}

func (AllThemeOption) String() string {
	return "t8425"
}

func (AllThemeOption) Path() string {
	return "/stock/sector"
}

func (AllThemeOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"dummy":""
  }
}`, AllThemeOption{}.String(),
	)), nil
}
