package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/serenite11/Links-Reduction-Api/internal/models"
)

type LinksPostgres struct {
	db *sqlx.DB
}

func NewLinksPostgres(db *sqlx.DB) *LinksPostgres {
	return &LinksPostgres{db: db}
}

func (r *LinksPostgres) CreateShortUrl(longUrl string) (string, error) {
	shortUrl := generateShortUrl(longUrl)
	query := fmt.Sprintf("INSERT INTO links (long_url,short_url) values ($1,$2)")
	_, err := r.db.Query(query, longUrl, shortUrl)
	if err != nil {
		return "", err
	}
	return shortUrl, nil
}

func (r *LinksPostgres) GetLongUrl(shortUrl string) (string, error) {
	var link models.Link
	query := fmt.Sprintf("SELECT long_url FROM links WHERE short_url=$1")
	err := r.db.Get(&link, query, shortUrl)
	if err != nil {
		return "", err
	}
	return link.LongUrl, nil
}
