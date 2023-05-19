package service

import (
	"github.com/serenite11/Links-Reduction-Api/internal/database"
	"github.com/serenite11/Links-Reduction-Api/internal/repository"
)

type Service struct {
	repo  *repository.Repository
	links *database.LinkInMemory
}

func NewService(repo *repository.Repository, links *database.LinkInMemory) *Service {
	return &Service{repo: repo, links: links}
}

func (s *Service) CreateShortUrl(longUrl string, shortUrl string) (string, error) {
	return s.repo.ILinksActions.CreateShortUrl(longUrl, shortUrl)
}

func (s *Service) GetLongUrl(shortUrl string) (string, error) {
	return s.repo.ILinksActions.GetLongUrl(shortUrl)
}

func (s *Service) CreateShortUrlMemory(longUrl string, shortUrl string) string {
	return s.links.CreateShortUrl(longUrl)
}
func (s *Service) GetLongUrlMemory(shortUrl string) string {
	return s.links.GetLongUrl(shortUrl)
}
