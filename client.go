package ebest_go

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
)

type Option interface {
	String() string
}

type Client struct {
	appKey      string
	appSecret   string
	accessToken string
	aCache      bool
	expire      time.Duration
	cli         *resty.Client
}

// TODO: remove and recreate
func saveToken(token string, expire time.Duration) error {

}

func NewClient(options ...ClientOption) *Client {
	client := &Client{}

	for i := range options {
		options[i](client)
	}
	cli := resty.New()
	cli.SetBaseURL("https://openapi.ebestsec.co.kr:8080")
	cli.SetDebug(true)
	client.cli = cli

	if client.aCache {
		f, err := os.Open("token")
		if err != nil && errors.Is(err, os.ErrNotExist) {
			tf, err := os.Create("token")
			if err != nil {
				panic(err)
			}
			defer tf.Close()
			tk, err := client.AccessToken(context.Background())
			if err != nil {
				panic(err)
			}

			client.accessToken = string(tk)
			date := time.Now().Add(client.expire)
			expiredate := date.Format("20060102")
			if _, err := tf.Write([]byte(tk + "::" + expiredate)); err != nil {
				panic(err)
			}
		} else if err != nil {
			panic(err)
		} else {
			bt, err := io.ReadAll(f)
			if err != nil {
				panic(err)
			}

			data := strings.Split(string(bt), "::")
			if len(data) != 2 {
				panic("invalid token file")
			}

			expireTime := time.Parse("20060102", data[1])
			if time.Now().After(expireTime) {
				//TODO: refresh token
			}

			client.accessToken = string(data[0])
		}
		defer f.Close()

	}

	cli.OnBeforeRequest(func(c *resty.Client, r *resty.Request) error {
		switch {
		case client.appKey == "":
			return errors.New("empty app key")
		case client.appSecret == "":
			return errors.New("empty app secret")
		case client.accessToken == "":
			return errors.New("empty access token")
		}
		r.SetHeader("authorization", "Bearer "+client.accessToken)
		return nil
	})

	return client
}

func (c *Client) AccessToken(ctx context.Context) (string, error) {
	res, err := c.cli.R().SetFormData(map[string]string{
		"grant_type":   "client_credentials",
		"appkey":       c.appKey,
		"appsecretkey": c.appSecret,
		"scope":        "oob",
	}).Post("/oauth2/token")
	if err != nil {
		return "", err
	}
	fmt.Println(res.String())

	return gjson.GetBytes(res.Body(), "access_token").String(), nil
}
