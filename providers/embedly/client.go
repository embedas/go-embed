package embedly

var (
	Host = "https://api.embed.ly"
)

type Client struct {
	key string
}

func NewClient(key string) *Client {
	return &Client{key}
}
