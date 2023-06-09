package service

import (
	"github.com/serenite11/Links-Reduction-Api/internal/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Service struct {
	LinksShortener
}

type LinksShortener interface {
	CreateShortUrl(longUrl string) (string, error)
	GetLongUrl(shortUrl string) (string, error)
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		NewLinksService(repo.LinksShortener),
	}
}
