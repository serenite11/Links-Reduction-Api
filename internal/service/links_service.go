package service

import "github.com/serenite11/Links-Reduction-Api/internal/repository"

type LinksService struct {
	repo repository.LinksShortener
}

func NewLinksService(repo repository.LinksShortener) *LinksService {
	return &LinksService{repo: repo}
}

func (l *LinksService) CreateShortUrl(longUrl string) (string, error) {
	return l.repo.CreateShortUrl(longUrl)
}

func (l *LinksService) GetLongUrl(shortUrl string) (string, error) {
	return l.repo.GetLongUrl(shortUrl)
}
