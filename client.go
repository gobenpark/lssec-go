package ebest_go

import (
	"context"
	"errors"

	"github.com/go-resty/resty/v2"
)

type Menu string

const (
	// 업종 기간별 추이
	IndustryTrandsOverTime = "t1514"
	// 전체업종
	TotalIndustry Menu = "t8436"
	// 예상지수
	ExpectedStockIndex Menu = "t1485"
	// 업종현재가
	CurrentPriceOfIndustry Menu = "t1511"
	// 업종별 종목시세
	PriceOfIndustry Menu = "t1516"
)

type Client struct {
	appKey      string
	appSecret   string
	accessToken string
	aCache      bool
	cli         *resty.Client
}

func NewClient(options ...ClientOption) *Client {
	client := &Client{}

	for i := range options {
		options[i](client)
	}
	cli := resty.New()
	cli.SetBaseURL("https://openapi.ebestsec.co.kr:8080")
	cli.OnBeforeRequest(func(c *resty.Client, r *resty.Request) error {
		switch {
		case client.accessToken == "":
			return errors.New("empty access token")
		case client.appKey == "":
			return errors.New("empty app key")
		case client.appSecret == "":
			return errors.New("empty app secret")
		}
		r.SetHeader("authorization", "Bearer "+client.accessToken)
		return nil
	})
	return client
}

func (c *Client) MarketData(ctx context.Context, menu Menu) {
	c.cli.R().
		SetContext(ctx).
		SetHeaders(map[string]string{
			"tr_cd":   string(menu),
			"tr_cont": "Y",
		}).
		SetBody(map[string]string{})
}

func (c *Client) AccessToken(ctx context.Context) {
	c.cli.R().SetBody(map[string]string{
		"grant_type": "client_credentials",
		"appkey":     c.appKey,
		"appsecret":  c.appSecret,
		"scope":      "oob",
	}).Post("/oauth2/token")
}

func (c *Client) PeriodicTrendOfStocks() {

	res, err := c.cli.R().
		SetHeaders(map[string]string{
			"tr_cd":       "t1513",
			"tr_cont":     "Y",
			"tr_cont_key": "",
		}).
		SetBody(map[string]string{
			"t1513InBlock": "t1513InBlock",
		}).
		Post("https://openapi.ebestsec.co.kr:8080")
	if err != nil {
		return nil, err
	}
}
