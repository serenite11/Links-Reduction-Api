package handlers

import (
	"errors"
	"net/http"
)

func validateCreateShortUrl(link string) error {
	if link == "" {
		return errors.New("Empty request")
	}
	if _, err := http.Get(link); err != nil {
		return errors.New("Url is not valid")
	}
	return nil
}
