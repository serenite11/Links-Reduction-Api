package service

import "github.com/serenite11/Links-Reduction-Api/internal/repository"

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateShort(longUrl string, shortUrl string) (string, error) {
	return s.repo.ILinksActions.CreateShortUrl(longUrl, shortUrl)
}

func (s *Service) GetLongUrl(shortUrl string) (string, error) {
	return s.repo.ILinksActions.GetLongUrl(shortUrl)
}
