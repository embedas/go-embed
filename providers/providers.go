// Package providers interacts with all supported Embed.as providers.
package providers

import (
	"log"
	"net/url"
	"regexp"
	"strings"
)

var embedHostDomains = map[string]*regexp.Regexp{
	"airtable.com":          nil,
	"www.are.na":            nil,
	"cinnamon.video":        nil,
	"codepen.io":            nil,
	"flic.kr":               nil,
	"www.flickr.com":        regexp.MustCompile("/photos/.+"),
	"giphy.com":             regexp.MustCompile("/gifs/.+"),
	"media.giphy.com":       regexp.MustCompile("/media/.+"),
	"gph.is":                nil,
	"gist.github.com":       nil,
	"gfycat.com":            nil,
	"hypem.com":             nil,
	"imgur.com":             nil,
	"instagr.am":            regexp.MustCompile("/p/.+"),
	"www.instagram.com":     regexp.MustCompile("/p/.+"),
	"www.kickstarter.com":   regexp.MustCompile("/projects/.+/.+"),
	"www.last.fm":           regexp.MustCompile("/music/.+"),
	"www.mixcloud.com":      nil,
	"www.rdio.com":          nil,
	"open.spotify.com":      nil,
	"quora.com":             nil,
	"www.quora.com":         nil,
	"share.getcloudapp.com": nil,
	"soundcloud.com":        nil,
	"snd.sc":                nil,
	"twitter.com":           regexp.MustCompile("/.+/status/.+"),
	"vimeo.com":             nil,
	"imgs.xkcd.com":         nil,
	"xkcd.com":              nil,
	"www.xkcd.com":          nil,
	"youtu.be":              nil,
	"youtube.ca":            nil,
	"youtube.jp":            nil,
	"youtube.com.br":        nil,
	"youtube.co.uk":         nil,
	"youtube.nl":            nil,
	"youtube.pl":            nil,
	"youtube.es":            nil,
	"youtube.ie":            nil,
	"it.youtube.com":        nil,
	"youtube.fr":            nil,
	"youtube.com":           nil,
	"www.youtube.com":       nil,
}

// EmbedURL returns whether or not the given URL is supported. If true, the
// client should fetch the rich media for the URL.
func EmbedURL(us string) bool {
	u, err := url.Parse(us)
	if err != nil {
		log.Printf("[ERROR] EmbedURL: unable to parse url: %s", err)
		return false
	}

	// Ensure domain is from a supported provider
	ex, domainOK := embedHostDomains[u.Host]
	domainOK = domainOK ||
		strings.HasSuffix(u.Host, ".bandcamp.com") ||
		strings.HasSuffix(u.Host, ".tumblr.com") ||
		strings.HasSuffix(u.Host, ".typeform.com")
	if !domainOK {
		return false
	}

	if ex == nil {
		// All URL paths are allowed
		return true
	}

	// Check URL path for valid pattern
	return ex.MatchString(u.Path)
}
