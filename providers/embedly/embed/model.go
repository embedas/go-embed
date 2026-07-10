// Package embed holds models for Embed.ly's Embed API.
package embed

import "github.com/embedas/go-embed/oembed"

// Options are properties that can be sent to the Embed API.
type Options struct {
	MaxWidth int
}

// Response is a rich media response from the Embed API.
type Response = oembed.Response
