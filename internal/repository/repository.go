package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	ILinksActions
}

type ILinksActions interface {
	CreateShortUrl(longUrl string, shortUrl string) (string, error)
	GetLongUrl(shortUrl string) (string, error)
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		NewLinksPostgres(db),
	}
}
