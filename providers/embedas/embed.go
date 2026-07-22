package embedas

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// Embed returns rich media responses for the given URLs.
func (c *Client) Embed(options Options, urls ...string) ([]Response, error) {
	responses := make([]Response, len(urls))
	for i, u := range urls {
		res, err := c.embed(options, u)
		if err != nil {
			return nil, err
		}
		responses[i] = *res
	}
	return responses, nil
}

// EmbedOne returns a rich media response for a single URL.
func (c *Client) EmbedOne(options Options, rawURL string) (*Response, error) {
	return c.embed(options, rawURL)
}

func (c *Client) embed(options Options, rawURL string) (*Response, error) {
	u, err := url.Parse(c.baseURL() + "/api")
	if err != nil {
		return nil, fmt.Errorf("url.Parse: %s", err)
	}

	q := u.Query()
	q.Set("url", rawURL)
	if options.MaxWidth > 0 {
		q.Set("maxwidth", strconv.Itoa(options.MaxWidth))
	}
	u.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Get(u.String())
	if err != nil {
		return nil, fmt.Errorf("GET: %s", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}
	return &response, nil
}
