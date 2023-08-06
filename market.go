package ebest_go

import (
	"context"
	"fmt"
)

// 업종 기간별 추이
// example"
//
//	{
//		"t1514InBlock": {
//		"upcode": "001",
//		"gubun1": " ",
//		"gubun2": "1",
//		"cts_date": " ",
//		"cnt": 1,
//		"rate_gbn": "1"
//		}
//	}
type IndustryTrendsOverTimeOption struct {
	Code string `json:"upcode"`
	//일: 1, 주: 2, 월: 3, 분: 4
	Date    string `json:"gubun2"`
	CtsDate string `json:"cts_date"`
	Count   int    `json:"cnt"`
	//1:거래량비중 - 2:거래대금비중
	Rate int `json:"rate_gbn"`
}

func (IndustryTrendsOverTimeOption) String() string {
	return "t1514"
}

func (i IndustryTrendsOverTimeOption) MarshalJSON() ([]byte, error) {
	body := fmt.Sprintf(`
{
	"%sInBlock": {
		"upcode": "%s",
		"gubun1": " ",
		"gubun2": "%s",
		"cts_date": "%s",
		"cnt": %d,
		"rate_gbn": "%d"
	}
}
`, i.String(), i.Code, i.Date, i.CtsDate, i.Count, i.Rate)

	return []byte(body), nil
}

// 전체 업종
type TotalIndustryOption struct{}

func (TotalIndustryOption) String() string {
	return "t8424"
}

func (i TotalIndustryOption) MarshalJSON() ([]byte, error) {

	return []byte(`
{
	"t8424InBlock": {
		"gubun1": ""
	}
}`), nil
}

// 예상지수
type ExpectedStockIndexOption struct {
	// 코스피:001 코스닥:301 KRX100:501 KP200:101 SRI:515 코스닥프리미어:404 KRX 보험:516 KRX 운송:517
	Code string
	// 1: 장전, 2: 장후
	Gubun string
}

func (ExpectedStockIndexOption) String() string {
	return "t1485"
}

func (i ExpectedStockIndexOption) MarshalJSON() ([]byte, error) {
	body := fmt.Sprintf(`
{
	"%sInBlock": {
		"upcode": "%s",
		"gubun": "%s"
	}
}
`, i.String(), i.Code, i.Gubun)

	return []byte(body), nil
}

type CurrentPriceOfIndustryOption struct {
	Code string
}

func (CurrentPriceOfIndustryOption) String() string {
	return "t1511"
}

func (i CurrentPriceOfIndustryOption) MarshalJSON() ([]byte, error) {
	body := fmt.Sprintf(`
{
	"%sInBlock": {
		"upcode": "%s"
	}
}
`, i.String(), i.Code)

	return []byte(body), nil
}

// 업종별 종목시세
type PriceOfIndustryOption struct {
	Code   string
	Gubun  string
	Ticker string
}

func (PriceOfIndustryOption) String() string {
	return "t1516"
}

func (i PriceOfIndustryOption) MarshalJSON() ([]byte, error) {
	body := fmt.Sprintf(`
{
	"%sInBlock": {
		"upcode": "%s",
		"gubun": "%s",
		"shcode": "%s"
	}
}
`, i.String(), i.Code, i.Gubun, i.Ticker)

	return []byte(body), nil
}

func (c *Client) MarketData(ctx context.Context, contKey string, data Option) ([]byte, error) {
	headers := map[string]string{
		"tr_cd":        data.String(),
		"tr_cont":      "N",
		"content-type": "application/json; charset=utf-8",
	}
	if contKey != "" {
		headers["tr_cont"] = "Y"
		headers["tr_cont_key"] = contKey
	}
	res, err := c.cli.R().
		SetContext(ctx).
		SetHeaders(headers).
		SetBody(data).Post("/indtp/market-data")

	if err != nil {
		return nil, err
	}
	return res.Body(), nil
}
