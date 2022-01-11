package giphy

import "net/http"

// Option function used by the Giphy client
type Option func(*Client)

// APIKey option used by the Giphy client
func APIKey(apiKey string) Option {
	return func(c *Client) {
		c.APIKey = apiKey
	}
}

// Rating option used by the Giphy client
func Rating(rating string) Option {
	return func(c *Client) {
		c.Rating = rating
	}
}

// Limit option used by the Giphy client
func Limit(limit int) Option {
	return func(c *Client) {
		c.Limit = limit
	}
}

// HTTPClient option used by the Giphy client
func HTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}
