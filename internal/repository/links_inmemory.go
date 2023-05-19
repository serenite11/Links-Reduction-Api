package repository

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

var symbols = []byte{
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
	'1', '2', '3', '4', '5', '6', '7', '8', '9', '0',
	'_',
}

type LinkInMemory struct {
	Links map[string]string `json:"links"`
}

func NewLinkInMemory(links map[string]string) *LinkInMemory {
	return &LinkInMemory{Links: links}
}

func (l *LinkInMemory) CreateShortUrl(longUrl string) (string, error) {
	shortUrl := generateShortUrl(longUrl)
	l.Links[longUrl] = shortUrl
	log.Print(l.Links)
	return l.Links[longUrl], nil
}

func (l *LinkInMemory) GetLongUrl(shortUrl string) (string, error) {
	for key, value := range l.Links {
		if value == shortUrl {
			return key, nil
		}
	}
	return "", errors.New("No origin url for this short url")
}
func generateShortUrl(url string) string {
	var hash string
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		hash += string(symbols[rand.Intn(62)])
	}
	return hash
}
