package url

import "time"

type GetUrlResponse struct {
	Url       string    `json:"url"`
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"createdAt"`
}

type CreateUrlResponse struct {
	Url  string `json:"url"`
	Code string `json:"code"`
}

type CreateUrlRequest struct {
	Url string `json:"url"`
}
