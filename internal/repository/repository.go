package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	LinksShortener
}

type LinksShortener interface {
	CreateShortUrl(longUrl string) (string, error)
	GetLongUrl(shortUrl string) (string, error)
}

func NewRepositoryPostgres(db *sqlx.DB) *Repository {
	return &Repository{
		NewLinksPostgres(db),
	}
}

func NewRepositoryInMemory(links map[string]string) *Repository {
	return &Repository{
		NewLinkInMemory(links),
	}
}
