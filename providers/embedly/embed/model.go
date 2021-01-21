package embed

type Options struct {
	MaxWidth int
}

type Response struct {
	URL          string `json:"url"`
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message,omitempty"`
	Type         string `json:"type"`
	Version      string `json:"version"`

	ProviderName    string `json:"provider_name"`
	ProviderURL     string `json:"provider_url"`
	OriginalURL     string `json:"original_url"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	AuthorName      string `json:"author_name"`
	AuthorURL       string `json:"author_url"`
	HTML            string `json:"html"`
	Width           int    `json:"width"`
	Height          int    `json:"height"`
	ThumbnailURL    string `json:"thumbnail_url"`
	ThumbnailWidth  int    `json:"thumbnail_width"`
	ThumbnailHeight int    `json:"thumbnail_height"`
}
