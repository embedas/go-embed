package embedly

var (
	Host = "https://api.embed.ly"
)

// Key is the Embed.ly API key.
var Key string

type Client struct {
	key string
}

// NewClient creates a Client with the given Embed.ly API key
func NewClient(key string) *Client {
	return &Client{key}
}

// DefaultClient creates a Client using the package's Key value
func DefaultClient() *Client {
	return &Client{Key}
}
