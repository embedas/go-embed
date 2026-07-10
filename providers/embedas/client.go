// Package embedas interacts with the Embed.as API.
package embedas

// Host is the base URL of the Embed.as API server.
var Host = "https://embed.as"

// Client for the Embed.as API. No key is needed — API keys for
// third-party providers are configured server-side.
type Client struct{}

// NewClient creates a new Embed.as Client.
func NewClient() *Client {
	return &Client{}
}
