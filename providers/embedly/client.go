package embedly

var (
	Host = "https://api.embed.ly"
)

// Key is the Embed.ly API key
var Key string

type Client struct {
	key string
}

func NewClient(key string) *Client {
	return &Client{key}
}

func DefaultClient() *Client {
	return &Client{Key}
}
