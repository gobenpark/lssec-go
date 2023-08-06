package ebest_go

import "time"

type ClientOption func(*Client)

func WithAuth(appKey, secret string) ClientOption {
	return func(c *Client) {
		c.appKey = appKey
		c.appSecret = secret
	}
}

func WithAccessToken(token string) ClientOption {
	return func(c *Client) {
		c.accessToken = token
	}
}

func WithAutomaticTokenCache(cache bool, expire time.Duration) ClientOption {
	return func(client *Client) {
		client.aCache = cache
	}
}
