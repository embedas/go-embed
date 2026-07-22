# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
# Run all tests
go test ./...

# Run a single test
go test ./providers/... -run TestName

# Build
go build ./...

# Vet
go vet ./...
```

No external dependencies — the module has no `require` entries in `go.mod`.

## Architecture

This library has three layers:

**`providers/` — URL support detection**
- `providers/providers.go`: The `EmbedURL(url string) bool` function determines whether a given URL is from a supported media provider. It uses a map of host domains to optional `*regexp.Regexp` path validators. Domains with `nil` regexps allow any path; those with a regexp require the path to match. Some providers (Bandcamp, Tumblr, Typeform) use subdomain suffix matching instead.

**`providers/embedas/` — Embed.as API client** (preferred)
- `embed.go`: `Embed(options, urls...)` / `EmbedOne(options, url)` fetch `GET {host}/api?url=...` and decode the JSON oEmbed response. Third-party *provider* keys are configured server-side on the Embed.as instance; the *caller's* Embed.as API key (below) is sent as the `X-API-Key` header.
- `client.go`: `NewClient(key)` talks to the default `Host` (`https://embed.as`); `NewClientWithHost(host, key)` targets a custom base URL (e.g. `http://embed.local` for local development). An empty host falls back to `Host`; an empty key sends no `X-API-Key` header (for instances that don't require auth).
- `model.go`: `Response` is an alias for the shared `oembed.Response`. `Options` carries request hints — currently `MaxWidth`, sent as the `maxwidth` query param when greater than 0 (0 leaves sizing to the API's server-side default).

**`oembed/` — shared oEmbed response model**
- `response.go`: `Response` holds the standard oEmbed fields (HTML, thumbnails, dimensions, etc.). No external dependencies.

**`providers/embedly/` — Embed.ly API client** (legacy)
- Two APIs are supported:
  - **Embed** (`embed.go`): calls `/1/oembed`, batches URLs in groups of 10, returns `[]embed.Response` with oEmbed fields (HTML, thumbnails, etc.)
  - **Extract** (`extract.go`): calls `/1/extract`, same batching logic, returns `[]Response` (the full `embedly.Response` with richer metadata: authors, keywords, entities, images)
- `client.go`: `Client` holds the API key. Use `NewClient(key)` or set the package-level `Key` var and call `DefaultClient()`.
- `model.go`: Defines `embedly.Response` (used by Extract) and media type constants.
- `embed/model.go`: Defines `embed.Options` and `embed.Response` (used by Embed API).
- `extract/model.go`: Defines `extract.Options` (used by Extract API).

**Adding a new provider**: add its domain to the `embedHostDomains` map in `providers/providers.go` with either `nil` (any path) or a compiled regexp for path validation. For subdomain-based providers, add a `strings.HasSuffix` check in `EmbedURL`.