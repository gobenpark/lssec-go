package ebest_go

import (
	"context"
)

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
