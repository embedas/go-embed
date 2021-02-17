// Package embedly interacts with Embed.ly's APIs.
package embedly

import (
	"encoding/json"
	"fmt"
	"github.com/embedas/go-embed/providers/embedly/embed"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Embed returns rich media responses for the given URLs.
func (c *Client) Embed(options embed.Options, urls ...string) ([]embed.Response, error) {
	responses := make([]embed.Response, len(urls))
	for i := 0; i < len(urls); i += 10 {
		to := len(urls)
		if to > i+10 {
			to = i + 10
		}
		res, err := c.embed(urls[i:to], options)
		if err != nil {
			return nil, err
		}

		reslen := to - i
		if reslen > len(res) {
			reslen = len(res)
		}
		for j := 0; j < reslen; j++ {
			responses[i+j] = res[j]
		}
	}
	return responses, nil
}

// EmbedOne returns a rich media response for a single URL.
func (c *Client) EmbedOne(options embed.Options, url string) (*embed.Response, error) {
	responses, err := c.embed([]string{url}, options)
	if err != nil {
		return nil, err
	}
	return &responses[0], nil
}

func (c *Client) embed(urls []string, options embed.Options) ([]embed.Response, error) {
	u, err := url.Parse(Host + "/1/oembed")
	if err != nil {
		return nil, fmt.Errorf("url.Parse: %s", err)
	}

	q := u.Query()
	q.Add("key", c.key)
	addInt(&q, "maxwidth", options.MaxWidth)

	for i, u := range urls {
		urls[i] = url.QueryEscape(u)
	}
	if len(urls) == 0 {
		return nil, fmt.Errorf("At least one URL is required")
	} else {
		for _, url := range urls {
			if len(url) == 0 {
				return nil, fmt.Errorf("A URL cannot be empty")
			}
		}
		q.Set("urls", strings.Join(urls, ","))
	}

	// Embedly needs literal commas in the `urls` field, so replace the URL-encoded values with the
	// literal commas here, up to the number of URLs-1.
	un, _ := url.QueryUnescape(q.Encode())
	u.RawQuery = un
	urlStr := u.String()

	cl := http.DefaultClient
	resp, err := cl.Get(urlStr)
	if err != nil {
		return nil, fmt.Errorf("GET: %s", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("Got non-2xx status code: %s. %q", resp.Status, body)
	}

	response := []embed.Response{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}
