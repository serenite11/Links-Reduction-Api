package database

import "github.com/serenite11/Links-Reduction-Api/internal/handlers"

type LinkInMemory struct {
	Links map[string]string `json:"links"`
}

func NewLinkInMemory() *LinkInMemory {
	return &LinkInMemory{Links: map[string]string{}}
}

func (l *LinkInMemory) CreateShortUrl(longUrl string) string {
	l.Links[longUrl] = handlers.GenerateShortUrl(longUrl)
	return l.Links[longUrl]
}

func (l *LinkInMemory) GetLongUrl(shortUrl string) string {
	for key, value := range l.Links {
		if value == shortUrl {
			return key
		}
	}
	return ""
}
