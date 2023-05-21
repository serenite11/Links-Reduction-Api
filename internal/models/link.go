package models

type Link struct {
	LongUrl  string `json:"long_url" db:"long_url" `
	ShortUrl string `json:"short_url" db:"short_url" `
}
