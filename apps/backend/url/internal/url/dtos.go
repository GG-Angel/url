package url

type CreateUrlRequest struct {
	Url  string   `json:"url"`
	Tags []string `json:"tags"`
	Slug *string  `json:"slug,omitempty"`
}
