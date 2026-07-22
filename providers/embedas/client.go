// Package embedas interacts with the Embed.as API.
package embedas

// Host is the default base URL of the Embed.as API server
var Host = "https://embed.as"

// Client for the Embed.as API. No key is needed — API keys for
// third-party providers are configured server-side.
type Client struct {
	// host is the base URL of the Embed.as API server. When empty, the
	// package-level Host is used.
	host string
}

// NewClient creates a new Embed.as Client that talks to the default Host.
func NewClient() *Client {
	return &Client{}
}

// NewClientWithHost creates a new Embed.as Client that talks to the API at the
// given base URL
func NewClientWithHost(host string) *Client {
	return &Client{host: host}
}

// baseURL returns the base URL this client should use.
func (c *Client) baseURL() string {
	if c.host != "" {
		return c.host
	}
	return Host
}
