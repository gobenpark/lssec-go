package ebest_go

import (
	"context"
)

func (c *Client) MarketSectorData(ctx context.Context, contKey string, data Option) ([]byte, error) {
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

func (c *Client) MarketData(ctx context.Context, contKey string, data Option) ([]byte, error) {
	headers := map[string]string{
		"tr_cd":        data.String(),
		"tr_cont":      "N",
		"content-type": "application/json; charset=utf-8",
	}
	//if contKey != "" {
	headers["tr_cont"] = "Y"
	//headers["tr_cont_key"] = contKey
	//}
	res, err := c.cli.R().
		SetContext(ctx).
		SetHeaders(headers).
		SetBody(data).Post("/stock/market-data")

	if err != nil {
		return nil, err
	}
	return res.Body(), nil
}
