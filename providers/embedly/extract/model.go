// Package extract holds models for Embed.ly's Extract API.
package extract

// Options are properties that can be sent to the Extract API.
type Options struct {
	MaxWidth     int
	MaxHeight    int
	Width        int
	Words        int
	Chars        int
	WMode        bool
	AllowScripts bool
	NoStyle      bool
	Autoplay     bool
	VideoSrc     bool
	Frame        bool
	Secure       bool
}
