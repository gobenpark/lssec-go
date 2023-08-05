package ebest_go

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

func WithAutomaticTokenCache(cache bool) ClientOption {
	return func(client *Client) {

	}
}
