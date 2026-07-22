// Package embedas interacts with the Embed.as API.
package embedas

// Host is the default base URL of the Embed.as API server
var Host = "https://embed.as"

// Client for the Embed.as API. The key, when set, is sent as the X-API-Key header
// on every request; Embed.as instances that require authentication reject requests
// without a valid key.
type Client struct {
	// host is the base URL of the Embed.as API server. When empty, the
	// package-level Host is used.
	host string
	// key is the Embed.as API key sent as X-API-Key. Empty sends no key.
	key string
}

// NewClient creates a new Embed.as Client that talks to the default Host with the
// given API key.
func NewClient(key string) *Client {
	return &Client{key: key}
}

// NewClientWithHost creates a new Embed.as Client that talks to the API at the
// given base URL with the given API key.
func NewClientWithHost(host, key string) *Client {
	return &Client{host: host, key: key}
}

// baseURL returns the base URL this client should use.
func (c *Client) baseURL() string {
	if c.host != "" {
		return c.host
	}
	return Host
}
