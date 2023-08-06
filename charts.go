package ebest_go

import (
	"context"
	"fmt"
	"time"
)

type TotalChartOption struct {
	Code string
	// 0: 틱,1: 분, 2:일, 3:주, 4:월
	ChartType string

	TickCount  int
	QueryCount int
	// 0: 전체, 1: 당일만
	Tdgb      string
	StartDate time.Time
	EndDate   time.Time
	CtsDate   time.Time
	CtsTime   time.Time
	CtsDaygb  string
}

func (TotalChartOption) String() string {
	return "t4203"
}

func (t TotalChartOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
    "shcode": "%s",
    "gubun": "%s",
    "ncnt": %d,
    "qrycnt": %d,
    "tdgb": "%s",
    "sdate": "%s",
    "edate": "%s",
    "cts_date": "%s",
    "cts_time": "%s",
    "cts_daygb": "%s"
  }
}`, t.String(),
		t.Code,
		t.ChartType,
		t.TickCount,
		t.QueryCount,
		t.Tdgb,
		t.StartDate.Format("20060102"),
		t.EndDate.Format("20060102"),
		t.CtsDate.Format("20060102"),
		t.CtsTime.Format("20060102"),
		t.CtsDaygb,
	)), nil
}

func (c *Client) Charts(ctx context.Context, contKey string, option Option) ([]byte, error) {
	headers := map[string]string{
		"tr_cd":        option.String(),
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
		SetBody(option).
		Post("/indtp/chart")
	if err != nil {
		return nil, err
	}
	return res.Body(), nil

}
