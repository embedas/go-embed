package embedas

import "github.com/embedas/go-embed/oembed"

// Response is an oEmbed response from the Embed.as API.
type Response = oembed.Response

// Options are properties that can be sent to the Embed.as API.
type Options struct {
	// MaxWidth, when greater than 0, is sent as the maxwidth hint to size the
	// returned embed. Zero leaves sizing to the API's server-side default.
	MaxWidth int
}
